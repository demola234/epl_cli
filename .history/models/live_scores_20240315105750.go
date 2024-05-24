
packagemodels

import "encoding/json"

func UnmarshalLiveScore(data []byte) (LiveScore, error) {
	var r LiveScore
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LiveScore) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LiveScore struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	Match []Match `json:"match"`
}

type Match struct {
	Scheduled   string      `json:"scheduled"`
	Competition Competition `json:"competition"`
	Time        string      `json:"time"`
	LastChanged string      `json:"last_changed"`
	FixtureID   int64       `json:"fixture_id"`
	Status      string      `json:"status"`
	Home        Away        `json:"home"`
	Added       string      `json:"added"`
	ID          int64       `json:"id"`
	Country     *Country    `json:"country"`
	Location    string      `json:"location"`
	Odds        Odds        `json:"odds"`
	Away        Away        `json:"away"`
	Federation  *Federation `json:"federation"`
	Outcomes    Outcomes    `json:"outcomes"`
	Scores      Scores      `json:"scores"`
	Urls        Urls        `json:"urls"`
}

type Away struct {
	Name      string `json:"name"`
	CountryID int64  `json:"country_id"`
	ID        int64  `json:"id"`
	Stadium   string `json:"stadium"`
}

type Competition struct {
	HasGroups         bool   `json:"has_groups"`
	Name              string `json:"name"`
	Tier              int64  `json:"tier"`
	NationalTeamsOnly bool   `json:"national_teams_only"`
	IsCup             bool   `json:"is_cup"`
	ID                int64  `json:"id"`
	IsLeague          bool   `json:"is_league"`
	Active            bool   `json:"active"`
}

type Country struct {
	Name     string `json:"name"`
	FifaCode string `json:"fifa_code"`
	ID       int64  `json:"id"`
	IsReal   bool   `json:"is_real"`
	UefaCode string `json:"uefa_code"`
}

type Federation struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type Odds struct {
	Live Live `json:"live"`
	Pre  Live `json:"pre"`
}

type Live struct {
	The1 *float64 `json:"1"`
	The2 *float64 `json:"2"`
	X    *float64 `json:"X"`
}

type Outcomes struct {
	HalfTime        *string     `json:"half_time"`
	FullTime        interface{} `json:"full_time"`
	ExtraTime       interface{} `json:"extra_time"`
	PenaltyShootout interface{} `json:"penalty_shootout"`
}

type Scores struct {
	Score   string `json:"score"`
	HTScore string `json:"ht_score"`
	FtScore string `json:"ft_score"`
	EtScore string `json:"et_score"`
	PSScore string `json:"ps_score"`
}

type Urls struct {
	Events     string `json:"events"`
	Statistics string `json:"statistics"`
	Lineups    string `json:"lineups"`
	Head2Head  string `json:"head2head"`
}
