package main

import (
	pb "auto_bidder/RPC/protos"
	"auto_bidder/RPC/scraper"
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
	port = ":50053"
)

type authService struct {
	pb.UnimplementedAuthServer
	IsCookie bool
}

func (s *authService) CheckStatus(ctx context.Context, in *pb.AuthEmpty) (*pb.AuthStatus, error) {

	if f, err := os.Stat("./login_cookie.json"); err == nil && f.Size() > 200 {
		if f != nil {
			s.IsCookie = true
			return &pb.AuthStatus{IsCookie: true}, nil
		}
	} else {
		s.IsCookie = false
		return &pb.AuthStatus{IsCookie: false}, err
	}

	s.IsCookie = false

	return &pb.AuthStatus{IsCookie: false}, nil
}

func (s *authService) Authenticate(ctx context.Context, in *pb.AuthCredentials) (*pb.AuthStatus, error) {

	if s.IsCookie {
		return &pb.AuthStatus{IsCookie: true}, nil
	} else {
		s := scraper.Scraper{}
		s.Start("https://www.freelancer.com/login")
		err := s.Login(scraper.LoginInfo{
			in.Username,
			in.Password,
			scraper.LoginPath{
				UsernameInputPath: in.UsernameInput,
				PasswordInputPath: in.PasswordInput,
				ButtonPath:        in.ButtonInput,
			},
		})
		// scraper.CheckError(s.Browser.Close())

		return &pb.AuthStatus{IsCookie: false}, err
	}

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// s := scraper.Scraper{}
	// s.Start("https://www.freelancer.com/")

	authServer := grpc.NewServer()
	pb.RegisterAuthServer(authServer, &authService{})

	if err := authServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
