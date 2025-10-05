package main

import (
	"net/http"
	"net/http/cgi"
)

// First bad example - uses cgi.Serve
func badExample() {
	// ruleid: insecure-cgi-module
	// This should trigger the rule - uses net/http/cgi import and cgi.Serve
	cgi.Serve(http.FileServer(http.Dir("/usr/share/doc")))
}

// Second bad example - uses cgi.Serve with different parameters
func anotherBadExample() {
	// ruleid: insecure-cgi-module
	// This should also trigger the rule - uses cgi.Serve function
	cgi.Serve(http.FileServer(http.Dir("/tmp")))
}
