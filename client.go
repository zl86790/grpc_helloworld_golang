package main

import (
	"context"
	"fmt"
	"helloworld/services"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	args := os.Args
	url := args[1]
	conn, err := grpc.Dial(url+":30010", grpc.WithInsecure())
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
