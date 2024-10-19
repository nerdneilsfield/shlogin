//go:build darwin
// +build darwin

package network

func ping_impl(host string) (string, error) {
	logger.Warn("ping_impl is not implemented for darwin")
	return "", nil
}
