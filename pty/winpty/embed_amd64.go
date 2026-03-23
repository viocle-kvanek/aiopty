//go:build windows && go1.16
// +build windows,go1.16

package winpty

import "embed"

//go:embed bin/x64/*
var f embed.FS
