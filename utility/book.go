package utility

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CreateRequest(url string) *http.Request {
	if url == "" {
		url = "https://dapi.kakao.com/v3/search/book"
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env파일을 찾을 수 없습니다!")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("request 객체를 생성하지 못했습니다.")
	}

	req.Header.Add("Authorization", os.Getenv("KAKAOAK"))

	return req
}

func GetBookByKeyword(keyword string) map[string]interface{} {
	req := CreateRequest("")

	q := req.URL.Query()
	q.Add("size", "10")
	q.Add("query", keyword)
	q.Add("target", "title")

	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bytes, err := io.ReadAll(res.Body)
	var jsonMap map[string]interface{}
	json.Unmarshal(bytes, &jsonMap)

	return jsonMap
}
