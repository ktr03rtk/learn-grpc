package handler

import (
	pb "client/gen/api"
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Handler interface {
	Bake(int) (sring error)
	Report() (string error)
	Close()
}

type handler struct {
	cc     *grpc.ClientConn
	client pb.PancakeBakerServiceClient
	context.Context
}

func NewHandler() (*handler, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	client := pb.NewPancakeBakerServiceClient(conn)

	md := metadata.New(map[string]string{"authorization": "Bearer secretword"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	return &handler{
		cc:      conn,
		client:  client,
		Context: ctx,
	}, nil
}

func (h *handler) Bake(id int) (string, error) {
	req := &pb.BakeRequest{
		Menu: encodePancakeMenu(id),
	}

	r, err := h.client.Bake(h.Context, req)
	if err != nil {
		return "", errors.Wrap(err, "failed to bake")
	}

	return r.Pancake.String(), nil
}

func (h *handler) Report() (string, error) {

	r, err := h.client.Report(h.Context, &pb.ReportRequest{})
	if err != nil {
		return "", errors.Wrap(err, "failed to get report")
	}

	return r.Report.String(), nil
}

func (h *handler) Close() {
	h.cc.Close()
}

func encodePancakeMenu(n int) pb.Pancake_Menu {
	switch n {
	case 1:
		return pb.Pancake_CLASSIC
	case 2:
		return pb.Pancake_BANANA_AND_WHIP
	case 3:
		return pb.Pancake_BACON_AND_CHEESE
	case 4:
		return pb.Pancake_MIX_BERRY
	case 5:
		return pb.Pancake_BAKED_MARSHMALLOW
	case 6:
		return pb.Pancake_SPICY_CURRY
	default:
		return pb.Pancake_UNKNOWN
	}
}
