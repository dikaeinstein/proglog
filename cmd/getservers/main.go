package main

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"

	api "github.com/dikaeinstein/proglog/api/v1"
)

func main() {
	addr := flag.String("addr", ":8400", "service address")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := api.NewLogClient(conn)
	r, err := c.GetServers(context.Background(), &api.GetServersRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("servers:")
	for _, server := range r.Servers {
		log.Printf("\t- %v\n", server)
	}
}
