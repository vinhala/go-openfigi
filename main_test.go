package main_test

import (
	"context"
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
	jobs := []openfigi.MappingJob[string]{job}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	opts := openfigi.DefaultApiMappingPostOpts{
		Body: optional.NewInterface(jobs),
	}
	res, _, err := client.DefaultApi.MappingPost(ctx, &opts)
	if err != nil {
		t.Error(err)
		return
	}
	if len(res) == 0 {
		t.Errorf("Expected atleast one result, got none")
		return
	}
	if res[0].Warning != "" {
		t.Errorf("Expected no warning, got %s", res[0].Warning)
		return
	}
}

func TestAPIKeyHeader(t *testing.T) {
	client := main.NewAPIClient(nil, "MyAPIKey")
	if client == nil {
		t.Error("Expected client to be not nil")
	}
	job := openfigi.MappingJob[string]{
		IdType:   "ID_WERTPAPIER",
		IdValue:  "851399",
		ExchCode: "US",
	}
	jobs := []openfigi.MappingJob[string]{job}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	opts := openfigi.DefaultApiMappingPostOpts{
		Body: optional.NewInterface(jobs),
	}
	_, resp, _ := client.DefaultApi.MappingPost(ctx, &opts)
	//Should have a header for the API key
	head := resp.Request.Header.Get("X-OPENFIGI-APIKEY")
	if head != "MyAPIKey" {
		t.Errorf("Expected header to be MyAPIKey, got %s", head)
		return
	}
}
