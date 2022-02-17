package handler

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"image.upload/gen/pb"
)

type ImageUploadHandler struct {
	sync.Mutex
	files map[string][]byte
}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{
		files: make(map[string][]byte),
	}
}

const outputDir = "output"

func (h *ImageUploadHandler) Upload(stream pb.ImageUploadService_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	meta := req.GetFileMeta()
	fileName := meta.Filename

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	uuid := u.String()
	buf := &bytes.Buffer{}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(outputDir, fileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}

		if _, err := f.Write(chunk); err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)

	h.files[uuid] = data

	err = stream.SendAndClose(&pb.ImageUploadResponse{
		Uuid:        uuid,
		Size:        int32(len(data)),
		Filename:    fileName,
		ContentType: mimeType,
	})

	return err
}
