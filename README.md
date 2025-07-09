ContentScrap is a Go-based API server that scrapes upcoming competitive programming contests from platforms like LeetCode and Codeforces. It exposes a single endpoint that returns all contests in a unified JSON format. To run the project, clone the repository using:

git clone https://github.com/0xamitr/contentScrap.git
cd contentScrap
go run main.go

This starts a server on http://localhost:8082. Hitting the /contests endpoint returns a list of upcoming contests. Example:

curl http://localhost:8082/

Response:

[
  {
    "title": "LeetCode Weekly Contest 412",
    "titleSlug": "weekly-contest-412",
    "start": Date Object,
    "duration": 5400
  },
  {
    "title": "Codeforces Round #939",
    "titleSlug": "codeforces-round-939",
    "start": Date Object,
    "duration": 7200
  }
]

The response fields are:
- title: name of the contest
- titleSlug: URL-friendly version
- start: UNIX timestamp (seconds)
- duration: duration in seconds

The project structure is:

contentScrap/
├── main.go                // starts the HTTP server and handles /contests
├── types/                 // to be done
│   └── contest.go         // struct definitions for contest data
├── scrap/
│   ├── leetcode
│   │   └── leetcode.go    // fetches and parses contests from LeetCode
│   ├── codeforces
│   │   └── codeforces.go  // fetches and parses contests from Codeforces
├── go.mod                 // module definition
└── go.sum                 // module checksums

The server uses only the Go standard library for HTTP, JSON parsing, and fetching content. More platforms can be added by implementing additional scrapers inside the utils/ folder and updating main.go to include them. The project is licensed under MIT and maintained by Amit Rathore (https://github.com/0xamitr).
