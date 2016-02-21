package main

import "fmt"

// BEGIN OMIT

type Config struct {
	Env, AccessKey, SecretKey string
}

func (c *Config) String() string {
	type c2 Config
	cs := c2(*c)
	cs.AccessKey = "(REDACTED)"
	cs.SecretKey = "(REDACTED)"
	return fmt.Sprintf("%+v", cs)
}

func main() {
	c := &Config{
		Env:       "Not secret environment",
		AccessKey: "Secret Access Key",
		SecretKey: "Very secret Key",
	}

	fmt.Printf("AccessKey: %s\n", c.AccessKey)
	fmt.Printf("Config: %+v", c)
}

// END OMIT
