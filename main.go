package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/whilp/git-urls"
)

// Minimalistic binary to convert a url to json. The purpose is to avoid using (usually incorrect) regular expressions
// in bash scripts. The json output can then be processed with jq.
//
// Example:
// $ url-to-json git@github.com:alokmenghrajani/url-to-json.git | jq -r .host
// github.com

type JsonURL struct {
	Scheme   string     `json:"scheme"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Host     string     `json:"host"`
	Port     string     `json:"port"`
	Path     []string   `json:"path"`
	Query    url.Values `json:"query"`
	Fragment string     `json:"fragment"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: url-to-json <url>")
		os.Exit(-1)
	}

	u, err := url.Parse(os.Args[1])
	if err != nil {
		// url.Parse fails on git urls, which are common. So try to parse with giturls.
		u, err = giturls.Parse(os.Args[1])
		panicOnError(err)
	}

	jsonUrl := JsonURL{
		Scheme:   u.Scheme,
		Username: u.User.Username(),
		Host:     u.Hostname(),
		Port:     u.Port(),
		Fragment: u.Fragment,
	}
	jsonUrl.Password, _ = u.User.Password()
	jsonUrl.Path = strings.FieldsFunc(u.Path, func(c rune) bool { return c == '/' })
	jsonUrl.Query, err = url.ParseQuery(u.RawQuery)
	panicOnError(err)

	jsonString, err := json.Marshal(jsonUrl)
	panicOnError(err)

	fmt.Println(string(jsonString))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
