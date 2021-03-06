package main

import (
	"io/ioutil"

	jenkins "github.com/bndr/gojenkins"
)

// Config is the set of parameters needed to configure the JenkinsCI provider.
type Config struct {
	ServerURL string
	CACert    string
	Username  string
	Password  string
}

func (c *Config) getAPIClient() (*jenkins.Jenkins, error) {

	client := jenkins.CreateJenkins(c.ServerURL, c.Username, c.Password)
	if c.CACert != "" {
		// provide CA certificate if server is using self-signed certificate
		client.Requester.CACert, _ = ioutil.ReadFile(c.CACert)
	}
	_, err := client.Init()
	if err != nil {
		return nil, err
	}

	// return the Jenkins API client
	return client, nil
}
