# Golang Urls Sanitize

[![Build Status](https://travis-ci.org/mantyr/urls.svg?branch=master)](https://travis-ci.org/mantyr/urls)
[![GoDoc](https://godoc.org/github.com/mantyr/urls?status.png)](http://godoc.org/github.com/mantyr/urls)
[![Software License](https://img.shields.io/badge/license-The%20Not%20Free%20License,%20Commercial%20License-brightgreen.svg)](LICENSE.md)

This stable version

## Installation

    $ go get github.com/mantyr/urls

## Example

```GO
package main

import (
    "github.com/mantyr/urls"
)

func main() {
    u := urls.New()
    u.SetBase("http://example.com/path")
    u.SetAddress("http://example.org/path/file")

    u.CheckScheme("http", "https", "")
    u.CheckHost()

    u.Is()
    u.IsHost()
    u.IsPath()
    u.IsFile("jpeg", "doc")
    u.String()
    // more example in url_test.go
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
