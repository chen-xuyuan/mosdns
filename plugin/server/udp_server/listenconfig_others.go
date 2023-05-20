//go:build !linux

package udp_server

import (
	"syscall"
)

func udpListenControl(_, _ string, c syscall.RawConn) error {
	return nil
}
