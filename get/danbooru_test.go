package get

import (
	"log"
	"testing"
)

func TestGetPost(t *testing.T) {
    config := GetConfig()

    post, err := GetPost(&config, 6000000)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("%+v", post)
}
