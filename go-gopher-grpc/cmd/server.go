/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/scraly/learning-go-by-examples/go-gopher-grpc/pkg/gopher"
	"google.golang.org/grpc"
)

const (
	port         = ":9000"
	KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"
)

// server is used to implement gopher.GopherServer.
type Server struct {
	pb.UnimplementedGopherServer
}

type Gopher struct {
	URL string `json: "url"`
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Schema gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		// Register reflection service on gRPC server.
		reflection.Register(grpcServer)

		// Register services
		pb.RegisterGopherServer(grpcServer, &Server{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

// GetGopher implements gopher.GopherServer
func (s *Server) GetGopher(ctx context.Context, req *pb.GopherRequest) (*pb.GopherRequest, error) {

	// Check request
	if req == nil {
		fmt.Println("request must not be nil")
		return req, xerrors.Errorf("request must not be nil")
	}

	log.Printf("Received: %+v", req)

	return req, nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
