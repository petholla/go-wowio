// x
package main

import (
	"os"
	"io"
	"fmt"
	"net/http"
)

func main() {
	data, _ := getChatacterInfo("fizzcrank", "iops")
	fmt.Println(data)
}

func getChatacterInfo(realm string, name string) (string, error) {
	url := "https://raider.io/api/v1/characters/profile"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
		fmt.Printf("error sending request: %s\n", err)
		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("error reading body: %s\n", err)
		}
		fmt.Printf("body: %s\n", body)
		os.Exit(1)
	}
	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("error reading body: %s\n", err)
		}
		return string(body), nil
	}
	return "", fmt.Errorf("Something went wrong.")
}
