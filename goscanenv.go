package goscanenv

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EnvOptions struct {
	EnvFile    string
	ExampleEnv string
	IgnoreEnv  string
}

func ScanEnv(opts ...EnvOptions) (bool, error) {
	options := EnvOptions{
		EnvFile:    ".env",
		ExampleEnv: ".env.example",
		IgnoreEnv:  ".envignore",
	}

	if len(opts) > 0 {
		options = opts[0]
	}

	err := godotenv.Load(options.EnvFile)

	if err != nil {
		log.Println("error: check the .env file.", err)
		return false, err
	}

	_, err = os.Stat(options.ExampleEnv)
	if err != nil {
		log.Printf("Example Env does not exist. %s",
			options.ExampleEnv)
		return false, err
	}

	exEnvs, err := godotenv.Read(options.ExampleEnv)

	if err != nil {
		log.Printf("Incorrect Syntax of %s. Please check keys are only using :, =.",
			options.ExampleEnv)
		return false, err
	}

	if len(exEnvs) == 0 {
		log.Printf("There is no key in %s. Please check if syntax is correct.",
			options.ExampleEnv)
		return true, nil
	}

	igEnvs := make(map[string]string)
	var ignoredEnvs []string

	_, err = os.Stat(options.IgnoreEnv)

	if err == nil {
		igEnvs, err = godotenv.Read(options.IgnoreEnv)
		if err != nil {
			log.Printf("Incorrect Syntax of %s. Please check keys are only using :, =.",
				options.IgnoreEnv)
		} else {
			ignoredEnvs = make([]string, len(igEnvs))
			i := 0
			for key := range igEnvs {
				ignoredEnvs[i] = key
				i++
			}
		}
	}

	missingEnvs := []string{}
	for key := range exEnvs {
		if os.Getenv(key) == "" && igEnvs[key] == "" {
			missingEnvs = append(missingEnvs, key)
		}
	}

	if len(missingEnvs) > 0 {
		fmt.Printf(
			`Env Scan Report
------------------------------------------------------------------------
Total Missing Environment Variables:
%d

Total Ignored Environment Variables:
%d

Missing Environment Variables: 
%s

Ignored Environment Variables: 
%s
------------------------------------------------------------------------
`, len(missingEnvs),
			len(igEnvs),
			strings.Join(missingEnvs, "\n"),
			strings.Join(ignoredEnvs, "\n"),
		)

		return false, fmt.Errorf("error: environment variables are missing. \n%v",
			missingEnvs)
	}

	return true, nil
}
