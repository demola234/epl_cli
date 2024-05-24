package api

import (
	"fmt"
	"net/http"

	"github.com/google/martian/v3/body"
)

func GetData() {
	url := fmt.Sprintln("https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR")

	client := &http.Client{}
	

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := body.Read(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	tableData :=  EplTable{}

}
