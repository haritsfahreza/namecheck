package namecheck

import (
	"errors"
	"net/http"
	"strings"

	"github.com/uwuh/namecheck/util"
)

func check(name, URL string) (string, error) {
	resp, err := util.HTTPClient.Get(strings.Replace(URL, "{name}", name, -1))
	if err != nil {
		if strings.Contains(err.Error(), "no such host") {
			return StatusAvailable, nil
		} else if strings.Contains(err.Error(), "Timeout") {
			return StatusUnknown, errors.New("Timeout")
		}

		return StatusUnknown, err
	}

	if resp.StatusCode == http.StatusOK {
		return StatusNotAvailable, nil
	}

	return StatusAvailable, nil
}

//Check name availability on all channel
func Check(name string) []Channel {
	channels := make([]Channel, len(Channels))

	for i, channel := range Channels {
		newChannel := channel
		newChannel.Status, newChannel.Error = check(name, channel.URL)
		channels[i] = newChannel
	}

	return channels
}
