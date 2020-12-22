# nestp
Pretty print json/xml in json.
Nestp is based on [go-prettyjson](https://github.com/hokaccha/go-prettyjson). 

## Example
```
$ echo '{"json":"{\"key\":\"val\"}","xml":"<?xml version=\"1.0\"?><foo>bar</foo>"}' | go run cmd/main.go
{
  "json": {
      "key": "val"
    },
  "xml": "
    <?xml version=\"1.0\"?>
    <foo>bar
    </foo>"
}
```
