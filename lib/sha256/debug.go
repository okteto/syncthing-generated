// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package sha256

import (
	"os"
	"strings"

	"github.com/syncthing/syncthing/lib/logger"
)

var (
	l = logger.DefaultLogger.NewFacility("sha256", "SHA256 hashing package")
)

func init() {
	l.SetDebug("sha256", strings.Contains(os.Getenv("STTRACE"), "beacon") || os.Getenv("STTRACE") == "all")
}

// SetLogger overrides the existing logger
func SetLogger(n logger.Logger) {
	l = n
}
