package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SearchIssues() (*IssuesSearchResult, error) {
	fmt.Println("Start query formatting")
	resp, err := http.Get("https://www.omdbapi.com/?i=tt3896198&apikey=7e0bb40a&t=Titanic")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
