package scrap

import (
	// "fmt"
	"contestScrap/scrap/leetcode"
	// "contestScrap/scrap/gfg"
	"contestScrap/scrap/codeforces"
)

func ScrapLeetcode() ([]map[string]interface{}, error) {
	return leetcode.GetContests()
}

// func ScrapGfg(){
// 	return gfg.GetContests()
// }

func ScrapCodeforces() ([]map[string]interface{}, error) {
	return codeforces.GetContests()
}

func GetContests() ([]map[string]interface{}, error) {
	cf, err := ScrapCodeforces()
	if err != nil {
		return nil, err
	}

	lc, err := ScrapLeetcode()
	if err != nil {
		return nil, err
	}

	combined := append(cf, lc...)
	return combined, nil
}