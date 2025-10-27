// Small go project for pulling io from raider.io
package main

import (
//	"os"
	"github.com/petholla/go-wowio/pkg/character"
	"io"
	"fmt"
	"github.com/spf13/pflag"
	"net/http"
	"encoding/json"
	"time"
)

func pprint(data any) {
	fmt.Printf("%+v\n", data)
	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting data: %s\n", err)
	}
	fmt.Printf(string(pretty) + "\n")
}

func main() {
	name := pflag.String("name", "", "Character name")
	pflag.Parse()
	fmt.Printf("name: %s\n", *name)
	character, err := getChatacterInfo("fizzcrank", *name)
	if err != nil {
		fmt.Println(err)
	} else {
		pprint(character)
		fmt.Println(character.Score())
	}
}

func getChatacterInfo(realm string, name string) (wowio.Character, error) {
	var character wowio.Character
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

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		return character, fmt.Errorf("Error sending request: %s", err)
	} else if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		var character wowio.Character
		body, _ := io.ReadAll(response.Body)
		err := json.Unmarshal(body, &character)
		if err != nil {
			return character, err
		}
		return character, nil
	}
	return character, fmt.Errorf("Something went wrong.")
}
