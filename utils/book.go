package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetBookByKeyword(keyword string) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", "https://dapi.kakao.com/v3/search/book", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", os.Getenv("KAKAOAK"))

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
	str := string(bytes)
	fmt.Println(str)
}
