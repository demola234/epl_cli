package api

import (
	"fmt"
	"net/http"
)

func getData() {
	url := fmt.Sprintf("https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR")

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Sprintf(re)
}
