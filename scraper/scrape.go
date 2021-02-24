package scraper

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	pb "auto_bidder/RPC/protos"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/ysmood/gson"
)

// Scraper ...
type Scraper struct {
	Browser *rod.Browser
	BidPath BidPath
	Page    *rod.Page
}

//go:embed evalProject.js
var evalProjectFile string

//go:embed eval.js
var evalFile string

func (s *Scraper) Start(url string, headless bool) {
	l := launcher.New().Headless(headless)

	u := l.MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	s.Browser = browser
	if f, err := os.Stat("./login_cookie.json"); err == nil && f.Size() > 200 {
		file, err := os.Open("./login_cookie.json")
		CheckError(err)
		var cookieParams []*proto.NetworkCookieParam
		bytes, err := ioutil.ReadAll(file)
		CheckError(err)
		err = json.Unmarshal(bytes, &cookieParams)
		CheckError(err)

		s.Page = browser.MustPage(url).
			MustSetCookies(cookieParams[0]).
			MustSetCookies(cookieParams[3]).
			MustWaitIdle().
			MustNavigate("https://www.freelancer.com/search/projects/").
			MustWaitIdle()

	} else {
		s.Page = browser.MustPage(url).MustWaitLoad()
	}

}

func (s *Scraper) Login(cred LoginInfo) error {
	usernameInput := s.Page.MustElementX(cred.UsernameInputPath)
	passwordInput := s.Page.MustElementX(cred.PasswordInputPath)
	buttonInput := s.Page.MustElementX(cred.ButtonPath)
	usernameInput.MustWaitInteractable().MustInput(cred.Username)
	passwordInput.MustWaitInteractable().MustInput(cred.Password)
	buttonInput.MustWaitInteractable().MustClick()

	s.Page.MustWaitNavigation()()

	cookies := s.Page.MustCookies()

	var filteredCookies []*proto.NetworkCookie

	for _, cookie := range cookies {
		if cookie.Name == "GETAFREE_AUTH_HASH_V2" {
			filteredCookies = append(filteredCookies, cookie)
		} else if cookie.Name == "GETAFREE_USER_ID" {
			filteredCookies = append(filteredCookies, cookie)
		}
	}

	if len(filteredCookies) < 2 {
		return errors.New("username or password is incorrect or check if your XPath is correct in the settings")
	}

	bytes, err := json.Marshal(proto.CookiesToParams(cookies))

	CheckError(err)

	file, err := os.Create("./login_cookie.json")

	CheckError(err)

	count, err := file.Write(bytes)

	CheckError(err)

	fmt.Println("wrote ", count)

	return nil
}

func (s *Scraper) GetProjects(pfunc func([]*pb.Projects)) func() error {

	// s.Page.MustWaitNavigation()()

	s.Page.WaitLoad()

	var projects []*pb.Projects

	stop, _ := s.Page.Expose("GetProjectList", func(v gson.JSON) (interface{}, error) {

		projectArr := v.Arr()

		for _, p := range projectArr {

			var project pb.Projects

			project.BiddingPrice = p.Get("biddingPrice").String()
			project.Description = p.Get("description").String()
			project.Currency = p.Get("currency").String()
			project.Link = p.Get("link").String()
			project.Selected = p.Get("selected").Bool()
			project.Skills = p.Get("skills").String()
			project.Title = p.Get("title").String()

			projects = append(projects, &project)

		}

		pfunc(projects)

		return nil, nil
	})

	// file, _ := os.Open("./evalProject.js")
	// bytes, _ := ioutil.ReadAll(file)

	_, err := s.Page.Evaluate(&rod.EvalOptions{
		JS: evalProjectFile,
	})

	CheckError(err)

	return stop

}

func (s *Scraper) SubscribeToProject(pfunc func(*pb.Projects)) {

	s.Page.MustWaitNavigation()()

	s.Page.MustWaitLoad()

	s.Page.Expose("GetProject", func(v gson.JSON) (interface{}, error) {
		var project pb.Projects

		project.BiddingPrice = v.Get("biddingPrice").String()
		project.Description = v.Get("description").String()
		project.Currency = v.Get("currency").String()
		project.Link = v.Get("link").String()
		project.Selected = v.Get("selected").Bool()
		project.Skills = v.Get("skills").String()
		project.Title = v.Get("title").String()

		pfunc(&project)

		return nil, nil
	})

	// file, _ := os.Open("./eval.js")
	// bytes, _ := ioutil.ReadAll(file)

	_, err := s.Page.Evaluate(&rod.EvalOptions{
		JS: evalFile,
	})

	CheckError(err)

}

func (s *Scraper) BidOnProject(url string, bidInfo BidInfo) bool {

	page := s.Browser.MustPage(url)

	page.WaitLoad()

	if bidInfo.BidAmount != "" {
		bidAmountInput := page.MustElementX(bidInfo.BidAmountPath)
		bidAmountInput.MustWaitInteractable().MustInput(bidInfo.BidAmount)
	}

	if bidInfo.ProjectDelivery != "" {

		projectDeliveryInput := page.MustElementX(bidInfo.ProjectDeliveryPath)
		projectDeliveryInput.MustWaitInteractable().MustInput(bidInfo.ProjectDelivery)
	}

	projectDescriptionInput := page.MustElementX(bidInfo.ProjectDescriptionPath)
	projectDescriptionInput.MustWaitInteractable().MustInput(bidInfo.ProjectDescription)

	bidButton := page.MustElementX(bidInfo.BidButtonPath)

	bidButton.MustWaitInteractable().MustClick()

	pageURL := page.MustInfo().URL

	fmt.Println(pageURL)

	defer page.MustClose()

	return strings.Contains(pageURL, "details")
}
