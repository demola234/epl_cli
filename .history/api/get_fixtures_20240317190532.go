package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/demola234/models"
)

func GetFixtureData() (fixtureData models.GetFixtureData, err error) {
	url := "https://livescore-api.com/api-client/fixtures/matches.json?&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR"

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
		return tableData, err
	}

	return tableData, nil

}
