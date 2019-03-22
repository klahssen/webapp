package emails

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"crypto/tls"

	pb "github.com/klahssen/webapp/pkg/domain"
	"gopkg.in/gomail.v2"
)

type srv struct {
	dialer *gomail.Dialer
}

//New returns an instance of an EmailsServer
func New() pb.EmailsServer {
	d := gomail.NewDialer("localhost", 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &srv{dialer: d}
}

func (s *srv) SendAccountConfirm(ctx context.Context, params *pb.EmailParams) (*pb.EmailResp, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "email@domain.com")
	m.SetHeader("To", params.Em)
	m.SetHeader("Subject", "Account Confirmation")
	m.SetBody("text/html", "Hello!</br>Please click <a href='http://www.google.com'>here</a>")
	err := s.dialer.DialAndSend(m)
	if err != nil {
		log.Printf("gomail dialer: dial and send: %v", err)
		return nil, fmt.Errorf("failed to send")
	}
	log.Printf("sent 'account confirmation' email, params %+v", params)
	return &pb.EmailResp{Uid: fmt.Sprintf("%d", rand.Int()), Server: "mock"}, nil
}
func (s *srv) SendEmailAddrConfirm(ctx context.Context, params *pb.EmailParams) (*pb.EmailResp, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "email@domain.com")
	m.SetHeader("To", params.Em)
	m.SetHeader("Subject", "Email Confirmation")
	m.SetBody("text/html", "Hello!</br>Please click <a href='http://www.google.com'>here</a>")
	err := s.dialer.DialAndSend(m)
	if err != nil {
		log.Printf("gomail dialer: dial and send: %v", err)
		return nil, fmt.Errorf("failed to send")
	}
	log.Printf("sent 'email address confirmation' email, params %+v", params)
	return &pb.EmailResp{Uid: fmt.Sprintf("%d", rand.Int()), Server: "mock"}, nil
}
func (s *srv) SendPasswordReset(ctx context.Context, params *pb.EmailParams) (*pb.EmailResp, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "email@domain.com")
	m.SetHeader("To", params.Em)
	m.SetHeader("Subject", "Password Reset")
	m.SetBody("text/html", "Hello!</br>Please click <a href='http://www.google.com'>here</a>")
	err := s.dialer.DialAndSend(m)
	if err != nil {
		log.Printf("gomail dialer: dial and send: %v", err)
		return nil, fmt.Errorf("failed to send")
	}
	log.Printf("sent 'password reset' email, params %+v", params)
	return &pb.EmailResp{Uid: fmt.Sprintf("%d", rand.Int()), Server: "mock"}, nil
}
