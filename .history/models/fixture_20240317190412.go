package models

import (
	"bytes"
	"encoding/json"
	"errors"
)

func UnmarshalFixturesEntity(data []byte) (FixturesEntity, error) {
	var r FixturesEntity
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FixturesEntity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FixturesEntity struct {
	Success bool        `json:"success"`
	Data    FixtureData `json:"data"`
}

type FixtureData struct {
	Fixtures []Fixture `json:"fixtures"`
	NextPage string    `json:"next_page"`
	PrevPage bool      `json:"prev_page"`
}

type Fixture struct {
	Odds             Odds                   `json:"odds"`
	HomeTranslations *HomeTranslationsUnion `json:"home_translations"`
	League           League                 `json:"league"`
	ID               int64                  `json:"id"`
	CompetitionID    int64                  `json:"competition_id"`
	HomeName         string                 `json:"home_name"`
	Competition      FixtureCompetition     `json:"competition"`
	Time             string                 `json:"time"`
	LeagueID         int64                  `json:"league_id"`
	Location         string                 `json:"location"`
	AwayID           int64                  `json:"away_id"`
	Date             string                 `json:"date"`
	GroupID          int64                  `json:"group_id"`
	Round            string                 `json:"round"`
	HomeID           int64                  `json:"home_id"`
	AwayName         string                 `json:"away_name"`
	AwayTranslations *AwayTranslationsUnion `json:"away_translations"`
	H2H              string                 `json:"h2h"`
}

type AwayTranslationsClass struct {
	Fa *string `json:"fa,omitempty"`
	Ar *string `json:"ar,omitempty"`
	Nl *string `json:"nl,omitempty"`
	Ko *string `json:"ko,omitempty"`
	Th *string `json:"th,omitempty"`
	Bg *string `json:"bg,omitempty"`
	El *string `json:"el,omitempty"`
	Ka *string `json:"ka,omitempty"`
	Ru *string `json:"ru,omitempty"`
}

type FixtureCompetition struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type HomeTranslationsClass struct {
	Fa *string `json:"fa,omitempty"`
	Ar *string `json:"ar,omitempty"`
	El *string `json:"el,omitempty"`
	Es *string `json:"es,omitempty"`
	Nl *string `json:"nl,omitempty"`
	Ko *string `json:"ko,omitempty"`
	Th *string `json:"th,omitempty"`
	Bg *string `json:"bg,omitempty"`
	Ka *string `json:"ka,omitempty"`
	Ru *string `json:"ru,omitempty"`
	Sv *string `json:"sv,omitempty"`
	Hu *string `json:"hu,omitempty"`
	Ro *string `json:"ro,omitempty"`
	Tr *string `json:"tr,omitempty"`
}

type League struct {
	Name      interface{} `json:"name"`
	CountryID interface{} `json:"country_id"`
	ID        interface{} `json:"id"`
}


type Live struct {
	The1 *float64 `json:"1"`
	The2 *float64 `json:"2"`
	X    *float64 `json:"X"`
}

type AwayTranslationsUnion struct {
	AnythingArray         []interface{}
	AwayTranslationsClass *AwayTranslationsClass
}

func (x *AwayTranslationsUnion) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.AwayTranslationsClass = nil
	var c AwayTranslationsClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.AwayTranslationsClass = &c
	}
	return nil
}

func (x *AwayTranslationsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingArray != nil, x.AnythingArray, x.AwayTranslationsClass != nil, x.AwayTranslationsClass, false, nil, false, nil, false)
}

type HomeTranslationsUnion struct {
	AnythingArray         []interface{}
	HomeTranslationsClass *HomeTranslationsClass
}

func (x *HomeTranslationsUnion) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.HomeTranslationsClass = nil
	var c HomeTranslationsClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.HomeTranslationsClass = &c
	}
	return nil
}

func (x *HomeTranslationsUnion) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingArray != nil, x.AnythingArray, x.HomeTranslationsClass != nil, x.HomeTranslationsClass, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
