package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab/canarybay/aws/integration/code-analyser/internal/files/model"
	"log"
	"net/http"
	"os"
)

// FlagInfo represent a flag code info sent to the Flagship API
type FlagInfo struct {
	FlagKey          string `json:"flagKey"`
	RepositoryURL    string `json:"repositoryUrl"`
	RepositoryBranch string `json:"repositoryBranch"`
	FilePath         string `json:"filePath"`
	LineNumber       int    `json:"lineNumber"`
	Code             string `json:"code"`
}

// SendFlagsToAPI takes file search result & sends flag info to the API
func SendFlagsToAPI(results []model.FileSearchResult, envId string) (err error) {
	flagInfos := []FlagInfo{}
	for _, fr := range results {
		for _, r := range fr.Results {
			flagInfos = append(flagInfos, FlagInfo{
				FlagKey:          r.FlagKey,
				RepositoryURL:    os.Getenv("REPOSITORY_URL"),
				RepositoryBranch: os.Getenv("REPOSITORY_BRANCH"),
				FilePath:         fr.File,
				LineNumber:       r.LineNumber,
				Code:             r.CodeLines,
			})
		}
	}

	callApi(envId, flagInfos)

	return err
}

func callApi(envId string, flagInfos []FlagInfo) {
	syncFlagsRoute := os.Getenv("FS_API") + "/account_environments/" + envId + "/flag_usages"

	json, _ := json.Marshal(flagInfos)

	fmt.Println("Calling PUT", syncFlagsRoute)

	req, err := http.NewRequest("PUT", syncFlagsRoute, bytes.NewBuffer(json))
	if err != nil {
		log.Fatal("Error in request", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpYXQiOjE2MTU0NjAwMjksImV4cCI6MTYxNTUwMzIyOSwic3ViIjoiYmtvNmQ4N3N1MWowMDI5OWljYzAiLCJyb2xlcyI6WyJST0xFX0FETUlOIiwiUk9MRV9BQlRBU1RZIiwiUk9MRV9VU0VSIl0sInVzZXJuYW1lIjoia2V2aW4uam9zZUBhYnRhc3R5LmNvbSJ9.oIhbuGZVG3xww9Tf7b-K6Ceeu1BwXZ7PY5tilNptVfGn7ep0DX-D_Fx9SCDBBzhA-9ZJSW8CyZ1i9ZgmGF8lt7Dw8v9kEl0HZIFG2_Wo-eoX_2PHTvoqx3DMFvngvqjHG_yzuA29cP_IcfvllqleXAzBH_RyhLNY0ZPudoHnGxZQVCFiUm0x8NSGa7WpImWKnfxk-9-PSH9Rss4I78xd2-SUF2Cf5IoSRd0InJwAZ_ggtRdaXVxajt3sSZ0gI2PBFCydC8uH_8N8QUUhcPapM02zOOAmJG2Ik_7jZHrnsUQsrvGZmZjqdHXuWG122UIQbaxWNWsV8o7njco3qeZ24HKRtd7q-vV1B1whMiwN7TuqBzC2fIV6uN8emIGeEwx7xD1xCM9PZvAWtPKrgvvKUfOJX9odwvLWNyYzHONhubVu001ZwwFUDBgjmaGkKxVxSSpuAHWpubB6xo-S0P4iE5uw1xkZZW0AwM4yzWjLeRyPALjRFcyQoq7WlOc03v-5IFET9IiA987ZJb5QHaQi-SRhqJQ1Yur_kcesfrQxIKqGdXq9oNpvo9cKwsTrFA8wRUkC3f8aX6Ot6jPgt6AoBwED5aedoW0n6S-KHehSxQsSFHZGNAofXJ4Z7di9Ok5sSpM3ktHJscrjbE7Z5JDSVF4AZqMxIcIuQiv6F8lSzpA")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error in request execution", err.Error())
	}

	fmt.Println("Response :", resp)
	defer resp.Body.Close()
}
