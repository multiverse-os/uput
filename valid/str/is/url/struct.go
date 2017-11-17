package urls

import (
    "net/url"
)

type Urls struct {
    base    *url.URL
    address *url.URL
    Error   error
    is      bool
}