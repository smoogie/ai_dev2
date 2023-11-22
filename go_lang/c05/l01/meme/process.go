package meme

/*
Wykonaj zadanie API o nazwie “meme”. Celem zadania jest nauczenie Cię pracy z generatorami grafik i dokumentów.
Zadanie polega na wygenerowaniu mema z podanego obrazka i podanego tekstu.
Mem ma być obrazkiem JPG o wymiarach 1080x1080. Powinien posiadać czarne tło,
dostarczoną grafikę na środku i podpis zawierający dostarczony tekst.
Grafikę możesz wygenerować za pomocą darmowych tokenów dostępnych w usłudze RenderForm (50 pierwszych grafik jest darmowych).
URL do wygenerowanej grafiki spełniającej wymagania wyślij do endpointa /answer/.
W razie jakichkolwiek problemów możesz sprawdzić hinty https://zadania.aidevs.pl/hint/meme
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
)

type respponseC05L01_meme struct {
	Code    int
	Msg     string
	Service string
	Image   string
	Text    string
}

type renderformResponse struct {
	RequestId string
	Href      string
}

func process(body []byte) (string, error) {
	responseJson := respponseC05L01_meme{}
	err := json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	image := responseJson.Image
	text := responseJson.Text
	fmt.Println("response code:", responseJson.Code)
	fmt.Println("response msg:", responseJson.Msg)
	fmt.Println("image msg:", image)
	fmt.Println("text msg:", text)
	return generateMeme(image, text)
}

func generateMeme(image string, text string) (string, error) {
	fmt.Println("---GET DATA---")
	//configure request
	url := "https://get.renderform.io/api/v2/render"
	jsonString := `{"template":"nice-spiders-scrape-madly-1940","data":{`
	jsonString += `"text.text":"` + text + `",`
	jsonString += `"image.src":"` + image + `"`
	jsonString += `}}`
	var jsonData = []byte(jsonString)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("X-API-KEY", os.Getenv("RENDERFROM_KEY"))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	responseJson := renderformResponse{}
	err = json.Unmarshal(body, &responseJson)
	fmt.Println("response href:", responseJson.Href)
	return "\"" + responseJson.Href + "\"", nil
}
