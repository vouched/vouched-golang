package example

import (
	"fmt"

	APIClient "github.com/vouched/vouched-golang"
	client "github.com/vouched/vouched-golang"
)

func submitVerificationJob() {
	c := APIClient.New("Set your private Vouched key from .env")
	var idPhotoBase64 = client.ReadImage("./tests/oh-license.jpeg")
	var userPhotoBase64 = client.ReadImage("./tests/oh-selfie.jpeg")
	params := map[string]interface{}{
		"userPhoto":   userPhotoBase64,
		"idPhoto":     idPhotoBase64,
		"callbackURL": "https://www.google.com",
		"type":        "id-verification",
		"firstName":   "Janice",
		"dob":         "06/22/1990", "lastName": "Way"}
	if resp, err := c.Submit(params); err != nil {
		fmt.Printf("Error: %+v\n", err)
	} else {
		fmt.Printf("Job: %+v\n", resp)
	}
}

func getAllJobs() {
	c := client.New("Set your private Vouched key from .env")
	params := map[string]interface{}{
		"page":       1,
		"sortBy":     "date",
		"sortOrder":  "desc",
		"from":       "1990-12-24T04:44:00+00:00",
		"to":         "2020-12-24T04:44:00+00:00",
		"type":       "id-verification",
		"token":      "SESSION_TOKEN",
		"status":     "active",
		"withPhotos": false,
		"pageSize":   2,
	}
	if resp, err := c.Jobs(params); err != nil {
		fmt.Printf("Error: %+v\n", err)
	} else {
		fmt.Printf("Jobs: %+v\n", resp)
	}
}

func removeJob() {
	c := APIClient.New("Set your private Vouched key from .env")
	if resp, err := c.RemoveJob("ZicnypPn"); err != nil {
		fmt.Printf("Error: %+v\n", err)
	} else {
		fmt.Printf(" %+v\n", resp)
	}
}
