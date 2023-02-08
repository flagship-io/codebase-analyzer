<p align="center">

<img  src="https://mk0abtastybwtpirqi5t.kinstacdn.com/wp-content/uploads/picture-solutions-persona-product-flagship.jpg"  width="211"  height="182"  alt="flagship"  />

</p>

<h3 align="center">Bring your features to life</h3>

[Website](https://flagship.io) | [Documentation](https://docs.developers.flagship.io/docs/codebase-analyzer) | [Installation Guide](https://docs.developers.flagship.io/docs/codebase-analyzer#run) | [Twitter](https://twitter.com/feature_flags)

[![Apache2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/flagship-io/flagship)](https://goreportcard.com/report/github.com/flagship-io/codebase-analyzer)
[![Go Reference](https://pkg.go.dev/badge/github.com/flagship-io/flagship.svg)](https://pkg.go.dev/github.com/flagship-io/codebase-analyzer)
[![codecov](https://codecov.io/gh/flagship-io/codebase-analyzer/branch/master/graph/badge.svg?token=71NXZN582Y)](https://codecov.io/gh/flagship-io/codebase-analyzer)

# Flagship Code Analyzer

Flagship Code Analyzer is a CLI and a docker image that can analyze your codebase and detect the usage of Flagship flags, in order to synchronize them with your Flag view in the platform.

## Run

### With CLI

```sh
go build

export FLAGSHIP_CLIENT_ID=FLAGSHIP_MANAGEMENT_API_CLIENT_ID
export FLAGSHIP_CLIENT_SECRET=FLAGSHIP_MANAGEMENT_API_CLIENT_SECRET
export ACCOUNT_ID=FLAGSHIP_ACCOUNT_ID
export ENVIRONMENT_ID=FLAGSHIP_ENVIRONMENT_ID
export REPOSITORY_URL=https://gitlab.com/org/repo
export REPOSITORY_BRANCH=master
export DIRECTORY=./
./code-analyzer
```

### With Docker

```sh
docker run -v $(pwd)/your_repo:/your_repo -e FLAGSHIP_CLIENT_ID=FLAGSHIP_MANAGEMENT_API_CLIENT_ID -e FLAGSHIP_CLIENT_SECRET=FLAGSHIP_MANAGEMENT_API_CLIENT_SECRET -e ACCOUNT_ID=FLAGSHIP_ACCOUNT_ID -e ENVIRONMENT_ID=your_env_id -e REPOSITORY_URL=https://gitlab.com/org/repo -e REPOSITORY_BRANCH=master -e DIRECTORY=/your_repo flagshipio/code-analyzer
```

### With Homebrew
```sh
export FLAGSHIP_CLIENT_ID=FLAGSHIP_MANAGEMENT_API_CLIENT_ID
export FLAGSHIP_CLIENT_SECRET=FLAGSHIP_MANAGEMENT_API_CLIENT_SECRET
export ACCOUNT_ID=FLAGSHIP_ACCOUNT_ID
export ENVIRONMENT_ID=FLAGSHIP_ENVIRONMENT_ID
export REPOSITORY_URL=https://gitlab.com/org/repo
export REPOSITORY_BRANCH=master
export DIRECTORY=./

brew tap flagship-io/flagship
brew install codebase-analyzer

codebase-analyzer
```

### Supported file languages
- .cs .fs
- .go
- .java
- .js .jsx
- .kt
- .m
- .php
- .py
- .swift
- .ts .tsx
- .vb

## Environment variables

### Flagship client ID (required)

This environment variable contains the Flagship client id necessary to authenticate request and send flags infos to the Flagship Platform

- example : `FLAGSHIP_CLIENT_ID=FLAGSHIP_MANAGEMENT_API_CLIENT_ID`

### Flagship client secret (required)

This environment variable contains the Flagship client secret necessary to authenticate request and send flags infos to the Flagship Platform

- example : `FLAGSHIP_CLIENT_SECRET=FLAGSHIP_MANAGEMENT_API_CLIENT_SECRET`

### Account ID (required)

This environment variable contains the Flagship account ID to synchronize flags usage for the matching environment

- example : `ACCOUNT_ID=your_flagship_env_id`

### Environment ID (required)

This environment variable contains the Flagship environment ID to synchronize flags usage for the matching environment

- example : `ENVIRONMENT_ID=your_flagship_env_id`

### Repository URL (required)

This environment variable should be set to the root URL of your repository, and is used to track the links of the files where your flags are used

- example : `REPOSITORY_URL=https://gitlab.com/org/repo`
- default value: ""

### Repository branch (required)

This environment variable should be set to the branch of the code you want to analyse, and is used to track the links of the files where your flags are used

- example : `REPOSITORY_BRANCH=master`
- default value: ""

### Files to exclude (optional)

This environment variable should contain a comma separated list of glob to exclude

- example : `FILES_TO_EXCLUDE=".git/*,go.mod,go.sum,main.go,internal/*,example/*"`
- default value: ""

### Directory to analyse (optional)

This environment variable should contain the directory you want to analyse

- example : `DIRECTORY="."`
- default value: "."

## Use with Gitlab CI

You can use the code analyzer to push flags references to your Flagship environment when code is pushed to a specific branch or tag

```yaml
analyze_flag_references:
  image: flagshipio/code-analyzer:master
  stage: analyze
  variables:
    REPOSITORY_URL: $CI_PROJECT_URL
    REPOSITORY_BRANCH: $CI_COMMIT_BRANCH
    FLAGSHIP_CLIENT_ID: FLAGSHIP_MANAGEMENT_API_CLIENT_ID
    FLAGSHIP_CLIENT_SECRET: FLAGSHIP_MANAGEMENT_API_CLIENT_SECRET
    ACCOUNT_ID: YOUR_ACCOUNT_ID
    ENVIRONMENT_ID: YOUR_ENVIRONMENT_ID
  script:
    - /root/code-analyser
  only:
    - master
```

### Dependencies

This repository needs go v1.13+ to work

### Install packages

`go mod download`

### Run

1. Go to the example directory, and copy paste a folder you want to analyse
2. Create a .env file to customize your environment variable
3. Run `go run *.go` in the example folder to run the code

### Test 

```
make test
```


## Contributors

- Guillaume Jacquart [@GuillaumeJacquart](https://github.com/guillaumejacquart)
- Kevin Jose [@kjose](https://github.com/kjose)
- Elias Cédric Laouiti [@eliaslaouiti](https://github.com/eliaslaouiti)
- Chadi Laoulaou [@Chadiii](https://github.com/chadiii)

## Licence

[Apache License.](https://github.com/flagship-io/codebase-analyzer/blob/master/LICENSE)

## About Flagship
​
<img src="https://www.flagship.io/wp-content/uploads/Flagship-horizontal-black-wake-AB.png" alt="drawing" width="150"/>
​
[Flagship by AB Tasty](https://www.flagship.io/) is a feature flagging platform for modern engineering and product teams. It eliminates the risks of future releases by separating code deployments from these releases :bulb: With Flagship, you have full control over the release process. You can:
​
- Switch features on or off through remote config.
- Automatically roll-out your features gradually to monitor performance and gather feedback from your most relevant users.
- Roll back any feature should any issues arise while testing in production.
- Segment users by granting access to a feature based on certain user attributes.
- Carry out A/B tests by easily assigning feature variations to groups of users.
​
<img src="https://www.flagship.io/wp-content/uploads/demo-setup.png" alt="drawing" width="600"/>
​
Flagship also allows you to choose whatever implementation method works for you from our many available SDKs or directly through a REST API. Additionally, our architecture is based on multi-cloud providers that offer high performance and highly-scalable managed services.
​
**To learn more:**
​
- [Solution overview](https://www.flagship.io/#showvideo) - A 5mn video demo :movie_camera:
- [Documentation](https://docs.developers.flagship.io/) - Our dev portal with guides, how tos, API and SDK references
- [Sign up for a free trial](https://www.flagship.io/sign-up/) - Create your free account
- [Guide to feature flagging](https://www.flagship.io/feature-flags/) - Everything you need to know about feature flag related use cases
- [Blog](https://www.flagship.io/blog/) - Additional resources about release management
