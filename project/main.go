package main

import (
	pb "auto_bidder/RPC/protos"
	"auto_bidder/RPC/scraper"
	"context"
	"fmt"
	"unsafe"

	_ "embed"
	"log"
	"net"

	"github.com/mattn/go-pointer"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

/*
#cgo CFLAGS: -I/usr/include/glib-2.0 -I/usr/lib/x86_64-linux-gnu/glib-2.0/include/ -I/usr/include/gdk-pixbuf-2.0/
#cgo LDFLAGS: libnotification.a -lnotify -lglib-2.0

#include <libnotify/notify.h>
#include <stdio.h>
#include "lib.h"
*/
import "C"

const (
	port = ":50052"
)

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

type projectService struct {
	pb.UnimplementedProjectServer
	crawler *scraper.Scraper
}

func (p *projectService) GetProjects(pInfo *pb.ProjectPathInfo, stream pb.Project_GetProjectsServer) error {

	c := make(chan bool)

	var pr []*pb.Projects

	p.crawler.BidPath = scraper.BidPath{
		BidAmountPath:          pInfo.BidAmountPath,
		ProjectDeliveryPath:    pInfo.ProjectDeliveryPath,
		ProjectDescriptionPath: pInfo.ProjectDescriptionPath,
		BidButtonPath:          pInfo.BidButtonPath,
	}

	stop := p.crawler.GetProjects(func(projects []*pb.Projects) {
		pr = projects
		// for _, project := range projects {
		// 	if err := stream.Send(project); err != nil {
		// 		fmt.Println(err)
		// 	}
		// }

		c <- true
	})

	<-c

	for _, project := range pr {
		if err := stream.Send(project); err != nil {
			fmt.Println(err)
		}
	}

	stop()

	return nil
}

func (p *projectService) SubscribeToProject(pEmpty *pb.ProjectEmpty, stream pb.Project_SubscribeToProjectServer) error {
	c := make(chan bool)

	p.crawler.SubscribeToProject(func(projects *pb.Projects) {
		C._register_callback(pointer.Save(&Callback{
			Func: notificationCallback,
			UserData: CallBackInfo{
				Link: projects.Link,
				BidInfo: scraper.BidInfo{
					BidAmount:          "",
					ProjectDelivery:    "",
					ProjectDescription: "Hello i am a full stack developer with 5+ years of experience in Web Development and i have experience in MERN and MEAN Stack, i will be able to accomplish your task.",
					BidPath:            p.crawler.BidPath,
				},
				Crawler: p.crawler,
			},
		}))

		n := C.notificationInfo{
			title:        C.CString(projects.Title),
			body:         C.CString(projects.Description),
			link:         C.CString(projects.Link),
			button_name:  C.CString("Bid"),
			loop_timeout: 3500,
		}

		C.send_notification(n)

		stream.Send(projects)
	})

	<-c

	return nil
}

func (p *projectService) BidOnProject(ctx context.Context, projectInfo *pb.ProjectInfo) (*pb.ProjectStatus, error) {

	bidInfo := scraper.BidInfo{
		BidAmount:          projectInfo.BidAmount,
		ProjectDelivery:    projectInfo.ProjectDelivery,
		ProjectDescription: projectInfo.ProjectDescription,
		BidPath: scraper.BidPath{
			BidAmountPath:          projectInfo.BidAmountPath,
			ProjectDeliveryPath:    projectInfo.ProjectDeliveryPath,
			ProjectDescriptionPath: projectInfo.ProjectDescriptionPath,
			BidButtonPath:          projectInfo.BidButtonPath,
		},
	}

	status := p.crawler.BidOnProject(projectInfo.Link, bidInfo)

	return &pb.ProjectStatus{Status: status}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	projectServer := grpc.NewServer()
	projectService := &projectService{crawler: &scraper.Scraper{}}
	projectService.crawler.Start("https://freelancer.com", false)
	pb.RegisterProjectServer(projectServer, projectService)

	if err := projectServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//export cb_proxy
func cb_proxy(v unsafe.Pointer) {
	cb := pointer.Restore(v).(*Callback)
	cb.Func(cb.UserData)
}
