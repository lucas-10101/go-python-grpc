package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/lucas-10101/training/go-grpc/pratice/proto"
	pb "github.com/lucas-10101/training/go-grpc/pratice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var ERR_UNSUPPORTED_VERSION = errors.New("unknown version")

type Server struct {
	pb.UnimplementedUUIDServiceServer
}

func (s *Server) GetUUID(ctx context.Context, req *proto.UUIDRequest) (*proto.UUIDResponse, error) {

	var generated uuid.UUID
	var err error

	switch req.Version {
	case 4:
		generated, err = uuid.NewUUID()
	case 5:
		namespace, err := uuid.Parse(req.Namespace)
		if err != nil {
			return nil, err
		}
		if namespace == uuid.Nil {
			return nil, errors.New("invalid namespace")
		} else if req.ValueToHash == "" {
			return nil, errors.New("value to hash is required")
		}
		generated = uuid.NewSHA1(namespace, []byte(req.ValueToHash))
	case 6:
		if generated, err = uuid.NewV6(); err != nil {
			return nil, err
		}
	case 7:
		if generated, err = uuid.NewV7(); err != nil {
			return nil, err
		}
	default:
		return nil, ERR_UNSUPPORTED_VERSION
	}

	if err != nil {
		return nil, err
	}

	value := generated.String()
	if value == "" {
		return nil, errors.New("failed to generate UUID")
	}

	return &proto.UUIDResponse{Uuid: value}, nil
}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		panic(err)
	}

	tls, err := credentials.NewServerTLSFromFile("../server.crt", "../server.key")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer(grpc.Creds(tls))
	server.RegisterService(&proto.UUIDService_ServiceDesc, &Server{})

	log.Printf("Listening on %s\n", listener.Addr().String())
	server.Serve(listener)

}
