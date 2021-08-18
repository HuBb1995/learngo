package model

import "encoding/json"

type Profile struct {
	Name      string
	Gender    string
	City      string
	Age       string
	Education string
	Marriage  string
	Height    string
	Income    string
}

type Chapter struct {
	Name    string
	Content string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var chapter Profile
	bytes, err := json.Marshal(o)
	if err != nil {
		return chapter, err
	}

	err = json.Unmarshal(bytes, &chapter)
	return chapter, err
}
