package main

import "github.com/simpleforce/simpleforce"

var (
	sfURL      = "Custom or instance URL, for example, 'https://na01.salesforce.com/'"
	sfUser     = "Username of the Salesforce account."
	sfPassword = "Password of the Salesforce account."
	sfToken    = "Security token, could be omitted if Trusted IP is configured."
)

func createClient() *simpleforce.Client {
	client := simpleforce.NewClient(sfURL, simpleforce.DefaultClientID, simpleforce.DefaultAPIVersion)
	if client == nil {
		// handle the error

		return nil
	}

	err := client.LoginPassword(sfUser, sfPassword, sfToken)
	if err != nil {
		// handle the error

		return nil
	}

	// Do some other stuff with the client instance if needed.

	return client
}

func main() {

}
