package protocol

import (
	"strconv"
	"strings"
)

const (
	StatusReply    = '+'
	ErrorReply     = '-'
	IntegerReply   = ':'
	BulkReply      = '$'
	MultiBulkReply = '*'

	OkReply   = "OK"
	PongReply = "PONG"
)

func GetReply(reply []byte) (interface{}, error) {
	replyType := reply[0]
	switch replyType {
	case StatusReply:
		return doStatusReply(reply[1:])
	case ErrorReply:
		return doErrorReply(reply[1:])
	case IntegerReply:
		return doIntegerReply(reply[1:])
	case BulkReply:
		return doBulkReply(reply[1:])
	case MultiBulkReply:
		return doMultiBulkReply(reply[1:])
	default:
		return nil, nil
	}
}

func doStatusReply(reply []byte) (string, error) {
	if len(reply) == 3 && reply[1] == 'O' && reply[2] == 'K' {
		return OkReply, nil
	}

	if len(reply) == 5 && reply[1] == 'P' && reply[2] == 'O' && reply[3] == 'N' && reply[4] == 'G' {
		return PongReply, nil
	}

	return string(reply), nil
}

func doErrorReply(reply []byte) (string, error) {
	return string(reply), nil
}

func doIntegerReply(reply []byte) (int, error) {
	pos := getFlagPos('\r', reply)
	result, err := strconv.Atoi(string(reply[:pos]))
	if err != nil {
		return 0, err
	}

	return result, nil
}

func doBulkReply(reply []byte) (interface{}, error) {
	pos := getFlagPos('\r', reply)
	pstart := 0
	if reply[:pos][0] == '$' {
		pstart = 1
	}

	vlen, err := strconv.Atoi(string(reply[pstart:pos]))
	if err != nil {
		return nil, err
	}
	if vlen == -1 {
		return nil, nil
	}

	start := pos + 2
	end := start + vlen
	return string(reply[start:end]), nil
}

func doMultiBulkReply(reply []byte) (interface{}, error) {
	replyStrs := strings.Split(string(reply), "\r\n")
	replylen := len(replyStrs)
	replyStrs = replyStrs[1 : replylen-1]

	r := []interface{}{}
	for i := 0; i < replylen-1; i++ {
		if i%2 == 1 {
			rv := strings.Join([]string{
				replyStrs[i-1],
				replyStrs[i],
			}, "\r\n") + "\r\n"

			value, err := doBulkReply([]byte(rv))
			if err != nil {
				return nil, err
			}

			r = append(r, value)
		}
	}

	return r, nil
}

func getFlagPos(flag byte, reply []byte) int {
	pos := 0
	for _, v := range reply {
		if v == flag {
			break
		}
		pos++
	}

	return pos
}
