package appgrpc

import (
	"flag"
	"fmt"
	"github.com/200Lab-Education/go-sdk/logger"
	"google.golang.org/grpc"
	"lesson-5-goland/common"
	"net"
)

type grpcServer struct {
	prefix      string
	port        int
	logger      logger.Logger
	server      *grpc.Server
	grpcHandler func(*grpc.Server)
}

func NewGrpcServer(prefix string) *grpcServer {
	return &grpcServer{
		prefix: prefix,
	}
}

func (gs *grpcServer) SetGrpcHandler(grpcHandler func(*grpc.Server)) {
	gs.grpcHandler = grpcHandler
}

func (gs *grpcServer) GetPrefix() string {
	return gs.prefix
}

func (gs *grpcServer) Get() interface{} {
	return gs
}

func (gs *grpcServer) Name() string {
	return gs.prefix
}

func (gs *grpcServer) InitFlags() {
	flag.IntVar(&gs.port, gs.prefix+"-port", 50051, "Port of grpc server")
}

func (gs *grpcServer) Configure() error {
	gs.logger = logger.GetCurrent().GetLogger(gs.prefix)
	gs.logger.Infoln("Setup gRPC service:", gs.prefix)
	gs.server = grpc.NewServer()

	return nil
}

func (gs *grpcServer) Run() error {
	_ = gs.Configure()

	go func() {
		defer common.AppRecover()
		if gs.grpcHandler != nil {
			gs.logger.Infoln("registering services...")
			gs.grpcHandler(gs.server)
		}

		address := fmt.Sprintf("0.0.0.0:%d", gs.port)
		lis, err := net.Listen("tcp", address)

		if err != nil {
			gs.logger.Errorln("Error %v", err)
		}

		gs.server.Serve(lis)

	}()

	return nil
}

func (gs *grpcServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		gs.server.Stop()
		c <- true
	}()
	return c
}
