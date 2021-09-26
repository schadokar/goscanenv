# goscanenv

`goscanenv` is a simple npm package which reads required environment variables
from `.env.example` and scan those in `.env`. It uses `dotenv` in background to load the `.env` file.

It returns a boolean.

## Install

```
go get -u github.com/schadokar/goscanenv
```

## Use

### With default files

- Environment File Name: `.env`
- Env Example File Name: `.env.example`
- Env Ignored File Name: `.envignore`

```go
package main

import (
	"fmt"
	"log"

	"github.com/schadokar/goscanenv"
)

func main() {
    scanResult, err := goscanenv.ScanEnv()

    if err != nil {
        log.Fatalln("Environment variables are missing.", err)   
        return
    }

    fmt.Println("All environment variables are set.",scanResult)
}
```

### With custom files

- Environment File Name: `.prod.env`
- Env Example File Name: `.prod.env.example`
- Env Ignored File Name: `.prod.envignore`

```go
package main

import (
    "fmt"

    "github.com/schadokar/godotenv"
)

func main() {
    opts := godotenv.EnvOptions{
		EnvFile:    "prod.env",
		ExampleEnv: "prod.env.example",
		IgnoreEnv:  "prod.envignore",
	}

    scanResult, err := godotenv.ScanEnv(opts)

    if err != nil {
        log.Fatalln("Environment variables are missing.", err)   
        return
    }

    fmt.Println("All environment variables are set.", scanResult)
}
```

## File Details

- `.env` file saves the original environment values.
- `.env.example` list all the env keys required by the application. Syntax of this env file must be like `.env`.
- `.envignore` list all the env keys which can be ignored if are missing.
