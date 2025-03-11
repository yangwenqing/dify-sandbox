//go:build linux && loong64

package static

var DEFAULT_PYTHON_LIB_REQUIREMENTS = []string{
	"/usr/local/lib/python3.10",
    "/usr/lib/python3.10",
    "/usr/lib/python3",
	"/usr/lib/loongarch-unknown-linux-gnu",
	"/etc/ssl/certs/ca-certificates.crt",
    "/etc/nsswitch.conf",
    "/etc/hosts",
    "/etc/resolv.conf",
    "/run/systemd/resolve/stub-resolv.conf",
    "/run/resolvconf/resolv.conf",
    "/etc/localtime",
    "/usr/share/zoneinfo",
    "/etc/timezone",
}

