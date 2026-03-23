//go:build windows && go1.16
// +build windows,go1.16

package winpty

import "embed"

//go:embed bin/ia32/*
var f embed.FS
