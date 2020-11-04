package database

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type DBlogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func NewDBLogger() *DBlogger {
	log.Info("Get new DB logger")
	return &DBlogger{
		SkipErrRecordNotFound: false,
	}
}

func (l *DBlogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	log.Info("set level ")
	return l
}

func (l *DBlogger) Info(ctx context.Context, s string, args ...interface{}) {
	log.Infof(s, args)
}

func (l *DBlogger) Warn(ctx context.Context, s string, args ...interface{}) {
	log.Warnf(s, args)
}

func (l *DBlogger) Error(ctx context.Context, s string, args ...interface{}) {
	log.Errorf(s, args)
}

func (l *DBlogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := log.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		log.WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		log.WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	log.WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
