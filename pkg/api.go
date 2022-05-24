package pkg

func (s *NetCardInfo) Get() error {
	s.file = "/home/john/wsp/gowp/netcardmng/ifcfg-enp3s0"
	err := s.read()
	if err != nil {
		return err
	}
	return nil
}

func (s *NetCardInfo) Flush(info *NetCardInfo) error {
	return nil
}
