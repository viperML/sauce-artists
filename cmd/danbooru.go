package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Response struct {
	Artist string `json:"tag_string_artist"`
}

func GetPost(config *Config, postId int64) Response {
	url, err := url.Parse("https://danbooru.donmai.us/posts/" + strconv.Itoa(int(postId)) + ".json")
	if err != nil {
		log.Fatal(err)
	}
	values := url.Query()
	values.Add("login", config.username)
	values.Add("api_key", config.api_key)
	url.RawQuery = values.Encode()
	log.Print(url.String())

	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "curl/8.9.1")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)

	ct := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	if ct != "application/json" {
		log.Fatalf("Content-Type was <%s>", ct)
	}


	bytes, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(string(bytes))

    var res Response
    json.Unmarshal(bytes, &res)
    return res
}
