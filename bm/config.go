package bm

// Config contains information that helps during the setup of a new Mollie client.
type Config struct {
	auth string
}

// NewConfig builds a Mollie configuration object,
// it takes t to indicate if our client is meant to create requests for testing
// and auth to indicate the authentication method we want to use.
func NewConfig(auth string) *Config {
	return &Config{
		auth: auth,
	}
}
