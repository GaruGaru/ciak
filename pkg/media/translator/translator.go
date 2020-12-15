package translator

import (
	"regexp"
	"strings"
)

var removeStrings = []string{
	"720p", "1080p", "x264", "BluRay", "mkv", "avi", "mp4",
	")", "(",
}

var nameRegex = regexp.MustCompile(`(.*?)(\d\d\d\d|S\d\dE\d\d)`)

var removeLogoRegex = regexp.MustCompile(`(\[.*?])`)

func Translate(name string) string {

	name = strings.Replace(name, ".", " ", -1)

	name = removeLogoRegex.ReplaceAllString(name, "")

	for _, rmv := range removeStrings {
		name = strings.Replace(name, rmv, "", -1)
	}

	allString := nameRegex.FindStringSubmatch(name)

	if len(allString) == 0 {
		return strings.TrimSpace(name)
	}

	return strings.TrimSpace(allString[1])

}
