package get

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/ratelimit"
)

type Response struct {
	Artist string `json:"tag_string_artist"`
}

func Get(config *Config, path string) ([]byte, error) {
	url, err := url.Parse("https://danbooru.donmai.us" + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	query := url.Query()
	query.Add("login", config.username)
	query.Add("api_key", config.api_key)
	url.RawQuery = query.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("User-Agent", "curl/8.9.1")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	ct := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	if ct != "application/json" {
		return nil, fmt.Errorf("request failed")
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return bytes, nil
}

func GetPost(config *Config, postId int64) (*Response, error) {
	path := fmt.Sprintf("/posts/%d.json", postId)
	bytes, err := Get(config, path)

	if err != nil {
		return nil, fmt.Errorf("failed to query post %d: %w", postId, err)
	}

	var resp Response
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response for %d: %w", postId, err)
	}

	return &resp, nil
}

func CollectAuthors(config *Config, posts []int64) {
	rl := ratelimit.New(9) // nginx limit

	results := map[string]([]int64){}

	for _, post := range posts {
		rl.Take()
		resp, err := GetPost(config, post)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%d -> %s", post, resp.Artist)

		entry, ok := results[resp.Artist]
		if !ok {
			entry = []int64{}
		}
		entry = append(entry, post)
		results[resp.Artist] = entry
	}

	m, err := json.Marshal(&results)
	if err != nil {
		log.Fatal(err)
	}

    fmt.Print(string(m))
}
