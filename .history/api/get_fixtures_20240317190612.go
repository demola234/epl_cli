package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/demola234/models"
)

func GetFixtureData() (fixtureData models.FixturesEntity, err error) {
	url := "https://livescore-api.com/api-client/fixtures/matches.json?&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR"

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return fixtureData, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fixtureData, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fixtureData, err
	}

	fixtureData, err = models.UnmarshalEplTable(body)

	if err != nil {
		return fixtureData, err
	}

	return fixtureData, nil

}
