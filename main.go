package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/whilp/git-urls"
)

// Minimalistic binary to convert a url to json. The purpose is to avoid using (usually incorrect) regular expressions
// in bash scripts. The json output can then be processed with jq.
//
// Example:
// $ url-to-json git@github.com:alokmenghrajani/url-to-json.git | jq -r .host
// github.com

type jsonURL struct {
	Scheme   string     `json:"scheme"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Host     string     `json:"host"`
	Port     uint16     `json:"port"`
	Path     []string   `json:"path"`
	Query    url.Values `json:"query"`
	Fragment string     `json:"fragment"`
}

func main() {
	process(os.Args)
}

func process(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: url-to-json <url>")
		os.Exit(-1)
	}

	jsonUrl := urlToJson(args[1])
	jsonString, err := json.Marshal(jsonUrl)
	panicOnError(err)
	fmt.Println(string(jsonString))
}

func urlToJson(urlString string) jsonURL {
	u, err := url.Parse(urlString)
	if err != nil {
		// url.Parse fails on git urls, which are common. So try to parse with giturls.
		u, err = giturls.Parse(urlString)
		panicOnError(err)
	}

	// Convert the fields we care about from url.URL to JsonURL
	jsonUrl := jsonURL{
		Scheme:   u.Scheme,
		Username: u.User.Username(),
		Host:     u.Hostname(),
		Fragment: u.Fragment,
	}
	jsonUrl.Password, _ = u.User.Password()
	port, _ := strconv.Atoi(u.Port())
	jsonUrl.Port = uint16(port)

	// Split part, ignoring empty
	jsonUrl.Path = strings.FieldsFunc(u.Path, func(c rune) bool { return c == '/' })
	// Convert query string to map
	jsonUrl.Query, err = url.ParseQuery(u.RawQuery)
	panicOnError(err)

	return jsonUrl
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
