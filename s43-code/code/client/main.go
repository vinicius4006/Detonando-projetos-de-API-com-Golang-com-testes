package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "network.golang/curso_grpc/protos"
)

func main() {
	serverAddr := "localhost:8080"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		panic("did not connect")
	}
	defer conn.Close()
	c := pb.NewSignVerifyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	signRequest := pb.SignRequest{Text: "Have a nice day"}

	signResponse, err := c.Sign(ctx, &signRequest)

	fmt.Println("Signature: ", signResponse.Signature)

}
