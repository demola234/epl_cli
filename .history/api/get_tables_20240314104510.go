package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetData() {
	url := fmt.Sprintln("https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR")

	client := &http.Client{}

	resp, err = client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	

	fmt.Sprintln(resp.Status)
}
