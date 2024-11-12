package main

import (
	"log"
	"os"
	"sauce-artists/db"

	"github.com/urfave/cli/v2"
)

func main() {
	var config AppConfig

	app := &cli.App{
		Name: "sauce-artists",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "danbooru_username",
				EnvVars:     []string{"DANBOORU_USERNAME"},
				Destination: &config.danbooru_apikey,
				Required:    true,
			},
		},
		Commands: []*cli.Command{
			{
				Name: "send",
				Action: func(cCtx *cli.Context) error {
					log.Printf("%+v", config)

					return nil
				},
			},
			{
				Name: "query",
				Action: func(cCtx *cli.Context) error {
					d, err := db.Init()
					if err != nil {
						log.Fatal(err)
					}

					db.WithDb(d, func(data *db.Data) error {
						data.Artists["foo"] = append(data.Artists["foo"], 1)

						return nil
					})

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
