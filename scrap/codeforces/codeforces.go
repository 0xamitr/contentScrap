package codeforces

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Contest struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	Phase              string `json:"phase"`
	DurationSeconds    int64  `json:"durationSeconds"`
	StartTimeSeconds   int64  `json:"startTimeSeconds"`
	RelativeTimeSeconds int64 `json:"relativeTimeSeconds"`
}

type ContestListResponse struct {
	Status string    `json:"status"`
	Result []Contest `json:"result"`
}

func GetContests() ([]map[string]interface{}, error) {
	res, err := http.Get("https://codeforces.com/api/contest.list")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response ContestListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Load IST timezone
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		// Fallback in case tzdata isn't available
		ist = time.FixedZone("IST", 5*60*60+30*60)
	}

	var contests []map[string]interface{}

	for _, contest := range response.Result {
		if contest.Phase == "BEFORE" {
			start := time.Unix(contest.StartTimeSeconds, 0).In(ist)

			contests = append(contests, map[string]interface{}{
				"title":    contest.Name,
				"start":    start.Format("2006-01-02T15:04:05-07:00"), // ISO8601 with offset
				"duration": contest.DurationSeconds / 60,
				"type": "Codeforces",

			})
		}
	}

	return contests, nil
}
