package main

import (
	"client/handler"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	fileName   = "cat.jpg"
	bufferSize = 1024
)

func main() {
	fmt.Println("Image upload start")

	h, err := handler.NewHandler()
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	if err := h.SendMetaData(fileName); err != nil {
		log.Fatal(err)
	}

	fp, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer fp.Close()

	buf := make([]byte, bufferSize)

	for {
		_, err := fp.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if err := h.SendData(buf); err != nil {
			log.Fatal(err)
		}
	}

	res, err := h.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("--------------- %+v\n", res)

	fmt.Println("Image upload success")
}
