//go:build linux || solaris
// +build linux solaris

package term

import "github.com/viocle-kvanek/aiopty/term/export"

const reqGetTermios = export.TCGETS
const reqSetTermios = export.TCSETS
