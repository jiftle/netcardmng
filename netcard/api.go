package netcard

import (
	"encoding/json"

	"github.com/jiftle/netcardmng/pkg"
)

type NetCardInformation struct {
	MAC       string
	IpAddr    string
	PrefixStr string
	Gateway   string
	DNS1      string
	DNS2      string
}

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

func GetNetCardInfo() (*NetCardInformation, error) {
	err := info.Get()
	if err != nil {
		return nil, err
	}
	s := &NetCardInformation{
		MAC:       info.MAC,
		IpAddr:    info.IpAddr,
		PrefixStr: info.PrefixStr,
		Gateway:   info.Gateway,
		DNS1:      info.DNS1,
		DNS2:      info.DNS2,
	}
	return s, nil
}

func SetIp(v string) error {
	info.IpAddr = v
	info.Flush(&info)
	return nil
}

func SetPrefix(v int32) error {
	info.Prefix = v
	info.Flush(&info)
	return nil
}

func SetPrefixStr(sv string) error {
	vv, err := pkg.SubNetMaskToLen(sv)
	if err != nil {
		return err
	}
	info.Prefix = int32(vv)
	info.Flush(&info)
	return nil
}

func SetGateway(v string) error {
	info.Gateway = v
	info.Flush(&info)
	return nil
}
func SetDNS1(v string) error {
	info.DNS1 = v
	info.Flush(&info)
	return nil
}
func SetDNS2(v string) error {
	info.DNS2 = v
	info.Flush(&info)
	return nil
}
