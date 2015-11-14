orderedmap.go
===============

[![Build Status](https://travis-ci.org/pocke/orderedmap.go.svg)](https://travis-ci.org/pocke/orderedmap.go)

Installation
-----------

```sh
go get github.com/pocke/orderedmap.go
```

Usage
-------

```go
m := omap.New()
json.NewDecoder(r).Decode(m)
json.NewEncoder(w).Encode(m)
```

License
-------

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.en)
