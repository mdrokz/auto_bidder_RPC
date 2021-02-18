package scraper

// LoginInfo ...
type LoginInfo struct {
	Username string
	Password string
	LoginPath
}

// LoginPath ...
type LoginPath struct {
	UsernameInputPath string
	PasswordInputPath string
	ButtonPath        string
}

// BidInfo ...
type BidInfo struct {
	BidAmount          string
	ProjectDelivery    string
	ProjectDescription string
	BidPath
}

// BidPath ...
type BidPath struct {
	BidAmountPath          string
	ProjectDeliveryPath    string
	ProjectDescriptionPath string
	BidButtonPath          string
}

// Project ...
type Project struct {
	Skills       string
	Title        string
	Link         string
	Description  string
	BiddingPrice string
	Selected     bool
	Currency     string
}
