// Package main implements a server for base64 service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/piyushparsai/base64grpc/proto"
	b64 "encoding/base64"
)

var (
	port = flag.Int("port", 18090, "The base64 service port")
)

// encodeserver is used to implement base64.Base64.
type base64server struct {
	pb.UnimplementedBase64Server
}

// EncodeString implements /base64.Base64/EncodeString
func (s *base64server) EncodeString(ctx context.Context, reqStr *pb.EncodingRequest) (*pb.EncodingResponse, error) {
	log.Printf("Received: %v", reqStr.GetStrToEncode())
	sEnc := b64.StdEncoding.EncodeToString([]byte(reqStr.GetStrToEncode()))
    fmt.Println(sEnc)
	return &pb.EncodingResponse{EncodedStr: sEnc}, nil
}

// DecodeString implements /base64.Base64/DecodeString
func (s *base64server) DecodeString(ctx context.Context, reqStr *pb.DecodingRequest) (*pb.DecodingResponse, error) {
	log.Printf("Received: %v", reqStr.GetStrToDecode())
	sDec, _ := b64.StdEncoding.DecodeString(reqStr.GetStrToDecode())
    fmt.Println(string(sDec))
	return &pb.DecodingResponse{DecodedStr: string(sDec)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBase64Server(s, &base64server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
