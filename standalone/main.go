package main

import (
	pb "auto_bidder/RPC/protos"
	"auto_bidder/RPC/scraper"
	"unsafe"

	"github.com/mattn/go-pointer"
)

/*
#cgo CFLAGS: -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/ -I/usr/include/gdk-pixbuf-2.0/
#cgo LDFLAGS: libnotification.a -lnotify -lglib-2.0

#include <libnotify/notify.h>
#include <stdio.h>
#include "lib.h"
*/
import "C"

type CallBackInfo struct {
	Link    string
	BidInfo scraper.BidInfo
	Crawler *scraper.Scraper
}

type Callback struct {
	Func     func(CallBackInfo)
	UserData CallBackInfo
}

func notificationCallback(c CallBackInfo) {
	c.Crawler.BidOnProject(c.Link, c.BidInfo)
}

func main() {

	c := make(chan bool)

	crawler := &scraper.Scraper{
		BidPath: scraper.BidPath{
			BidAmountPath:          "/html/body/app-root/app-logged-in-shell/div/div[2]/ng-component/app-project-view-details/fl-page-layout/main/fl-container/fl-page-layout-single/fl-grid/fl-col[1]/app-project-details-freelancer/app-bid-form/fl-card/fl-bit/fl-bit[2]/fl-bit[1]/fl-bit/fl-bit[1]/fl-input/fl-bit/fl-bit[2]/input",
			ProjectDeliveryPath:    "/html/body/app-root/app-logged-in-shell/div/div[2]/ng-component/app-project-view-details/fl-page-layout/main/fl-container/fl-page-layout-single/fl-grid/fl-col[1]/app-project-details-freelancer/app-bid-form/fl-card/fl-bit/fl-bit[2]/fl-bit[1]/fl-bit/fl-bit[2]/fl-input/fl-bit/fl-bit[1]/input",
			ProjectDescriptionPath: "/html/body/app-root/app-logged-in-shell/div/div[2]/ng-component/app-project-view-details/fl-page-layout/main/fl-container/fl-page-layout-single/fl-grid/fl-col[1]/app-project-details-freelancer/app-bid-form/fl-card/fl-bit/fl-bit[2]/fl-bit[1]/fl-textarea/textarea",
			BidButtonPath:          "/html/body/app-root/app-logged-in-shell/div/div[2]/ng-component/app-project-view-details/fl-page-layout/main/fl-container/fl-page-layout-single/fl-grid/fl-col[1]/app-project-details-freelancer/app-bid-form/fl-card/fl-bit/fl-bit[2]/fl-bit[2]/fl-button/button",
		},
	}

	// _, e := os.Open("/usr/bin/login_cookie.json")

	// fmt.Println(e)

	crawler.Start("https://freelancer.com", true)

	crawler.SubscribeToProject(func(project *pb.Projects) {
		C._register_callback(pointer.Save(&Callback{
			Func: notificationCallback,
			UserData: CallBackInfo{
				Link: project.Link,
				BidInfo: scraper.BidInfo{
					BidAmount:          "",
					ProjectDelivery:    "",
					ProjectDescription: "Hello i am a full stack developer with 5+ years of experience in Web Development and i have experience in MERN and MEAN Stack, i will be able to accomplish your task.",
					BidPath:            crawler.BidPath,
				},
				Crawler: crawler,
			},
		}))

		n := C.notificationInfo{
			title:        C.CString(project.Title),
			body:         C.CString(project.Description),
			link:         C.CString(project.Link),
			button_name:  C.CString("Bid"),
			loop_timeout: 3500,
		}

		C.send_notification(n)
	})

	<-c
}

//export cb_proxy
func cb_proxy(v unsafe.Pointer) {
	cb := pointer.Restore(v).(*Callback)
	cb.Func(cb.UserData)
}
