package urlscan

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GetReportByUUID Gets a urlscan report by the task uuid
func GetReportByUUID(ctx context.Context, uuid string) (*ScanResult, error) {
	result := &ScanResult{}
	code, err := req(ctx, http.MethodGet, nil, fmt.Sprintf("result/%s", uuid), nil, result, "")
	if err != nil {
		return nil, errors.Wrap(err, "Fail to get result query")
	}

	switch code {
	case 200:
		return result, nil
	default:
		return nil, fmt.Errorf("status: %d", code)
	}
}
