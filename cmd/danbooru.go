package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.uber.org/ratelimit"
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

	ct := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	if ct != "application/json" {
		log.Fatalf("Content-Type was <%s>", ct)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var res Response
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func CollectAuthors(config *Config, posts []int64) {
	rl := ratelimit.New(9) // nginx limit

	for i, post := range posts {
		rl.Take()
		resp := GetPost(config, post)
		log.Print(resp.Artist)

		if i == 5 {
			log.Fatal("finish")
		}
	}
}
