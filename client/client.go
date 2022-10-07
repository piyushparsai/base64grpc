package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/piyushparsai/base64grpc/proto"
)

var (
	addr = flag.String("addr", "localhost:18090", "the server address")
	plainstring = flag.String("str", "Piyush Parsai", "string to encode")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewBase64Client(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	re, err := client.EncodeString(ctx, &pb.EncodingRequest{StrToEncode: *plainstring})
	if err != nil {
		log.Fatalf("could not encode: %v", err)
	}
	log.Printf("Encoded String: %s", re.GetEncodedStr())

	log.Printf("Decode the string to get original string.")

	rd, err := client.DecodeString(ctx, &pb.DecodingRequest{StrToDecode: re.GetEncodedStr()})
	if err != nil {
		log.Fatalf("could not decode: %v", err)
	}
	log.Printf("Decoded String: %s", rd.GetDecodedStr())
}