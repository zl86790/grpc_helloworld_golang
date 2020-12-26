package main

import (
	"context"
	"fmt"
	"helloworld/services"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	client := services.NewHiServiceClient(conn)
	resp, err := client.GetHiResponed(context.Background(), &services.HiRequest{Say: "Hello"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp.Responed)

}
