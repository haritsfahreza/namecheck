package namecheck_test

import (
	"testing"

	"github.com/uwuh/namecheck"
)

func TestFindChannelByCode(t *testing.T) {
	channel, err := namecheck.FindChannelByCode(nil, ".com")

	if err != nil {
		t.Errorf("Error must be nil. expected: %s actual: %s", "nil", err.Error())
		return
	}

	if channel.Code != ".com" {
		t.Errorf("Invalid channel code. expected: %s actual: %s", ".com", channel.Code)
		return
	}

	channel, err = namecheck.FindChannelByCode(nil, "invalidcode")
	if err == nil {
		t.Errorf("Error must not be nil. expected: %s actual: %v", "no channel found", err)
		return
	}
}

func TestFindChannelsByType(t *testing.T) {
	channels := namecheck.FindChannelsByType(nil, namecheck.TypeDomain)

	if len(channels) <= 0 {
		t.Errorf("Channels must not be empty. expected: %s actual: %d", "> 0", len(channels))
		return
	}

	for _, ch := range channels {
		if ch.Type != namecheck.TypeDomain {
			t.Errorf("Invalid channel type on channel %s. expected: %s actual: %s", ch.Code, namecheck.TypeDomain, ch.Type)
			return
		}
	}

}
