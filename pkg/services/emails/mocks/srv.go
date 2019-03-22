package emails

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "github.com/klahssen/webapp/pkg/domain"
)

type srv struct{}

//New returns an instance of an EmailsServer
func New() pb.EmailsServer {
	return &srv{}
}

func (s *srv) SendAccountConfirm(ctx context.Context, params *pb.EmailParams) (*pb.EmailResp, error) {
	log.Printf("sent 'account confirmation' email, params %+v", params)
	return &pb.EmailResp{Uid: fmt.Sprintf("%d", rand.Int()), Server: "mock"}, nil
}
func (s *srv) SendEmailAddrConfirm(ctx context.Context, params *pb.EmailParams) (*pb.EmailResp, error) {
	log.Printf("sent 'email address confirmation' email, params %+v", params)
	return &pb.EmailResp{Uid: fmt.Sprintf("%d", rand.Int()), Server: "mock"}, nil
}
func (s *srv) SendPasswordReset(ctx context.Context, params *pb.EmailParams) (*pb.EmailResp, error) {
	log.Printf("sent 'password reset' email, params %+v", params)
	return &pb.EmailResp{Uid: fmt.Sprintf("%d", rand.Int()), Server: "mock"}, nil
}
