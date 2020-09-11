# url-to-json
Minimalistic binary to convert a url to json. The purpose is to avoid using (often incorrect) regular expressions
in bash scripts. The json output can then be processed with [jq](https://stedolan.github.io/jq/).

Example:
```
$ url-to-json git@github.com:alokmenghrajani/url-to-json.git | jq -r .host
github.com
```

# Building
go build
