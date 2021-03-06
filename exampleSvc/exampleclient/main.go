package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {
	// cred := credentials.NewClientTLSFromCert(&tls.Certificate{}, "")
	cred, err := credentials.NewClientTLSFromFile("../server.crt", "")
	if err != nil {
		fmt.Println("unexpected err", err)
		os.Exit(1)
	}

	cc, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(cred))
	if err != nil {
		fmt.Println("unexpected err", err)
		os.Exit(1)
	}

	client := NewExampleServiceClient(cc)

	resp, err := client.SayHello(context.Background(), &SayRequest{
		SayWhat: "this is an example client",
	})
	if err != nil {
		fmt.Println("unexpected err", err)
		os.Exit(1)
	}

	fmt.Println("resp:", resp.SaidWhat)

	_, err = client.ReturnsUnimplementedCode(context.Background(), &SayRequest{})
	if err != nil {
		status, _ := status.FromError(err)
		fmt.Printf("ReturnsUnimplementedCode returned code %d with message %s", status.Code(), status.Message())
	}

}
