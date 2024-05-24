package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/demola234/models"
)

func GetData() (tableData models.EplTable, err error) {
	url := "<LiveApiKey>"

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
