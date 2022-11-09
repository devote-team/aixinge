package utils

import (
	"errors"
	"net"
)

func ExternalIP() (net.IP, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, face := range faces {
		if face.Flags&net.FlagUp == 0 {
			continue
		}
		if face.Flags&net.FlagLoopback != 0 {
			continue
		}
		address, err := face.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range address {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		// not an ipv4 address
		return nil
	}
	return ip
}
