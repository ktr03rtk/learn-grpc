package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ktr03rtk/learn-grpc/mongodb_crud/todopb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mongoDBHandler struct {
	dbClient     *mongo.Client
	dbCollection *mongo.Collection
}

type todoItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   string             `bson:"user_id"`
	Title    string             `bson:"title"`
	Detail   string             `bson:"detail"`
	Deadline time.Time          `bson:"deadline"`
}

func newMongoDBHandler() (*mongoDBHandler, error) {
	fmt.Println("Connecting to MongoDB")

	dbClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:password@localhost:27017"))
	if err != nil {
		return nil, err
	}
	if err := dbClient.Connect(context.Background()); err != nil {
		return nil, err
	}

	dbCollection := dbClient.Database("mydb").Collection("todo")

	return &mongoDBHandler{
		dbClient,
		dbCollection,
	}, nil
}

func (h *mongoDBHandler) disconnect() {
	if err := h.dbClient.Disconnect(context.Background()); err != nil {
		log.Fatalf("faied to disconnect MongoDB %v\n", err)
	}
}

func (h *mongoDBHandler) CreateTodo(ctx context.Context, req *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {
	fmt.Println("create Todo request")
	todo := req.GetTodo()

	data := todoItem{
		UserID:   todo.GetUserId(),
		Title:    todo.GetTitle(),
		Detail:   todo.GetDetail(),
		Deadline: time.Unix(todo.GetDeadline().GetSeconds(), int64(todo.GetDeadline().GetNanos())),
	}

	res, err := h.dbCollection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v\n", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "cannot convert OID")
	}

	return &todopb.CreateTodoResponse{
		Todo: &todopb.Todo{
			Id:     oid.Hex(),
			UserId: todo.GetUserId(),
			Title:  todo.GetTitle(),
			Detail: todo.GetDetail(),
			Deadline: &timestamppb.Timestamp{
				Seconds: todo.GetDeadline().GetSeconds(),
				Nanos:   todo.GetDeadline().GetNanos(),
			},
		},
	}, nil
}

func (h *mongoDBHandler) ReadTodo(ctx context.Context, req *todopb.ReadTodoRequest) (*todopb.ReadTodoResponse, error) {
	fmt.Println("read Todo request")

	todoID := req.GetTodoId()
	oid, err := primitive.ObjectIDFromHex(todoID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse ID")
	}

	data := &todoItem{}
	filter := bson.M{"_id": oid}

	res := h.dbCollection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("cannot find todo with specified ID: %v\n", err))
	}

	return &todopb.ReadTodoResponse{
		Todo: dataToTodoPb(data),
	}, nil
}

func dataToTodoPb(data *todoItem) *todopb.Todo {
	return &todopb.Todo{
		Id:     data.ID.Hex(),
		UserId: data.UserID,
		Title:  data.Title,
		Detail: data.Detail,
		Deadline: &timestamppb.Timestamp{
			Seconds: data.Deadline.Unix(),
			Nanos:   int32(data.Deadline.Nanosecond()),
		},
	}
}

func (h *mongoDBHandler) UpdateTodo(ctx context.Context, req *todopb.UpdateTodoRequest) (*todopb.UpdateTodoResponse, error) {
	fmt.Println("update Todo request")

	todo := req.GetTodo()
	oid, err := primitive.ObjectIDFromHex(todo.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse ID")
	}

	data := &todoItem{}
	filter := bson.M{"_id": oid}

	res := h.dbCollection.FindOne(ctx, filter)
	if err = res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find todo with specified ID: %v", err)
	}

	data.UserID = todo.GetUserId()
	data.Title = todo.GetTitle()
	data.Detail = todo.GetDetail()
	data.Deadline = time.Unix(todo.GetDeadline().GetSeconds(), int64(todo.GetDeadline().GetNanos()))

	_, updateErr := h.dbCollection.ReplaceOne(ctx, filter, data)
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot update object in MongoDB: %v", updateErr))
	}

	return &todopb.UpdateTodoResponse{
		Todo: dataToTodoPb(data),
	}, nil
}

func (h *mongoDBHandler) DeleteTodo(ctx context.Context, req *todopb.DeleteTodoRequest) (*todopb.DeleteTodoResponse, error) {
	fmt.Println("delete Todo request")

	oid, err := primitive.ObjectIDFromHex(req.GetTodoId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse ID")
	}

	filter := bson.M{"_id": oid}

	res, err := h.dbCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot delete object in MongoDB: %v", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "cannot find todo in MongoDB")
	}

	return &todopb.DeleteTodoResponse{TodoId: req.GetTodoId()}, nil
}
