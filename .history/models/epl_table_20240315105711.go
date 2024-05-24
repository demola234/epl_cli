package models

import "encoding/json"

func UnmarshalEplTable(data []byte) (EplTable, error) {
	var r EplTable
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EplTable) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EplTable struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Table []Table `json:"table"`
}

type Table struct {
	LeagueID      string    `json:"league_id"`
	SeasonID      string    `json:"season_id"`
	Name          string    `json:"name"`
	Rank          string    `json:"rank"`
	Points        string    `json:"points"`
	Matches       string    `json:"matches"`
	GoalDiff      string    `json:"goal_diff"`
	GoalsScored   string    `json:"goals_scored"`
	GoalsConceded string    `json:"goals_conceded"`
	Lost          string    `json:"lost"`
	Drawn         string    `json:"drawn"`
	Won           string    `json:"won"`
	TeamID        string    `json:"team_id"`
	CompetitionID string    `json:"competition_id"`
	GroupID       string    `json:"group_id"`
	GroupName     string    `json:"group_name"`
	StageName     StageName `json:"stage_name"`
	StageID       string    `json:"stage_id"`
}

type StageName string

const (
	RegularSeason StageName = "Regular Season"
)
