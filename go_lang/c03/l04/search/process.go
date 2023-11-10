package search

/*
Dziś zadanie jest proste, ale nie łatwe —
zaimportuj do swojej bazy wektorowej, spis wszystkich linków z newslettera unknowNews z adresu:
https://unknow.news/archiwum.json
[jeśli zależy Ci na czasie, możesz dodać pierwsze 300 rekordów]

Następnie wykonaj zadanie API o nazwie “search” — odpowiedz w nim na zwrócone przez API pytanie.
Odpowiedź musi być adresem URL kierującym do jednego z linków unknowNews. Powodzenia!
*/
import (
	"ai_dev/open_ai_help"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

type respponseC03L04_functions struct {
	Code     int
	Msg      string
	Question string
}

type Article struct {
	Id     int
	Title  string
	Url    string
	Info   string
	C_date string
	Uuid   string
}

func process(body []byte) (string, error) {
	responseJson := respponseC03L04_functions{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	question := responseJson.Question
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("question msg:", question)
	return findLink(question)
}

func findLink(question string) (string, error) {
	questionEmbedding, err := getEmbeddingForQuestion(question)
	if err != nil {
		return "", err
	}
	article, err := getArticle(questionEmbedding)
	if err != nil {
		return "", err
	}
	return "\"" + article.Url + "\"", nil
}

func getEmbeddingForQuestion(question string) ([]float32, error) {
	fmt.Println("OPEN AI BUILD EMBEDDING")
	response, err := open_ai_help.SendEmbeddingRequest(question)
	return response, err
}

func getArticle(questionEmbedding []float32) (Article, error) {
	uuid, err := searchSimilarEmbedding(questionEmbedding)
	if err != nil {
		return Article{}, err
	}
	return getArticleByUuId(uuid)
}

func searchSimilarEmbedding(questionEmbedding []float32) (string, error) {
	//TODO: FIND SIMILAR IN QDRANT

	fmt.Println("Open QDrant connection")
	conn, err := grpc.DialContext(context.Background(), os.Getenv("QDRANT_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	pointsClient := pb.NewPointsClient(conn)

	fmt.Println("Search in QDrant")
	// Unfiltered search
	unfilteredSearchResult, err := pointsClient.Search(ctx, &pb.SearchPoints{
		CollectionName: "articles",
		Vector:         questionEmbedding,
		Limit:          1,
		// Include all payload and vectors in the search result
		WithVectors: &pb.WithVectorsSelector{SelectorOptions: &pb.WithVectorsSelector_Enable{Enable: true}},
		WithPayload: &pb.WithPayloadSelector{SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true}},
	})
	if err != nil {
		return "", err
	}
	results := unfilteredSearchResult.GetResult()
	result := results[0]
	fmt.Println("Found result with uuid:", result.Id.GetUuid())
	return result.Id.GetUuid(), err
}

func getArticleByUuId(uuid string) (Article, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB"))
	defer db.Close()

	if err != nil {
		return Article{}, err
	}
	row := db.QueryRow("SELECT * FROM articles WHERE uuid=?", uuid)
	var article Article
	err = row.Scan(&article.Id, &article.Title, &article.Url, &article.Info, &article.C_date, &article.Uuid)
	if err != nil {
		return Article{}, err
	}
	return article, nil
}
