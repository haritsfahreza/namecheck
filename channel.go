package namecheck

import (
	"context"
	"errors"
)

//ChannelType value
type ChannelType string

//ChannelStatus value
type ChannelStatus string

const (
	//StatusAvailable represents status when
	//the name is available on a channel
	StatusAvailable ChannelStatus = "V"

	//StatusNotAvailable represents status when
	//the name is not available on a channel
	StatusNotAvailable ChannelStatus = "X"

	//StatusUnknown represents status when there is
	//an unknown problem when checking the availability of the name
	StatusUnknown ChannelStatus = "?"

	//TypeDomain represents DNS channel type
	TypeDomain ChannelType = "domain"

	//TypeSocial represents Social media channel type
	TypeSocial ChannelType = "social"
)

//Channel represents channel that will be checked
type Channel struct {
	Code   string
	URL    string
	Type   ChannelType
	Status ChannelStatus
	Error  error
}

//DefaultChannels contains list of available channel
var DefaultChannels = []*Channel{{
	Code: ".biz",
	URL:  "http://{name}.biz",
	Type: TypeDomain,
}, {
	Code: ".cc",
	URL:  "http://{name}.cc",
	Type: TypeDomain,
}, {
	Code: ".com",
	URL:  "http://{name}.com",
	Type: TypeDomain,
}, {
	Code: ".co",
	URL:  "http://{name}.co",
	Type: TypeDomain,
}, {
	Code: ".co.id",
	URL:  "http://{name}.co.id",
	Type: TypeDomain,
}, {
	Code: ".id",
	URL:  "http://{name}.id",
	Type: TypeDomain,
}, {
	Code: ".in",
	URL:  "http://{name}.info",
	Type: TypeDomain,
}, {
	Code: ".info",
	URL:  "http://{name}.info",
	Type: TypeDomain,
}, {
	Code: ".io",
	URL:  "http://{name}.io",
	Type: TypeDomain,
}, {
	Code: ".ly",
	URL:  "http://{name}.ly",
	Type: TypeDomain,
}, {
	Code: ".me",
	URL:  "http://{name}.me",
	Type: TypeDomain,
}, {
	Code: ".name",
	URL:  "http://{name}.name",
	Type: TypeDomain,
}, {
	Code: ".net",
	URL:  "http://{name}.net",
	Type: TypeDomain,
}, {
	Code: ".org",
	URL:  "http://{name}.org",
	Type: TypeDomain,
}, {
	Code: ".us",
	URL:  "http://{name}.us",
	Type: TypeDomain,
}, {
	Code: ".xyz",
	URL:  "http://{name}.xyz",
	Type: TypeDomain,
}, {
	Code: "Facebook",
	URL:  "https://facebook.com/{name}",
	Type: TypeSocial,
}, {
	Code: "YouTube",
	URL:  "https://youtube.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Instagram",
	URL:  "https://instagram.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Twitter",
	URL:  "https://twitter.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Github",
	URL:  "https://github.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Gitlab",
	URL:  "https://gitlab.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Bitbucket",
	URL:  "https://bitbucket.org/{name}",
	Type: TypeSocial,
}, {
	Code: "Reddit",
	URL:  "https://www.reddit.com/user/{name}",
	Type: TypeSocial,
}, {
	Code: "Medium",
	URL:  "https://www.medium.com/@{name}",
	Type: TypeSocial,
}, {
	Code: "Vimeo",
	URL:  "https://vimeo.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Artstation",
	URL:  "https://www.artstation.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Tumblr",
	URL:  "https://{name}.tumblr.com",
	Type: TypeSocial,
}, {
	Code: "Keybase.io",
	URL:  "https://keybase.io/{name}",
	Type: TypeSocial,
}, {
	Code: "Vkontakte",
	URL:  "https://vk.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Pinterest",
	URL:  "https://pinterest.com/{name}",
	Type: TypeSocial,
}, {
	Code: "Slack",
	URL:  "https://{name}.slack.com",
	Type: TypeSocial,
}, {
	Code: "Dev.to",
	URL:  "https://dev.to/{name}",
	Type: TypeSocial,
}}

//FindChannelByCode is used to find channel by its channel code
func FindChannelByCode(ctx context.Context, code string) (*Channel, error) {
	for _, channel := range DefaultChannels {
		if channel.Code == code {
			return channel, nil
		}
	}
	return nil, errors.New("No channel found. Please use the available channel code")
}

//FindChannelsByType is used to find channels by its channel type
func FindChannelsByType(ctx context.Context, chType ChannelType) []*Channel {
	result := []*Channel{}
	for _, channel := range DefaultChannels {
		if channel.Type == chType {
			result = append(result, channel)
		}
	}
	return result
}
