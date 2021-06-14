package cuckoorpc

import (
	context "context"
	"log"
	"net"
	"strconv"

	"CuckooGo/internal/filter"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CuckooServer struct {
	ft *filter.Filter
}

func (s *CuckooServer) Lookup(ctx context.Context, in *RequestData) (*ReplyData, error) {
	res := s.ft.Lookup(in.Data)
	return &ReplyData{Data: res}, nil
}

func (s *CuckooServer) Insert(ctx context.Context, in *RequestData) (*ReplyData, error) {
	res := s.ft.InsertUnique(in.Data)
	return &ReplyData{Data: res}, nil
}

func (s *CuckooServer) Delete(ctx context.Context, in *RequestData) (*ReplyData, error) {
	res := s.ft.Delete(in.Data)
	return &ReplyData{Data: res}, nil
}

func (s *CuckooServer) Reset(ctx context.Context, in *NullMessage) (*NullMessage, error) {
	s.ft.Reset()
	return &NullMessage{}, nil
}

func (s *CuckooServer) Count(ctx context.Context, in *NullMessage) (*ReplyUint, error) {
	res := s.ft.Count()
	return &ReplyUint{Data: uint64(res)}, nil
}

type CuckooRpcServer struct {
	Port uint
	ft   *filter.Filter
}

func RpcServer(port uint, ft *filter.Filter) *CuckooRpcServer {
	return &CuckooRpcServer{
		Port: port,
		ft:   ft,
	}
}

func (self *CuckooRpcServer) Listen() {
	var err error
	port := strconv.FormatUint(uint64(self.Port), 10)
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	RegisterCuckooRpcServiceServer(server, &CuckooServer{ft: self.ft})
	reflection.Register(server)
	log.Printf("RPC Listen at :%s", port)
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
	defer self.ft.SaveFile()
}
