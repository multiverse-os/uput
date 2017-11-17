package urls

import (
    "github.com/mantyr/runner"
    "strings"
    "net/url"
)

func New() (u *Urls) {
    u = new(Urls)
    u.is = true
    return
}

func (u *Urls) SetBase(address string) *Urls {
    address = runner.Trim(address)
    if address == "" {
        u.is = false
        return u
    }

    var err error
    u.base, err = url.Parse(address)
    if err != nil {
        u.Error = err
        u.is = false
    }
    return u
}

func (u *Urls) SetAddress(address string) *Urls {
    address = runner.Trim(address)
    if address == "" {
        u.is = false
        return u
    }

    var err error
    u.address, err = url.Parse(address)
    if err != nil {
        u.Error = err
        u.is = false
    }
    return u
}

// Example:
//  CheckSheme("https", "http", "")
func (u *Urls) CheckScheme(params ...string) *Urls {
    if !u.is {
        return u
    }
    if !runner.InSlice(u.address.Scheme, params) {
        u.is = false
    }
    if u.address.Scheme == "" {
        u.address.Scheme = u.base.Scheme
    }
    return u
}

func (u *Urls) CheckHost() *Urls {
    if !u.is {
        return u
    }
    if u.address.Host == "" {
        u.address.Host = u.base.Host
        u.address.Path = "/"+strings.TrimLeft(u.address.Path, "./")
    }
    return u
}

func (u *Urls) Is() bool {
    return u.is
}

func (u *Urls) IsHost() bool {
    if u.is && u.address.Host == u.base.Host {
        return true
    }
    return false
}

func (u *Urls) IsPath() bool {
    if !u.is || u.address.Path == "" || u.address.Path == "/" {
        return false
    }
    return true
}

// Example:
//  IsFile("doc", "tiff", "jpeg", "jpeg", "zip", "gz")
func (u *Urls) IsFile(params ...string) bool {
    if !u.is {
        return false
    }
    i := strings.LastIndex(u.address.Path, ".")
    var file_type string
    if i < 0 {
        return false
    }
    file_type = u.address.Path[i+1:]
    file_type = strings.ToLower(file_type)

    for _, f_type := range params {
        if file_type == strings.ToLower(f_type) {
            return true
        }
    }

    return false
}

func (u *Urls) String() string {
    if u.is {
        return u.address.String()
    }
    return ""
}