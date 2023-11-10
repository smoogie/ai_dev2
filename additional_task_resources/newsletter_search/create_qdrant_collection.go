package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/qdrant/go-client/qdrant"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func CreateQdrantCollection(c *cli.Context) error {
	fmt.Println("Start collection configuration")

	// Set up a connection to the server.
	conn, err := grpc.DialContext(context.Background(), os.Getenv("QDRANT_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect: ", err)
		return err
	}
	defer conn.Close()
	// create grpc collection client
	collections_client := pb.NewCollectionsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var defaultSegmentNumber uint64 = 2
	_, err = collections_client.Create(ctx, &pb.CreateCollection{
		CollectionName: CollectionName,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     VectorSize,
				Distance: Distance,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &defaultSegmentNumber,
		},
	})
	if err != nil {
		fmt.Println("Could not create collection:", err)
		return err
	} else {
		fmt.Println("Collection", CollectionName, "created")
	}

	return nil
}
