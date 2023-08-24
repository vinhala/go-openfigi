package main_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/antihax/optional"
	main "github.com/ga42quy/go-openfigi"
	"github.com/ga42quy/go-openfigi/openfigi"
)

func TestNewClient(t *testing.T) {
	client := main.NewAPIClient(nil, "")
	if client == nil {
		t.Error("Expected client to be not nil")
	}
}

func TestPostMapping(t *testing.T) {
	client := main.NewAPIClient(nil, "")
	if client == nil {
		t.Error("Expected client to be not nil")
	}
	job := openfigi.MappingJob[string]{
		IdType:   "ID_WERTPAPIER",
		IdValue:  "851399",
		ExchCode: "US",
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	opts := openfigi.DefaultApiMappingPostOpts{
		Body: optional.NewInterface(job),
	}
	resp, _, err := client.DefaultApi.MappingPost(ctx, &opts)
	if err != nil {
		t.Error(err)
		return
	}
	if len(resp) == 0 {
		t.Errorf("Expected atleast one result, got none")
		return
	}
	if resp[0].Warning != "" {
		t.Errorf("Expected no warning, got %s", resp[0].Warning)
		return
	}
	fmt.Printf("%+v\n", resp[0])
}
