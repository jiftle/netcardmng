package pkg

func (s *NetCardInfo) Get() error {
	err := s.read()
	if err != nil {
		return err
	}
	if s.BootProto == "dhcp" {
		ip, mask, mac, err := GetLocalAddr()
		if err != nil {
			return err
		}
		s.IpAddr = ip
		s.PrefixStr = mask
		nPrefix, _ := SubNetMaskToLen(mask)
		s.Prefix = int32(nPrefix)
		s.MAC = mac
	}
	return nil
}

func (s *NetCardInfo) Flush(info *NetCardInfo) error {
	return s.write()
}
