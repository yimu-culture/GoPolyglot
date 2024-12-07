package common

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"strconv"
)

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func CommandParameterAdd(data ...interface{}) {
	for i := 0; i < len(data); i++ {
		CommandParams[data[i].(string)] = data[i+1]
		i = i + 1
	}

}

func CommandParameterGet(key string) interface{} {
	return CommandParams[key]
}

func IntToStr(i int) string {
	s := strconv.Itoa(i)

	return s
}
