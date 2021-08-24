package main

import (
	"context"
	example "github.com/mbier/example-grpc/.gen"
	"github.com/pioz/faker"
	"google.golang.org/grpc"
	"log"
	"os/exec"
	"time"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := example.NewPersonServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Save(ctx, &example.CreateRequest{
		Person: &example.Person{
			Name:  faker.FullName(),
			Id:    faker.Int32(),
			Email: faker.Email(),
			Phones: []*example.Person_PhoneNumber{
				{
					Number: faker.PhoneNumber(),
					Type:   example.Person_PhoneType(faker.Int32InRange(0, 2)),
				},
			},
		}},
	)
	if err != nil {
		log.Fatalf("could not save: %v", err)
	}

	resp, err := c.List(ctx, &example.ListRequest{})
	if err != nil {
		log.Fatalf("could not list: %v", err)
	}

	log.Printf("List: %v", resp.GetPerson())

}
