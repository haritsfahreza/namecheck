package namecheck

const (
	//StatusAvailable represents status when
	//the name is available on a channel
	StatusAvailable string = "V"

	//StatusNotAvailable represents status when
	//the name is not available on a channel
	StatusNotAvailable string = "X"

	//StatusUnknown represents status when there is
	//an unknown problem when checking the availability of the name
	StatusUnknown string = "?"
)

//Channel represents channel that will be checked
type Channel struct {
	Code   string
	URL    string
	Status string
	Error  error
}

//DefaultChannels contains list of available channel
var DefaultChannels = []Channel{{
	Code: ".com",
	URL:  "http://{name}.com",
}, {
	Code: ".co",
	URL:  "http://{name}.co",
}, {
	Code: ".co.id",
	URL:  "http://{name}.co.id",
}, {
	Code: ".id",
	URL:  "http://{name}.id",
}, {
	Code: ".info",
	URL:  "http://{name}.info",
}, {
	Code: ".io",
	URL:  "http://{name}.io",
}, {
	Code: ".me",
	URL:  "http://{name}.me",
}, {
	Code: ".name",
	URL:  "http://{name}.name",
}, {
	Code: ".net",
	URL:  "http://{name}.net",
}, {
	Code: ".org",
	URL:  "http://{name}.org",
}, {
	Code: ".us",
	URL:  "http://{name}.us",
}, {
	Code: ".xyz",
	URL:  "http://{name}.xyz",
}, {
	Code: "Facebook",
	URL:  "https://facebook.com/{name}",
}, {
	Code: "YouTube",
	URL:  "https://youtube.com/{name}",
}, {
	Code: "Instagram",
	URL:  "https://instagram.com/{name}",
}, {
	Code: "Twitter",
	URL:  "https://twitter.com/{name}",
}, {
	Code: "Google+",
	URL:  "https://plus.google.com/+{name}",
}, {
	Code: "Github",
	URL:  "https://github.com/{name}",
}}
