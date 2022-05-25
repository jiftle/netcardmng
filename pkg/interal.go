package pkg

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// read file from file
func (s *NetCardInfo) read() error {
	byt, err := ioutil.ReadFile(s.File)
	if err != nil {
		return err
	}

	arLine := strings.Split(string(byt), "\n")

	for _, line := range arLine {
		arKv := strings.Split(line, "=")
		if len(arKv) == 2 {
			k := TrimStr(arKv[0])
			v := TrimStr(arKv[1])
			if k == "IPADDR" {
				s.IpAddr = v
			} else if k == "PREFIX" {
				nv := TrimStr(v)
				nn, err := strconv.ParseUint(nv, 10, 32)
				if err != nil {
					return err
				}
				s.Prefix = int32(nn)
				s.PrefixStr = LenToSubNetMask(int(nn))
			} else if k == "GATEWAY" {
				s.Gateway = v
			} else if k == "DNS1" {
				s.DNS1 = v
			} else if k == "DNS2" {
				s.DNS2 = v
			} else if k == "BOOTPROTO" {
				s.BootProto = v
			} else if k == "ONBOOT" {
				s.OnBoot = ternary(v == "yes", true, false)
			} else if k == "DEVICE" {
				s.Device = v
			} else if k == "UUID" {
				s.UUID = v
			} else if k == "NAME" {
				s.Name = v
			} else if k == "TYPE" {
				s.Type = v
			}
		}
	}
	if s.BootProto == "dhcp" {
		s.readDNS()
	}

	return nil
}

// 三目运算的函数
func ternary(a bool, b, c bool) bool {
	if a {
		return b
	}
	return c
}

func TrimStr(s string) string {
	ss := strings.Trim(s, " ")
	ch := ss[0:1]
	if ch == "\"" {
		ns := ss[1 : len(ss)-1]
		return ns
	} else {
		return ss
	}
}

func (s *NetCardInfo) write() error {
	byt, err := ioutil.ReadFile(s.File)
	if err != nil {
		return err
	}

	arLine := strings.Split(string(byt), "\n")
	arNewLine := make([]string, 0)

	for _, line := range arLine {
		nLine := line
		arKv := strings.Split(line, "=")
		if len(arKv) == 2 {
			k := TrimStr(arKv[0])
			if k == "IPADDR" {
				vv := s.IpAddr
				nLine = fmt.Sprintf("%s=%v", k, vv)
			} else if k == "PREFIX" {
				vv := int32(s.Prefix)
				nLine = fmt.Sprintf("%s=%v", k, vv)
			} else if k == "GATEWAY" {
				vv := s.Gateway
				nLine = fmt.Sprintf("%s=%v", k, vv)
			} else if k == "DNS1" {
				vv := s.DNS1
				nLine = fmt.Sprintf("%s=%v", k, vv)
			}
		}
		arNewLine = append(arNewLine, nLine)
	}

	content := ""
	content = strings.Join(arNewLine, "\n")

	err = ioutil.WriteFile(s.File, []byte(content), 0655)
	if err != nil {
		return err
	}

	return nil
}
