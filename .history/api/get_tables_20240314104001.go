package api

import (
	"encoding/json"
	"fmt"
	"http/net"
)

func getData() {
	url := fmt.Sprintf("https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR")

	resp, err := client.Get(url)

	fmt.Sprintf()
}