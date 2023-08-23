package main

import (
	"github.com/gregjones/httpcache"

	"github.com/ga42quy/go-openfigi/openfigi"
)

// Create a new default API configuration
// This will control the used http client, default headers and caching
func newAPIConfiguration() *openfigi.Configuration {
	return &openfigi.Configuration{
		BasePath:      "https://api.openfigi.com/v3",
		DefaultHeader: make(map[string]string),
		UserAgent:     "go-openfigi",
	}
}

// Create a new default API client for the OpenFIGI API with optional caching
// You can either pass an API key here or pass it later through the context of each request
func NewAPIClient(cache *httpcache.Cache, apiKey string) *openfigi.APIClient {
	config := newAPIConfiguration()
	if cache != nil {
		config.HTTPClient = httpcache.NewTransport(*cache).Client()
	}
	if len(apiKey) > 0 {
		config.DefaultHeader["X-OPENFIGI-APIKEY"] = apiKey
	}
	return openfigi.NewAPIClient(config)
}
