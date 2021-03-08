# Run

docker run -v ./your_repo ./your_repo -e FLAGSHIP_TOKEN=your_token -e REPOSITORY_URL=https://gitlab.com/org/repo -e REPOSITORY_BRANCH=master -e DIRECTORY=./your_repo flagshipio/code-analyzer

# Environment variables

## Flagship token (required)
This environment variable contains the Flagship token necessary to send Flag infos to the Flagship Platform

## Repository URL (required)
This environment variable should be set to the root URL of your repository, and is used to track the links of the files where your flags are used

- example : `REPOSITORY_URL=https://gitlab.com/org/repo`
- default value: ""

## Repository branch (required)
This environment variable should be set to the branch of the code you want to analyse, and is used to track the links of the files where your flags are used
- example : `REPOSITORY_BRANCH=master`
- default value: ""

## Files to exclude (optional)
This environment variable should contain a comma separated list of glob to exclude
- example : `FILES_TO_EXCLUDE=".git/*,go.mod,go.sum,main.go,internal/*,example/*"`
- default value: ""

## Directory to analyse
This environment variable should contain the directory you want to analyse
- example : `DIRECTORY="."`
- default value: "."

# Contribute

## Dependencies
This repository needs go v1.13+ to work

## Install packages
`go mod download`

## Run
1. Go to the example directory, and copy paste a folder you want to analyse
2. Create a .env file to customize your environment variable
3. Run `go run .go` in the example folder to run the code