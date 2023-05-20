//go:build !linux

package coremain

import (
	"syscall"
)

func udpListenControl(_, _ string, c syscall.RawConn) error {
	return nil
}
