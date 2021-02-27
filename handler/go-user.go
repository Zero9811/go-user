package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	gouser "go-user/proto"
)

type GoUser struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GoUser) Call(ctx context.Context, req *gouser.Request, rsp *gouser.Response) error {
	log.Info("Received GoUser.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GoUser) Stream(ctx context.Context, req *gouser.StreamingRequest, stream gouser.GoUser_StreamStream) error {
	log.Infof("Received GoUser.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&gouser.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GoUser) PingPong(ctx context.Context, stream gouser.GoUser_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&gouser.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
