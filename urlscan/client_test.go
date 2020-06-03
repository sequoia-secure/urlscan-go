package urlscan_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vertoforce/urlscan-go/urlscan"
)

type config struct {
	ApiKey string `json:"api_key"`
}

var cfg config

func init() {
	cfg.ApiKey = os.Getenv("URLSCAN_API_KEY")
	if cfg.ApiKey == "" {
		log.Fatal("no API KEY, environment variable URLSCAN_API_KEY is required.")
	}
}

func TestSubmitScan(t *testing.T) {
	client := urlscan.NewClient(cfg.ApiKey)
	task, err := client.Submit(context.Background(), urlscan.SubmitArguments{
		URL: "https://cookpad.com",
	})

	require.NoError(t, err)
	err = task.WaitForReport(context.Background())
	require.NoError(t, err)
}
