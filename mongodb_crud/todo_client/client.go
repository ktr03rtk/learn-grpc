package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ktr03rtk/learn-grpc/mongodb_crud/todopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("TODO client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("failed to connect %v\n", err)
	}
	defer cc.Close()

	c := todopb.NewTodoServiceClient(cc)

	// *********************************************
	// * create Todo
	// *********************************************
	fmt.Println("creating the todo")

	now := time.Now()
	todo := &todopb.Todo{
		UserId: "abc123",
		Title:  "buy",
		Detail: "buy a cup of coffee",
		Deadline: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	createTodoRes, err := c.CreateTodo(context.Background(), &todopb.CreateTodoRequest{
		Todo: todo,
	})
	if err != nil {
		log.Fatalf("failed to create todo: %v\n", err)

	}

	fmt.Printf("todo has been created: %v\n", createTodoRes)

	todoId := createTodoRes.GetTodo().GetId()

	// *********************************************
	// * read Todo
	// *********************************************
	fmt.Println("read the todo")

	_, readErr := c.ReadTodo(context.Background(), &todopb.ReadTodoRequest{
		TodoId: "hogehoge",
	})
	if readErr != nil {
		fmt.Printf("failed to read todo: %v\n", readErr)
	}

	readTodoRes, err := c.ReadTodo(context.Background(), &todopb.ReadTodoRequest{
		TodoId: todoId,
	})
	if err != nil {
		fmt.Printf("failed to read todo: %v\n", err)
	}

	fmt.Printf("todo was read: %v\n", readTodoRes)
}
