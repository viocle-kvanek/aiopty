package winpty

import (
	"fmt"
	"os"
	"sync"

	"github.com/viocle-kvanek/aiopty/pty/common"
)

var errUnsupported = fmt.Errorf("unsupported os or arch")

type WinPty struct {
	opt      *common.Options
	pty      uintptr
	conin    *os.File
	conout   *os.File
	process  uintptr
	mu       sync.Mutex
	isClosed bool
}

// Open create a WinPty using path as the command to run.
func Open(path string) (*WinPty, error) {
	return openWithOptions(&common.Options{Path: path})
}

// OpenWithOptions create a WinPty with Options.
func OpenWithOptions(opt *common.Options) (*WinPty, error) {
	return openWithOptions(opt)
}

// SetSize is used to set the WinPty windows size.
func (p *WinPty) SetSize(size *common.WinSize) (err error) {
	return p.setSize(size)
}

// Close WinPty.
func (p *WinPty) Close() (err error) {
	return p.close()
}

// Read from WinPty.
func (p *WinPty) Read(b []byte) (n int, err error) {
	return p.read(b)
}

// Write to WinPty.
func (p *WinPty) Write(b []byte) (n int, err error) {
	return p.write(b)
}
