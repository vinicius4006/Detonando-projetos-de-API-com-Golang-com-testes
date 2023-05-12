package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
	pb "network.golang/curso_grpc/protos"
)

type server struct {
	pb.UnimplementedSignVerifyServer
}

var (
	privateKey *rsa.PrivateKey
)

func (s *server) Sign(ctx context.Context, input *pb.SignRequest) (*pb.SignResponse, error) {

	hashed := sha256.Sum256([]byte(input.Text))

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	output := hex.EncodeToString(signature)
	return &pb.SignResponse{Signature: output}, nil

}

// Get the private key:

func GetPrivateKey() (*rsa.PrivateKey, error) {
	key, err := ioutil.ReadFile("../private_key.pem")
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)
	der, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return der, err
}

func main() {
	var err error
	privateKey, err = GetPrivateKey()
	if err != nil {
		panic(err)
	}
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("\nfailed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSignVerifyServer(s, &server{})
	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
