package codeforces

import (
    "fmt"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
)

type Contest struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Phase             string `json:"phase"`
	DurationSeconds   int64  `json:"durationSeconds"`
	StartTimeSeconds  int64  `json:"startTimeSeconds"`
	RelativeTimeSeconds int64 `json:"relativeTimeSeconds"`
}

type ContestListResponse struct {
	Status string    `json:"status"`
	Result []Contest `json:"result"`
}

func GetContests() ([]map[string]interface{}, error) {
	res, err := http.Get("https://codeforces.com/api/contest.list")
	
	if(err != nil){
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
 
	var Response ContestListResponse
	err = json.Unmarshal(body, &Response)

	var contests []map[string]interface{}

	for _, contest := range Response.Result{
		if(contest.Phase == "BEFORE"){
			contests = append(contests, map[string]interface{}{
				"title":    contest.Name,
				"start":    time.Unix(contest.StartTimeSeconds, 0).Format("2006-01-02 15:04"),
				"duration": contest.DurationSeconds / 60,
			})
		}
	}
		return contests, nil
}
