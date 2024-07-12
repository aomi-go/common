package logger

var Factory = &factory{}

type Creator func(name string) Logger

type factory struct {
	creator Creator
}

func (f *factory) SetCreator(creator Creator) {
	f.creator = creator
}

func (f *factory) GetLogger(name string) Logger {
	if nil == f.creator {
		panic("logger factory creator is nil, please set it first: logger.Factory.SetCreator()")
	}
	return f.creator(name)
}
