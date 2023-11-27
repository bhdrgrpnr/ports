package logging

import (
	"context"
	"io/ioutil"
	"os"
	
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func SetLogOutput() {
	log.SetOutput(ioutil.Discard)

	// Send high level logs to stderr
	log.AddHook(&writer.Hook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})

	// Send info and debug logs to stdout
	log.AddHook(&writer.Hook{
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
		},
	})
}

func LogRequest(ctx context.Context, handler string, body log.Fields, err bool) {
	if body == nil {
		body = make(log.Fields)
	}
	body["handler"] = handler

	if err {
		log.WithFields(body).Error("REQUEST")
	} else {
		log.WithFields(body).Info("REQUEST")
	}
}

func LogRequestError(ctx context.Context, handler string, err error) {
	body := make(log.Fields)
	body["error"] = err.Error()

	LogRequest(ctx, handler, body, true)
}

func LogInfo(ctx context.Context, scope string, msg string) {
	fields := make(log.Fields)

	fields["scope"] = scope

	log.WithFields(fields).Info(msg)
}

func LogWarn(ctx context.Context, scope string, msg string) {
	fields := make(log.Fields)

	fields["scope"] = scope


	log.WithFields(fields).Warn(msg)
}

func LogError(ctx context.Context, scope string, msg string, err error) {
	fields := make(log.Fields)

	fields["scope"] = scope
	fields["error"] = err.Error()

	log.WithFields(fields).Error(msg)
}

func LogFatal(ctx context.Context, scope string, msg string, err error) {
	fields := make(log.Fields)

	fields["scope"] = scope
	fields["error"] = err.Error()

	log.WithFields(fields).Fatal(msg)
}
