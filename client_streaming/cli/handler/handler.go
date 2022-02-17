package handler

import (
	"client/gen/pb"
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Handler interface {
	SendMetaData(string) error
	SendData([]byte) error
	CloseAndRecv() (string, error)
	Close()
}

type handler struct {
	cc     *grpc.ClientConn
	client pb.ImageUploadService_UploadClient
}

func NewHandler() (Handler, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	c, err := pb.NewImageUploadServiceClient(conn).Upload(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client")
	}

	return &handler{
		cc:     conn,
		client: c,
	}, nil
}

func (h *handler) SendMetaData(fileName string) error {
	if err := h.client.Send(&pb.ImageUploadRequest{
		File: &pb.ImageUploadRequest_FileMeta_{
			FileMeta: &pb.ImageUploadRequest_FileMeta{
				Filename: fileName,
			},
		},
	}); err != nil {
		return errors.Wrap(err, "failed to send filename")
	}

	return nil
}

func (h *handler) SendData(data []byte) error {
	if err := h.client.Send(&pb.ImageUploadRequest{
		File: &pb.ImageUploadRequest_Data{
			Data: data,
		},
	}); err != nil {
		return errors.Wrap(err, "failed to send data")
	}

	return nil
}

func (h *handler) CloseAndRecv() (string, error) {
	res, err := h.client.CloseAndRecv()
	if err != nil {
		return "", errors.Wrap(err, "failed to receive responce")
	}

	return res.String(), nil
}

func (h *handler) Close() {
	h.cc.Close()
}
