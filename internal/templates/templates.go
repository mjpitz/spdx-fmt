// Copyright (C) 2023 Mya Pitzeruse
// SPDX-License-Identifier: MIT

package templates

import (
	"embed"
)

//go:embed *.md.tmpl
var fs embed.FS

// Lookup attempts to resolve a template name from the in-memory file systems.
func Lookup(name string) ([]byte, bool) {
	contents, err := fs.ReadFile(name + ".md.tmpl")

	return contents, err == nil
}
