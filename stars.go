package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	dotenv "github.com/joho/godotenv"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: Replace [starred](https://github.com/maguowei/starred) with custom
func main() {
	err := dotenv.Load()
	handle(err)

	client := http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://api.github.com/users/hyperupcall/starred", nil)
	handle(err)

	env, hasToken := os.LookupEnv("GITHUB_TOKEN")
	if hasToken {
		req.Header.Add("Authorization", "token "+env)
	}

	res, err := client.Do(req)
	handle(err)
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	handle(err)

	type Repo struct {
		name             string
		full_name        string
		description      string
		fork             bool
		url              string
		size             int
		stargazers_count int
		watchers_count   int
		language         string
		forks_count      string
		license          struct {
			key     string
			name    string
			spdx_id string
			url     string
		}
		topics []string
	}

	var repos []Repo
	err = json.Unmarshal(content, &repos)
	handle(err)
	fmt.Printf("%+v", repos)
}
