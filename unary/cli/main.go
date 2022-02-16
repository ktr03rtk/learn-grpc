package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "client/gen/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPancakeBakerServiceClient(conn)

	md := metadata.New(map[string]string{"authorization": "Bearer secretword"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	fmt.Println("Pancake client start")

	scanner := bufio.NewScanner(os.Stdin)

Loop:
	for {
		fmt.Println("1: Order pancake")
		fmt.Println("2: Inquire status")
		fmt.Println("3: Exit")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			fmt.Println("Select a Pancake")
			fmt.Println("1: CLASSIC")
			fmt.Println("2: BANANA_AND_WHIP")
			fmt.Println("3: BACON_AND_CHEESE")
			fmt.Println("4: MIX_BERRY")
			fmt.Println("5: BAKED_MARSHMALLOW")
			fmt.Println("6: SPICY_CURRY")

			scanner.Scan()
			in := scanner.Text()

			switch in {
			case "1", "2", "3", "4", "5", "6":
				menu, err := strconv.Atoi(in)
				if err != nil {
					log.Fatalf("failed to strconv: %v", err)
				}

				req := &pb.BakeRequest{
					Menu: encodePancakeMenu(menu),
				}

				r, err := c.Bake(ctx, req)
				if err != nil {
					log.Fatalf("failed to bake: %v", err)
				}

				fmt.Println("Order accepted")
				fmt.Printf("%+v\n", r)
			default:
				fmt.Println("invalid command")
				continue
			}
			continue
		case "2":
			r, err := c.Report(ctx, &pb.ReportRequest{})
			if err != nil {
				log.Fatalf("failed to get report: %v", err)
			}
			fmt.Printf("Order status: %v\n", r)
			continue
		case "3":
			fmt.Println("bye")
			break Loop
		default:
			fmt.Println("invalid command")
			continue
		}
	}
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
