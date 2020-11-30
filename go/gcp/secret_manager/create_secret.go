package main

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {
	// import secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	ctx := context.Background()
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	req := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", "<Project ID>"),
		SecretId: "<Secret Name>",
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
			Labels: map[string]string{
				"<label1>": "<label1 value>",
				"<label2>": "<label2 value>",
			},
		},
	}
	resp, err := c.CreateSecret(ctx, req)
	if err != nil {
		log.Fatalf("failed to create secret: %v", err)
	}

	payload := []byte("<Secret>")

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
