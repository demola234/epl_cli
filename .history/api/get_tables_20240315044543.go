package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/demola234/models"
)

func GetData() tableData EplTable {
	url := "https://livescore-api.com/api-client/leagues/table.json?competition_id=2&key=ZSZxPcrq7SOFjrhA&secret=WikIyrJXCtYBVvxVoxsdJUERNtx6UOHR"

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

	tableData, err := models.UnmarshalEplTable(body)

	if err != nil {
		fmt.Println(err)
	}

	// table := fmt.Sprintf("EPL Table Data: %v", tableData)

	fmt.Printf("%s %s %s %s %s %s %s\n", "Rank", "Name", "Points", "GoalDiff", "GroupName", "Drawn", "Won")

	// Format details
	for _, row := range tableData.Data.Table {
		// Print row details

		fmt.Printf("%s %s %s %s %s %s %s\n", row.Rank, row.Name,
			row.Points, row.GoalDiff,
			row.GroupName, row.Drawn, row.Won)
	}

}
