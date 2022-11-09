package utils

import (
	"log"
	"testing"
)

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	log.Println(ip.String(), err)
}
