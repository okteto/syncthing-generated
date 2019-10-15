// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package events

import (
	"os"
	"strings"

	liblogger "github.com/syncthing/syncthing/lib/logger"
)

var (
	dl = liblogger.DefaultLogger.NewFacility("events", "Event generation and logging")
)

func init() {
	dl.SetDebug("events", strings.Contains(os.Getenv("STTRACE"), "events") || os.Getenv("STTRACE") == "all")
}

// SetLogger overrides the existing logger
func SetLogger(n liblogger.Logger) {
	dl = n
}
