# Flagship Code Analyzer

Flagship Code Analyzer is a CLI and a docker image that can analyze your codebase and detect the usage of Flagship flags, in order to synchronize them with your Flag view in the platform.

# Run

docker run -v $(pwd)/your_repo:/your_repo -e FLAGSHIP_TOKEN=your_token -e REPOSITORY_URL=https://gitlab.com/org/repo -e REPOSITORY_BRANCH=master -e DIRECTORY=/your_repo flagshipio/code-analyzer

# Environment variables

## Flagship token (required)

This environment variable contains the Flagship token necessary to send Flag infos to the Flagship Platform

- example : `FLAGSHIP_TOKEN=your_token`

## Environment ID (required)

This environment variable contains the Flagship environment ID to synchronize flag usage for the matching environment

- example : `ENVIRONMENT_ID=your_flagship_env_id`

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

## Directory to analyse (optional)

This environment variable should contain the directory you want to analyse

- example : `DIRECTORY="."`
- default value: "."

# Use with Gitlab CI

You can use the code analyzer to push flag references to your Flagship environment when code is pushed to a specific branch or tag

```yaml
analyze_flag_references:
  image: flagshipio/code-analyzer:master
  stage: analyze
  variables:
    REPOSITORY_URL: $CI_PROJECT_URL
    REPOSITORY_BRANCH: $CI_COMMIT_BRANCH
    FLAGSHIP_TOKEN: YOUR_FLAGSHIP_TOKEN
    ENVIRONMENT_ID: YOUR_ENVIRONMENT_ID
  script:
    - /root/code-analyser
  only:
    - master
```

# Contribute

## Dependencies

This repository needs go v1.13+ to work

## Install packages

`go mod download`

## Run

1. Go to the example directory, and copy paste a folder you want to analyse
2. Create a .env file to customize your environment variable
3. Run `go run .go` in the example folder to run the code
