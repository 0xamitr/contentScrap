package leetcode

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
	// "fmt"
)

type Contest struct {
	Title           string `json:"title"`
	TitleSlug       string `json:"titleSlug"`
	Start           int64  `json:"startTime"`
	Duration        int64  `json:"duration"`
	OriginStartTime int64  `json:"originStartTime"`
	IsVirtual       bool   `json:"isVirtual"`
}

type GraphQLResponse struct {
	Data struct {
		All []Contest `json:"allContests"`
	} `json:"data"`
}

func GetContests() ([]map[string]interface{}, error) {
	query := `{"query":"query { allContests { title titleSlug startTime duration originStartTime isVirtual } }"}`

	req, err := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var gqlResp GraphQLResponse
	if err := json.Unmarshal(body, &gqlResp); err != nil {
		return nil, err
	}

	var contests []map[string]interface{}
	currentTime := time.Now().Unix()

	for _, contest := range gqlResp.Data.All {
		if contest.Start > currentTime && !contest.IsVirtual {
			contests = append(contests, map[string]interface{}{
				"title":    contest.Title,
				"start":    time.Unix(contest.Start, 0).Format("2006-01-02 15:04"),
				"duration": contest.Duration / 60,
			})
		}
	}
	return contests, nil
}
