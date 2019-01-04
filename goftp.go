// Copyright 2015 Muir Manders.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package goftp provides a high-level FTP client for go.
*/
package goftp

import (
	"regexp"
)

// Dial creates an FTP client using the default config. See DialConfig for
// information about "hosts".
func Dial(hosts ...string) (*Client, error) {
	return DialConfig(Config{}, hosts...)
}

// DialConfig creates an FTP client using the given config. "hosts" is a list
// of IP addresses or hostnames with an optional port (defaults to 21).
func DialConfig(config Config, hosts ...string) (*Client, error) {
	return newClient(config, hosts), nil
}

var hasPort = regexp.MustCompile(`^[^:]+:\d+$|\]:\d+$`)
