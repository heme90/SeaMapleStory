package static

type ServiceLogger struct {

}

func (s *ServiceLogger) Logf(msg string, i ...interface{}) {
	panic("implement me")
}

func (s *ServiceLogger) Log() bool {
	panic("implement me")
}

func (s *ServiceLogger) Fatalf(pos interface{}, msg string, args ...interface{}) {
	panic("implement me")
}

func (s *ServiceLogger) Warnl(pos interface{}, fmt_ string, args ...interface{}) {
	panic("implement me")
}

func (s *ServiceLogger) Debug_checknil() bool {
	panic("implement me")
}
