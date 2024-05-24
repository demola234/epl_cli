package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

func GetData() {
	url := fmt.Sprintln("https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR")

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status:", resp.StatusCode)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	tableData, err := UnmarshalEplTable(body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tableData)

}
