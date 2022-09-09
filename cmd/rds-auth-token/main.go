package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
)

func usage() {
	commandName := os.Args[0]
	fmt.Fprintf(os.Stderr, `Generates the RDS IAM credentials using IAM

Usage:
	%s --user <user> --host <host> --port <port> [--region <region>]

`, commandName)
}

func main() {
	dbUser := flag.String("username", "", "database user")
	dbHost := flag.String("hostname", "", "database host. Must be the actual RDS endpoint.")
	dbPort := flag.String("port", "", "database port")
	region := flag.String("region", "", "database region. Optional.")
	flag.Parse()
	dbEndpoint := fmt.Sprintf("%s:%s", *dbHost, *dbPort)

	if *dbUser == "" || *dbHost == "" || *dbPort == "" {
		usage()
		log.Fatal("Error: missing args")
	}
	if !strings.HasSuffix(*dbHost, "rds.amazonaws.com") {
		usage()
		log.Fatal("Error: host must be of the format: <name>.<id>.<region>.rds.amazonaws.com")
	}

	if *region == "" {
		// Extract region from host
		r := strings.Split(*dbHost, ".")[2]
		region = &r
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal("failed authing into AWS: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.Background(), dbEndpoint, *region, *dbUser, cfg.Credentials,
	)
	if err != nil {
		log.Fatal("failed to create authentication token: " + err.Error())
	}

	fmt.Println(authenticationToken)
}
