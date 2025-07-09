# ContentScrap

ContentScrap is a Go-based API server that scrapes upcoming competitive programming contests from platforms like LeetCode and Codeforces. It exposes a simple HTTP endpoint that returns contest data in a unified JSON format.

## Getting Started

### Clone and Run

```bash
git clone https://github.com/0xamitr/contentScrap.git
cd contentScrap
go run main.go
```

The server will start at `http://localhost:8080`.

### Fetch Contests

```bash
curl http://localhost:8080/
```

## Example Response

```json
[
  {
    "title": "LeetCode Weekly Contest 412",
    "titleSlug": "weekly-contest-412",
    "start": Date object,
    "duration": 5400
  },
  {
    "title": "Codeforces Round #939",
    "titleSlug": "codeforces-round-939",
    "start": Date object,
    "duration": 7200
  }
]
```

## Project Structure

```
contentScrap/
├── main.go             // starts the HTTP server and handles /contests
├── types/              // to be done
│   └── contest.go      // struct definitions for contest data
├── scrap/
│   ├── leetcode/
│   │   └── leetcode.go     // fetches and parses contests from LeetCode
│   ├── codeforces/
│   │   └── codeforces.go   // fetches and parses contests from Codeforces
├── go.mod              // module definition
└── go.sum              // module checksums
```

## License

MIT

## Author

Amit Rathore ([0xamitr](https://github.com/0xamitr))