package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"client/handler"
)

func main() {
	fmt.Println("Pancake client start")

	h, err := handler.NewHandler()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

Loop:
	for {
		in := selectMode(scanner)

		switch in {
		case "1":
			id, err := selectPancake(scanner)
			if err != nil {
				log.Fatalf("failed to select Pancake: %v", err)
			}

			switch id {
			case 1, 2, 3, 4, 5, 6:

				r, err := h.Bake(id)
				if err != nil {
					log.Fatalf("failed to bake: %v", err)
				}

				fmt.Println("Order accepted")
				fmt.Printf("%+v\n", r)
			default:
				fmt.Println("invalid command")
			}
			continue
		case "2":
			r, err := h.Report()
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
		}
	}
}

func selectMode(scanner *bufio.Scanner) string {
	fmt.Println("1: Order pancake")
	fmt.Println("2: Inquire status")
	fmt.Println("3: Exit")

	scanner.Scan()
	in := scanner.Text()

	return in
}

func selectPancake(scanner *bufio.Scanner) (int, error) {
	fmt.Println("Select a Pancake")
	fmt.Println("1: CLASSIC")
	fmt.Println("2: BANANA_AND_WHIP")
	fmt.Println("3: BACON_AND_CHEESE")
	fmt.Println("4: MIX_BERRY")
	fmt.Println("5: BAKED_MARSHMALLOW")
	fmt.Println("6: SPICY_CURRY")

	scanner.Scan()
	in := scanner.Text()

	id, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}

	return id, nil
}
