# nestp
Pretty print nested json/xml data in string.
Nestp is based on [go-prettyjson](https://github.com/hokaccha/go-prettyjson). 

## Example
```
$ echo '{"hoge":"<?xml version=\"1.0\"?><foo>bar</foo>"}' | go run cmd/main.go
{
  "hoge": "
    <?xml version=\"1.0\"?>
    <foo>bar
    </foo>"
}
```

## Support 
Currently only supported xml in JSON string.
