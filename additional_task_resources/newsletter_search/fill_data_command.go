package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	pb "github.com/qdrant/go-client/qdrant"
	"github.com/sashabaranov/go-openai"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Articles struct {
	Articles []Article
}

type Article struct {
	Id    int
	Title string
	Url   string
	Info  string
	Date  string
	Uuid  string
}

const MaxPositions = 300
const StartAt = 297
const CollectionName = "articles"
const VectorSize uint64 = 1536
const Distance = pb.Distance_Dot

func Command(c *cli.Context) error {
	fmt.Println("Reading json file")
	jsonFile, err := os.Open("archiwum.json")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var articles Articles
	err = json.Unmarshal(byteValue, &articles)
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}

	fmt.Println("Open DB connection")
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}

	fmt.Println("Open QDrant connection")
	conn, err := grpc.DialContext(context.Background(), os.Getenv("QDRANT_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		fmt.Println("error:", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	pointsClient := pb.NewPointsClient(conn)

	fmt.Println("Starting adding data")
	endOn, err := strconv.Atoi(os.Getenv("LAST_POSITION"))
	startFrom, err := strconv.Atoi(os.Getenv("START_FROM"))
	for i := startFrom; i < len(articles.Articles) && i < endOn; i++ {
		if i%10 == 0 {
			if i != 0 {
				fmt.Println("Adding positions", i, "-", i+9)
			} else {
				fmt.Println("Adding positions 0 - 9")
			}
		}
		article := articles.Articles[i]
		uuid := uuid.New()
		article.Id = i + 1
		article.Uuid = uuid.String()
		var embedding []float32
		embedding, err = getEmbedding(article.Info)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return err
		}
		err = insertEmbeddingToQdrant(embedding, article, pointsClient, ctx)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return err
		}
		err = insertDataToMySQL(db, article)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return err
		}
	}

	return nil
}

func getEmbedding(textForEmbedding string) ([]float32, error) {
	client := openai.NewClient(os.Getenv("OPEN_AI_KEY"))
	queryReq := openai.EmbeddingRequest{
		Input: []string{textForEmbedding},
		Model: openai.AdaEmbeddingV2,
	}

	// Create an embedding for the user query
	queryResponse, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		fmt.Printf("Embedding error: %v\n", err)
		return make([]float32, 1), err
	}
	queryEmbedding := queryResponse.Data[0].Embedding
	return queryEmbedding, nil
}

func insertEmbeddingToQdrant(embedding []float32, article Article, pointsClient pb.PointsClient, ctx context.Context) error {
	// Upsert points
	waitUpsert := true
	upsertPoints := []*pb.PointStruct{
		{
			// Point Id is number or UUID
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Uuid{article.Uuid},
			},
			Vectors: &pb.Vectors{VectorsOptions: &pb.Vectors_Vector{Vector: &pb.Vector{Data: embedding}}},
			Payload: map[string]*pb.Value{
				"Uuid": {
					Kind: &pb.Value_StringValue{StringValue: article.Uuid},
				},
			},
		},
	}
	_, err := pointsClient.Upsert(ctx, &pb.UpsertPoints{
		CollectionName: CollectionName,
		Wait:           &waitUpsert,
		Points:         upsertPoints,
	})
	return err
}

func insertDataToMySQL(db *sql.DB, article Article) error {

	sql := "INSERT INTO articles(title, url, info, c_date, uuid) VALUES (?, ?, ?, ?, ?)"
	res, err := db.Exec(sql, article.Title, article.Url, article.Info, article.Date, article.Uuid)

	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()

	if err != nil {
		return err
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
	return nil
}
