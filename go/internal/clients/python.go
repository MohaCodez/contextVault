package clients

import (
	"errors"
	"net/http"
	"time"
)

type PythonClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewPythonClient(baseURL string) *PythonClient {
	return &PythonClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (c *PythonClient) Health() error {
	resp, err := c.HTTPClient.Get(c.BaseURL + "/health")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("python service unhealthy")
	}

	return nil
}
