# url-to-json
[![license](http://img.shields.io/badge/license-apache_2.0-blue.svg?style=flat)](https://raw.githubusercontent.com/alokmenghrajani/url-to-json/master/LICENSE) [![travis](https://img.shields.io/travis/alokmenghrajani/url-to-json/master.svg?maxAge=3600&logo=travis&label=travis)](https://travis-ci.org/alokmenghrajani/url-to-json) [![coverage](https://coveralls.io/repos/github/alokmenghrajani/url-to-json/badge.svg?branch=master)](https://coveralls.io/github/alokmenghrajani/url-to-json?branch=master) [![report](https://goreportcard.com/badge/github.com/alokmenghrajani/url-to-json)](https://goreportcard.com/report/github.com/alokmenghrajani/url-to-json)

Minimalistic binary to convert a url to json. The purpose is to avoid using (often incorrect) regular expressions
in bash scripts. The json output can then be processed with [jq](https://stedolan.github.io/jq/).

Example:
```
$ url-to-json git@github.com:alokmenghrajani/url-to-json.git | jq -r .host
github.com
```

# Building
go build
