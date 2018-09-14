package namecheck_test

import (
	"errors"
	"testing"

	"github.com/uwuh/namecheck"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func init() {
	httpmock.Activate()
}

func TestCheck(t *testing.T) {
	defer httpmock.Deactivate()

	mockChannels := []namecheck.Channel{{
		Code: ".com",
		URL:  "http://{name}.com",
		Type: namecheck.TypeDomain,
	}, {
		Code: "found",
		URL:  "https://found.com/{name}",
		Type: namecheck.TypeSocial,
	}, {
		Code: "notfound",
		URL:  "https://notfound.com/{name}",
		Type: namecheck.TypeSocial,
	}, {
		Code: "timeout",
		URL:  "https://timeout.com/{name}",
		Type: namecheck.TypeSocial,
	}, {
		Code: "error",
		URL:  "https://error.com/{name}",
		Type: namecheck.TypeSocial,
	}}

	httpmock.Reset()
	httpmock.RegisterResponder("GET", "http://test.com", httpmock.NewErrorResponder(errors.New("no such host")))
	httpmock.RegisterResponder("GET", "https://found.com/test", httpmock.NewStringResponder(200, "{}"))
	httpmock.RegisterResponder("GET", "https://notfound.com/test", httpmock.NewStringResponder(404, "{}"))
	httpmock.RegisterResponder("GET", "https://timeout.com/test", httpmock.NewErrorResponder(errors.New("Timeout")))
	httpmock.RegisterResponder("GET", "https://error.com/test", httpmock.NewErrorResponder(errors.New("test")))

	results, _ := namecheck.Check("test", mockChannels)

	if len(results) != len(mockChannels) {
		t.Errorf("Total results must be equal with channels. expected: %d actual: %d", len(mockChannels), len(results))
		return
	}

	for _, result := range results {
		if (result.Code == ".com" || result.Code == "notfound") && result.Status != namecheck.StatusAvailable {
			t.Errorf("(%s) result.Status must be available. expected: %s actual: %s",
				result.Code, namecheck.StatusAvailable, result.Status)
		}

		if result.Code == "found" && result.Status != namecheck.StatusNotAvailable {
			t.Errorf("(found) result.Status must be not available. expected: %s actual: %s",
				namecheck.StatusNotAvailable, result.Status)
		}

		if result.Code == "timeout" || result.Code == "error" {
			if result.Status != namecheck.StatusUnknown {
				t.Errorf("(%s) result.Status must be unknown. expected: %s actual: %s",
					result.Code, namecheck.StatusUnknown, result.Status)
			}

			if result.Error == nil {
				t.Errorf("(%s) result.Error must be not nil", result.Code)
			}
		}
	}

}
