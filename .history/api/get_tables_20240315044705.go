package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/demola234/models"
)

func GetData() (tableData models.EplTable, err error) {
	url := "https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR"

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return tableData, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return tableData, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return tableData, err
	}

	tableData, err = models.UnmarshalEplTable(body)

	if err != nil {
		fmt.Println(err)
	}

	return tableData

}
