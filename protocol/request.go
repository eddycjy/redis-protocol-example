package protocol

import (
	"strconv"
	"strings"
)

func GetRequest(args []string) []byte {
	req := []string{
		"*" + strconv.Itoa(len(args)),
	}

	for _, arg := range args {
		req = append(req, "$"+strconv.Itoa(len(arg)))
		req = append(req, arg)
	}

	str := strings.Join(req, "\r\n")
	return []byte(str + "\r\n")
}
