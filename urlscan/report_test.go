package urlscan

import (
	"context"
	"testing"
)

func TestGetReportByUUID(t *testing.T) {
	report, err := GetReportByUUID(context.Background(), "8c23ec03-64b2-442a-8db7-86681049f17e")
	if err != nil {
		t.Error(err)
		return
	}
	if report.Task.ScreenshotURL == "" {
		t.Errorf("no screenshot url")
	}
}
