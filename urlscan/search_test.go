package urlscan_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vertoforce/urlscan-go/urlscan"
)

func TestSearch(t *testing.T) {
	resp, err := urlscan.Search(context.Background(), urlscan.SearchArguments{
		Query: "ip:163.43.24.70",
	})

	require.NoError(t, err)
	assert.NotEqual(t, 0, len(resp.Results))
	assert.NotEqual(t, "", resp.Results[0].ID)
}

func TestSearchSize(t *testing.T) {
	resp, err := urlscan.Search(context.Background(), urlscan.SearchArguments{
		Query: "ip:163.43.24.70",
		Size:  1,
	})

	require.NoError(t, err)
	assert.Equal(t, 1, len(resp.Results))
}

func TestSearchOffset(t *testing.T) {
	resp1, err := urlscan.Search(context.Background(), urlscan.SearchArguments{
		Query:  "ip:163.43.24.70",
		Size:   1,
		Offset: 0,
	})

	require.NoError(t, err)
	assert.Equal(t, 1, len(resp1.Results))

	resp2, err := urlscan.Search(context.Background(), urlscan.SearchArguments{
		Query:  "ip:163.43.24.70",
		Size:   1,
		Offset: 1,
	})

	require.NoError(t, err)
	assert.Equal(t, 1, len(resp2.Results))

	assert.NotEqual(t, resp1.Results[0].ID, resp2.Results[0].ID)
}
