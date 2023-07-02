package logging

type Log interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugln(args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorln(args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalln(args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infoln(args ...interface{})
	Infow(msg string, keysAndValues ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Panicln(args ...interface{})
	Panicw(msg string, keysAndValues ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnln(args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
}
