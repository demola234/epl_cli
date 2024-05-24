package api

import (
	"fmt"
	"io"
	"net/http"
)

func GetData() {
	url := fmt.Sprintln("https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR")

	client := &http.Client{}

	resp, err = client.Do(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GoTube: Bad status: %s (%s)", resp.Status, http.StatusText(resp.StatusCode))
	}

	var body io.Reader

	fmt.Sprintln(resp.Status)
}
