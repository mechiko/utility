package utility

import "net"

// GetFreePort asks the kernel for a currently available TCP port on 127.0.0.1.
// Note: The port is no longer reserved once returned; callers should bind immediately
// (or consider a function that returns a bound listener).
func GetFreePort() (port int, err error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
