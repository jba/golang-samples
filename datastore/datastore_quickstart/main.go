// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START datastore_quickstart]
// Sample datastore_quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"

	// Imports the Google Cloud Datastore client package
	"cloud.google.com/go/datastore"
)

type Entity struct {
	Value string
}

func main() {
	ctx := context.Background()

	// Your Google Cloud Platform project ID
	projectID := "YOUR_PROJECT_ID"

	// Creates a client
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// The kind of the entity to retrieve
	kind := "Person"
	// The name/ID of the entity to retrieve
	name := "Bob"
	// The Datastore key for the entity
	key := datastore.NewKey(ctx, kind, name, 0, nil)

	entity := new(Entity)

	// Retrieves the entity
	if err := client.Get(ctx, key, entity); err != nil {
		log.Fatalf("Failed to get entity: %v", err)
	}

	fmt.Printf("Fetched entity: %v", key.String())
}

// [END datastore_quickstart]