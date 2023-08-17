package utils

import (
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var linkRegexPattern = regexp.MustCompile(`((ftp|http|https):\/\/)(www.)?(?P<domain>(?:[a-z\d](?:[a-z\d-]*[a-z\d])?\.)+[a-z]{2,}|((\d{1,3}\.){3}\d{1,3}))(.+)?`)

type UrlValidator struct {
	blackListSet map[string]bool
}

func NewURLValidator() (*UrlValidator, error) {
	blackList := make(map[string]bool)
	validator := &UrlValidator{
		blackListSet: blackList,
	}
	f, err := os.Open("./block.txt")
	defer f.Close()
	if err != nil {
		return validator, nil
	}
	domains, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	for _, link := range strings.Split(string(domains), "\n") {
		if link == "" {
			continue
		}
		blackList[link] = true
	}
	log.Println("Blacklist loaded", len(blackList))
	return validator, nil
}

func (u *UrlValidator) Validate(url string) error {
	match := linkRegexPattern.FindStringSubmatch(url)
	if len(match) == 0 {
		return errors.New("invalid URL")
	}
	domain := match[4]
	if u.blackListSet[domain] {
		return errors.New("this url is not available")
	}
	return nil
}
