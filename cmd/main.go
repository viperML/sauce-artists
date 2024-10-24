package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type FavResp struct {
	id         int32
	name       string
	creator_id int32
	// post_ids   []int32
}

func Execute() {
	config := GetConfig()

	url, err := url.Parse("https://danbooru.donmai.us/favorite_groups/" + fmt.Sprint(config.favgroup) + ".json")
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
	body := string(bytes)
	log.Print(body)

	var favresp map[string](any)
    err = json.Unmarshal(bytes, &favresp)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%+v", favresp)
}
