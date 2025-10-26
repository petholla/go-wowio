// Small go project for pulling io from raider.io
package main

import (
//	"os"
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)

func pprint(data any) {
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting data: %s\n", err)
	}
	fmt.Printf(string(pretty) + "\n")
}

func main() {
	character, err := getChatacterInfo("fizzcrank", "iops")
	if err != nil {
		fmt.Println(err)
	} else {
		pprint(character)
	}
}

type Character struct {
	Name string `json:"name"`
	Class string `json:"class"`
	Spec string `json:"active_spec_name"`
	Role string `json:"active_spec_role"`
	Race string `json:"race"`
	Faction string `json:"faction"`
	BestRuns []Run `json:"mythic_plus_best_runs"`
}

type Run struct {
	Dungeon string `json:"dungeon"`
	MythicLevel int `json:"mythic_level"`
	Score float32 `json:"score"`
}

func getChatacterInfo(realm string, name string) (Character, error) {
	var character Character
	url := "https://raider.io/api/v1/characters/profile"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return character, fmt.Errorf("Error creating request: %s", err)
	}

	// Add query params
	query := request.URL.Query()
	query.Add("region", "us")
	query.Add("realm", realm)
	query.Add("name", name)
	query.Add("fields", "mythic_plus_scores_by_season:current,mythic_plus_best_runs")
	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		body, err2 := io.ReadAll(response.Body)
		if err2 != nil {
			return character, fmt.Errorf("Error reading body: %s\n", err)
		} else {
			return character, fmt.Errorf("Error getting character data: %s", body)
		}
	}
	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		var character Character
		body, _ := io.ReadAll(response.Body)
		err := json.Unmarshal(body, &character)
		if err != nil {
			return character, err
		}
		return character, nil
	}
	return character, fmt.Errorf("Something went wrong.")
}
