package logging

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

type SentryLogger struct {
	*zap.SugaredLogger
}

func (l *SentryLogger) Error(args ...interface{}) {
	log := l.SugaredLogger
	e := sentryError(fmt.Errorf("%s", args...))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Error(args...)
}

func (l *SentryLogger) Errorf(template string, args ...interface{}) {
	log := l.SugaredLogger
	e := sentryError(fmt.Errorf(template, args...))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Errorf(template, args...)
}

func (l *SentryLogger) Errorln(args ...interface{}) {
	log := l.SugaredLogger
	e := sentryError(fmt.Errorf("%s\n", args...))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Errorln(args...)
}

func (l *SentryLogger) Errorw(msg string, keysAndValues ...interface{}) {
	log := l.SugaredLogger
	e := sentryError(fmt.Errorf(msg))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Errorw(msg, keysAndValues...)
}

func (l *SentryLogger) Warn(args ...interface{}) {
	log := l.SugaredLogger
	e := sentryWarn(getMessage("", args))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Warn(args...)
}

func (l *SentryLogger) Warnf(template string, args ...interface{}) {
	log := l.SugaredLogger
	e := sentryWarn(getMessage(template, args))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Warnf(template, args...)
}

func (l *SentryLogger) Warnln(args ...interface{}) {
	log := l.SugaredLogger
	e := sentryWarn(getMessage("", args))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Warnln(args...)
}

func (l *SentryLogger) Warnw(msg string, keysAndValues ...interface{}) {
	log := l.SugaredLogger
	e := sentryWarn(getMessage(msg, nil))
	if e != nil {
		log = l.SugaredLogger.With("sentryEvent", e)
	}
	log.Warnw(msg, keysAndValues...)
}

func getMessage(template string, fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, fmtArgs...)
	}

	if len(fmtArgs) == 1 {
		if str, ok := fmtArgs[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(fmtArgs...)
}

func sentryError(err error) *sentry.EventID {
	if err != nil {
		hub := sentry.CurrentHub().Clone()
		if hub == nil {
			return nil
		}
		return hub.CaptureException(err)
	}
	return nil
}

func sentryWarn(message string, args ...interface{}) *sentry.EventID {
	if message != "" {
		hub := sentry.CurrentHub().Clone()
		if hub == nil {
			return nil
		}
		client := hub.Client()
		if client == nil {
			return nil
		}
		event := client.EventFromMessage(message, sentry.LevelWarning)
		return hub.CaptureEvent(event)
	}
	return nil
}
