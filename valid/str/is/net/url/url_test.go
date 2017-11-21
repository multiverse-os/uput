package urls

import (
    "testing"
)

type ParseTest struct {
    name       string
    base       string
    address    string
    is         bool
    isHost     bool
    isPath     bool
    isFile     bool
    result     string
}

var parseTests = []ParseTest{
    {"", "http://example.com/path", "https://example.org/path", true, false, true, false, "https://example.org/path"},
    {"", "http://example.com/path", "https://example.com/path", true, true, true, false, "https://example.com/path"},
    {"", "http://example.com/path", "/path2", true, true, true, false, "http://example.com/path2"},
    {"", "https://example.com/path", "./path2", true, true, true, false, "https://example.com/path2"},
    {"", "", "./path2", false, false, false, false, ""},
    {"", "http://example.com/path", "http://example.com/path/file.doc", true, true, true, true, "http://example.com/path/file.doc"},
    {"", "http://example.com/path", "http://example.com/path/file.doc?test=1234", true, true, true, true, "http://example.com/path/file.doc?test=1234"},
    {"", "http://example.com/path", "http://example.com/path/file.avi", true, true, true, false, "http://example.com/path/file.avi"},
}


func TestNormal(t *testing.T) {
    for _, test := range parseTests {
        u := New()
        u.SetBase(test.base)
        u.SetAddress(test.address)

        u.CheckScheme("http", "https", "")
        u.CheckHost()

        if u.Is() != test.is {
            t.Errorf("%q Is expected %q got %q", test, test.is, u.Is())
        }
        if u.IsHost() != test.isHost {
            t.Errorf("%q IsHost expected %q got %q", test, test.is, u.IsHost())
        }
        if u.IsPath() != test.isPath {
            t.Errorf("%q IsPath expected %q got %q", test, test.is, u.IsPath())
        }
        if u.String() != test.result {
            t.Errorf("%q String expected %q got %q", test, test.result, u.String())
        }
        if u.IsFile("pdf", "jpeg", "tiff", "doc", "zip", "gz") != test.isFile {
            t.Errorf("%q IsFile expected %q got %q", test, test.result, u.IsFile())
        }
    }
}