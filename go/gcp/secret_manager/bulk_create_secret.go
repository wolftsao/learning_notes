package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"gopkg.in/yaml.v2"
)

// Secret Type for single secret
type Secret struct {
	Name   string            `yaml: name`
	Value  string            `yaml: value`
	Labels map[string]string `yaml: labels`
}

// SecretFile Type to hold the whole yaml data
type SecretFile struct {
	Project string   `yaml: project`
	Secrets []Secret `yaml: inline`
}

func main() {
	// replace <path to secret yaml> with the yaml file which stores the secrets
	yamlFile, err := ioutil.ReadFile("./secrets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var t SecretFile

	err = yaml.Unmarshal(yamlFile, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	ctx := context.Background()
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	for _, v := range t.Secrets {
		ctx := context.Background()
		c, err = secretmanager.NewClient(ctx)
		if err != nil {
			log.Fatalf("failed to setup client: %v", err)
		}

		req := &secretmanagerpb.CreateSecretRequest{
			Parent:   fmt.Sprintf("projects/%s", t.Project),
			SecretId: v.Name,
			Secret: &secretmanagerpb.Secret{
				Replication: &secretmanagerpb.Replication{
					Replication: &secretmanagerpb.Replication_Automatic_{
						Automatic: &secretmanagerpb.Replication_Automatic{},
					},
				},
				Labels: v.Labels,
			},
		}
		resp, err := c.CreateSecret(ctx, req)
		if err != nil {
			log.Fatalf("failed to create secret: %v", err)
		}

		payload := []byte(v.Value)

		addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
			Parent: resp.Name,
			Payload: &secretmanagerpb.SecretPayload{
				Data: payload,
			},
		}

		// Call the API.
		version, err := c.AddSecretVersion(ctx, addSecretVersionReq)
		if err != nil {
			log.Fatalf("failed to add secret version: %v", err)
		}

		fmt.Println(version)
	}
}
