package main

import (
	pb "auto_bidder/RPC/protos"
	"auto_bidder/RPC/scraper"
	"context"

	_ "embed"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
	port = ":50052"
)

type projectService struct {
	pb.UnimplementedProjectServer
	crawler *scraper.Scraper
}

func (p *projectService) GetProjects(pEmpty *pb.ProjectEmpty, stream pb.Project_GetProjectsServer) error {

	c := make(chan bool)

	stop := p.crawler.GetProjects(func(projects []*pb.Projects) {
		for _, project := range projects {
			if err := stream.Send(project); err != nil {
				fmt.Println(err)
			}
		}

		c <- true
	})

	<-c

	stop()

	return nil
}

func (p *projectService) SubscribeToProject(pEmpty *pb.ProjectEmpty, stream pb.Project_SubscribeToProjectServer) error {
	c := make(chan bool)

	p.crawler.SubscribeToProject(func(projects *pb.Projects) {
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
	projectService.crawler.Start("https://freelancer.com")
	pb.RegisterProjectServer(projectServer, projectService)

	if err := projectServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
