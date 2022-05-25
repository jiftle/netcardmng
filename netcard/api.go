package netcard

import (
	"encoding/json"

	"github.com/jiftle/netcardmng/pkg"
)

type NetCardInformation struct {
	MAC       string `json:"MAC"`
	IpAddr    string `json:"IpAddr"`
	PrefixStr string `json:"PrefixStr"`
	Gateway   string `json:"Gateway"`
	DNS1      string `json:"DNS1"`
	DNS2      string `json:"DNS2"`
}

var info pkg.NetCardInfo

func Set(file string) error {
	//etc/sysconfig/network-scripts/tmp-ifcfg-enp0s3
	info.File = file
	return nil
}

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

func SetNetCardInfo(v NetCardInformation) error {
	info.IpAddr = v.IpAddr
	vv, err := pkg.SubNetMaskToLen(v.PrefixStr)
	if err != nil {
		return err
	}
	info.Prefix = int32(vv)
	info.Gateway = v.Gateway
	info.DNS1 = v.DNS1
	info.Flush(&info)

	return nil
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
