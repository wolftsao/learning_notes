package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"wolftsao/gce_sm/helper"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"gopkg.in/yaml.v2"
)

func main() {
	args := os.Args

	secretYaml, err := ioutil.ReadFile(args[2])
	if err != nil {
		log.Fatalf("Failed to get secret yaml file: #%v\n", err)
	}

	var secretFile helper.SecretFile
	err = yaml.Unmarshal(secretYaml, &secretFile)
	if err != nil {
		log.Fatalf("Failed to parse secret yaml file: %v\n", err)
	}

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to setup client: %v\n", err)
	}

	var wg sync.WaitGroup

	switch args[1] {
	case "bulk-create":
		helper.BulkCreateSecrets(ctx, client, &wg, secretFile)
	case "bulk-delete":
		helper.BulkDeleteSecrets(ctx, client, &wg, secretFile)
	default:
		log.Fatalf("Unsupported operation: %v\n", args[1])
	}

	wg.Wait()
}
