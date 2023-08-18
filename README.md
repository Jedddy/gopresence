Discord Presence
================

Create and use a custom discord presence.
Your activity name will be your `Application` name and it can only be changed in your [application's](https://discord.com/developers/applications) settings.


## Installation
```
go get github.com/jedddy/gopresence
```

## Getting the Client ID
- Create an Application [here](https://discord.com/developers/applications).
- Go to `OAuth` section and copy the `Client ID`

## Creating Assets
- Go to your Application settings
- Go to the `Rich Presence` section and click `Art Assets`
- Upload your image asset (min: 512x512) and copy the ID.

## Usage Example

```go
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
        // All Activity arguments are optional
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

			Buttons: []gopresence.Button{ // maximum of 2 buttons
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

```
See more about activity objects [here](https://discord.com/developers/docs/topics/gateway-events#activity-object).
