This Repo has been moved to https://github.com/vulncheck-oss/sdk-go

# The VulnCheck SDK For Go
Bring the VulnCheck API to your Go applications.

[![Release](https://img.shields.io/github/v/release/vulncheck-oss/sdk-go)](https://github.com/vulncheck-oss/sdk-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/vulncheck-oss/sdk-go)](https://goreportcard.com/report/github.com/vulncheck-oss/sdk-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/vulncheck-oss/sdk-go.svg)](https://pkg.go.dev/github.com/vulncheck-oss/sdk-go)
[![Lint](https://github.com/vulncheck-oss/sdk-go/actions/workflows/lint.yml/badge.svg)](https://github.com/vulncheck-oss/sdk-go/actions/workflows/lint.yml)
[![Tests](https://github.com/vulncheck-oss/sdk-go/actions/workflows/test.yml/badge.svg)](https://github.com/vulncheck-oss/sdk-go/actions/workflows/test.yml)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/vulncheck-oss/sdk-go/pulls)

## Installation

```bash
go get github.com/vulncheck-oss/sdk-go
```


## Examples

### Connecting to the API

```go
package main

import (
	"fmt"
	"github.com/vulncheck-oss/sdk-go"
)

func main() {
    client := sdk.Connect("https://api.vulncheck.com", "vulncheck_token")
    fmt.Println(client.GetUrl())
}
```

### Available Methods

### PURL
```go
response, err := client.GetPurl("pkg:hex/coherence@0.1.2")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### CPE
```go
response, err := client.GetCpe("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### BACKUP

```go
require github.com/vulncheck-oss/sdk
```
to:
```go
require github.com/vulncheck-oss/sdk-go
```
