package netcard

import (
	"encoding/json"
	"netcardmng/pkg"
)

var info pkg.NetCardInfo

func GetInfo() (string, error) {
	err := info.Get()
	if err != nil {
		return "", err
	}
	byt, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	return string(byt), nil
}

func SetIp(v string) (string, error) {
	info.IpAddr = v
	info.Flush(&info)
	return "", nil
}
func SetPrefix(v int32) (string, error) {
	info.Prefix = v
	info.Flush(&info)
	return "", nil
}
func SetGateway(v string) (string, error) {
	info.Gateway = v
	info.Flush(&info)
	return "", nil
}
func SetDNS1(v string) (string, error) {
	info.DNS1 = v
	info.Flush(&info)
	return "", nil
}
func SetDNS2(v string) (string, error) {
	info.DNS2 = v
	info.Flush(&info)
	return "", nil
}
