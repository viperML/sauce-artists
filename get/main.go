package get

import (
	"encoding/json"
	"fmt"
	"log"
)

type FavResp struct {
	Name    string
	PostIds []int64 `json:"post_ids"`
}

func Execute() {
	config := GetConfig()

    bytes, err := Get(&config, fmt.Sprintf("/favorite_groups/%s.json", config.favgroup))

	var favresp FavResp
	err = json.Unmarshal(bytes, &favresp)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", favresp)

    CollectAuthors(&config, favresp.PostIds)
}
