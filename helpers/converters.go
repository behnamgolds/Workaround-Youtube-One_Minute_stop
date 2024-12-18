package helpers

import (
	"errors"
	"regexp"
)

const YT_WATCH_URL = `watch\?v=`

var watch_regex = regexp.MustCompile(YT_WATCH_URL)

func Convert(input string) (string, error) {
	if !watch_regex.MatchString(input) {
		return "", errors.New("Input url is malformed")
	}
	result := watch_regex.ReplaceAllString(input, "embed/")
	return result, nil
}
