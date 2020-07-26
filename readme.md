# Sitemap

A sitemap.xml parser for go

## Install

```bash
go get github.com/thxi/sitemap
```

## Examle

To fetch the sitemap:

```go
u := "https://myanimelist.net/sitemap/index.xml"
r, err := FetchFromURL(u)
// or use FetchFromURLwithClient()
// r, err := FetchFromURLwithClient(u, myCustomHttpClient)
if err != nil {
  panic(err)
}
defer r.Close()
// ...
```

Additionally, you can use Fetch(url).
This will firstly check **/robots.txt** for `Sitemap:...` section. It will default to **/sitemap.xml** if the section could not be found.

```go
u := "https://myanimelist.net"
r, err := Fetch(u)
if err != nil {
  panic(err)
}
defer r.Close()
// ...
```

All Fetch(...) methods will return `FetchError` if something goes wrong.
The error can be unwrapped for inspection by using `errors.Unwrap(err)`
