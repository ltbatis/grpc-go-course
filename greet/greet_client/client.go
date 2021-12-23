package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ltbatista/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, i'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Sarting to do an Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Lucas",
			LastName:  "Batista",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC:  %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
