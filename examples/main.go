package main

import (
	"log"
	"time"

	"github.com/jedddy/gopresence"
)

func main() {
	gp, err := gopresence.New("CLIENT_ID")

	if err != nil {
		log.Fatal(err)
	}

	now := time.Now().Unix()

	for {
		err := gp.SetActivity(gopresence.Activity{
			State:   "Activity State",
			Details: "Activity Details",
			Timestamps: &gopresence.Timestamps{
				Start: now,
			},
			Assets: &gopresence.Assets{
				LargeImage: "20210904_113603", // Id of the image asset
			},
			Party: &gopresence.Party{
				ID:   "Party ID",
				Size: [2]int{1, 10}, // size must be a two item array, [currentSize, maxSize]
			},

			Buttons: &[2]gopresence.Button{
				{
					Label: "Button 1",
					Url:   "https://pkg.go.dev/github.com/jedddy/gopresence",
				},
				{
					Label: "Button 2",
					Url:   "https://pkg.go.dev/github.com/jedddy/gopresence",
				},
			},
		})

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second)
	}
}
