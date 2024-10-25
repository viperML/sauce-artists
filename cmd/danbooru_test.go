package cmd

import (
	"log"
	"testing"
)

func TestGetPost(t *testing.T) {
    config := GetConfig()

    post := GetPost(&config, 6000000)
    log.Printf("%+v", post)
}
