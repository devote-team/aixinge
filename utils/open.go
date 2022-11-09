package utils

import (
	"aixinge/global"
	"fmt"
)

func Open() {
	uri := `http://`
	ip, err := ExternalIP()
	if err == nil {
		uri += ip.String()
	} else {
		uri += `localhost`
	}
	uri += fmt.Sprintf("%s%d", `:`, global.CONFIG.System.Port)
	OpenUri(uri)
}
