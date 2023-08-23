# go-openfigi
Go client for the OpenFIGI API with support for caching
## Thanks
Many thanks to the OpenFIGI API-team for maintaining a perfect OpenAPI specification!
Also great thanks to @gregjones and all involved contributors for creating the awesome httpcache library!

## Usage
1. (Optional) Get an OpenFIGI API key for higher rates
2. (Optional) Select and create an http cache by using one of the factories in the `cache` package
3. Create a new `openfigi.APIClient` by calling the factory method `NewAPIClient`
4. Make requests against the api client, types for the client are in the `openfigi` package