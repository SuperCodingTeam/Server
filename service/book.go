package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func createRequest(url string) *http.Request {
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

func checkTarget(target string) bool {
	for _, data := range [4]string{"title", "isbn", "publisher", "person"} {
		if data == target {
			return true
		}
	}

	return false
}

func GetBookByKeyword(query string, target string) []map[string]interface{} {
	if !checkTarget(target) {
		return nil
	}

	req := createRequest("")

	q := req.URL.Query()
	q.Add("size", "10")
	q.Add("query", query)
	q.Add("target", target)

	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(bytes, &jsonMap); err != nil {
		panic(err)
	}

	documents, ok := jsonMap["documents"].([]interface{})
	if !ok {
		log.Println("documents 키를 찾을 수 없거나 올바른 타입이 아닙니다.")
		return nil
	}

	var results []map[string]interface{}
	for _, doc := range documents {
		book, ok := doc.(map[string]interface{})
		if !ok {
			continue
		}

		results = append(results, map[string]interface{}{
			"authors":     book["authors"],
			"contents":    book["contents"],
			"isbn":        book["isbn"],
			"publisher":   book["publisher"],
			"thumbnail":   book["thumbnail"],
			"title":       book["title"],
			"translators": book["translators"],
		})
	}

	return results
}
