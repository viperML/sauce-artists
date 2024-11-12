package get

import (
	"log"
	"os"
)

type Config struct {
	username string
	api_key  string
	favgroup string
}

func GetConfig() Config {
	username, found := os.LookupEnv("DANBOORU_USERNAME")
	if !found {
		log.Fatal("DANBOORU_USERNAME not set")
	}

	api_key, found := os.LookupEnv("DANBOORU_APIKEY")
	if !found {
		log.Fatal("DANBOORU_APIKEY not set")
	}

	favgroup, found := os.LookupEnv("DANBOORU_FAVGROUP")
	if !found {
		log.Fatal("DANBOORU_FAVGROUP not set")
	}

	res := Config{
		username: username,
		api_key:  api_key,
		favgroup: favgroup,
	}

	return res
}
