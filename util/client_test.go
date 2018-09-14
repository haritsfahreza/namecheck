package util_test

import (
	"testing"

	"github.com/uwuh/namecheck/util"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func init() {
	httpmock.Activate()
}

func TestIsConnected(t *testing.T) {
	defer httpmock.Deactivate()

	t.Run("success", func(t *testing.T) {
		httpmock.Reset()
		httpmock.RegisterResponder("GET", "https://www.google.com",
			httpmock.NewStringResponder(200, ""))

		if result := util.IsConnected(); !result {
			t.Errorf("Must be connected, expected: %v actual: %v", true, result)
		}
	})

	t.Run("failed to connect", func(t *testing.T) {
		httpmock.Reset()
		httpmock.RegisterResponder("GET", "https://www.google.com", httpmock.ConnectionFailure)

		if result := util.IsConnected(); result {
			t.Errorf("Must be disconnected, expected: %v actual: %v", false, result)
		}
	})
}
