package mock

import (
	"embed"
	"encoding/json"
	"io/ioutil"

	"github.com/SofaSubs/server/model"
)

//go:embed subs.json
var subsJSON embed.FS

func GetSubs() (map[string]model.Sub, error) {
	jsonFile, err := subsJSON.Open("subs.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var subs map[string]model.Sub
	if err := json.Unmarshal(byteValue, &subs); err != nil {
		return nil, err
	} else {
		return subs, nil
	}
}
