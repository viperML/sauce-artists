package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adrg/xdg"
)

type Data struct {
	lastModified int64              `json:"last_modified"`
	Artists      map[string][]int64 `json:"artists"`
}

type Db struct {
	path string
}

func Init() (*Db, error) {
	path := fmt.Sprintf("%s/sauce-artists", xdg.CacheHome)
	log.Printf("database path: %s", path)

	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			seed := new(Data)
			seed.lastModified = time.Now().Unix()
			seed.Artists = map[string][]int64{}

			bs, err := json.Marshal(&seed)
			if err != nil {
				return nil, fmt.Errorf("Failed to init database: %w", err)
			}
			os.WriteFile(path, bs, 0644)
		} else {
			return nil, fmt.Errorf("Unknown error when opening the database: %w", err)
		}
	}

	res := Db{path}
	return &res, nil
}

func WithDb(db *Db, transaction func(*Data) error) error {
	bs, err := os.ReadFile(db.path)
	if err != nil {
		return fmt.Errorf("Failed to open database: %w", err)
	}

	data := new(Data)

	err = json.Unmarshal(bs, data)
	if err != nil {
		return fmt.Errorf("Failed to decode database: %w", err)
	}

	err = transaction(data)
	if err != nil {
		return fmt.Errorf("Failed to perform transaction: %w", err)
	}

	data.lastModified = time.Now().Unix()

	newbs, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Failed to encode database: %w", err)
	}

	os.WriteFile(db.path, newbs, 0644)

	return nil
}
