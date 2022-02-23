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

	deadline := time.Now().AddDate(0, 0, 10)
	todo := &todopb.Todo{
		UserId: "abc123",
		Title:  "buy",
		Detail: "buy a cup of coffee",
		Deadline: &timestamp.Timestamp{
			Seconds: deadline.Unix(),
			Nanos:   int32(deadline.Nanosecond()),
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
	fmt.Println("reading the todo")

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

	// *********************************************
	// * update Todo
	// *********************************************
	fmt.Println("updating the todo")

	newDeadline := time.Now().AddDate(0, 0, 10)
	newTodo := &todopb.Todo{
		Id:     todoId,
		UserId: "changed user",
		Title:  "sell",
		Detail: "sell my car",
		Deadline: &timestamp.Timestamp{
			Seconds: newDeadline.Unix(),
			Nanos:   int32(newDeadline.Nanosecond()),
		},
	}

	updateRes, updateErr := c.UpdateTodo(context.Background(), &todopb.UpdateTodoRequest{Todo: newTodo})
	if updateErr != nil {
		fmt.Printf("failed to update: %v\n", updateErr)
	}

	fmt.Printf("todo was updated: %v\n", updateRes)

	// *********************************************
	// * delete Todo
	// *********************************************
	fmt.Println("deleting the todo")

	deletelRes, deleteErr := c.DeleteTodo(context.Background(), &todopb.DeleteTodoRequest{TodoId: todoId})

	if deleteErr != nil {
		fmt.Printf("failed to delete: %v\n", deleteErr)
	}

	fmt.Printf("todo was deleted: %v\n", deletelRes)
}
