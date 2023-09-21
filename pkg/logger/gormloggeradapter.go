package logger

import (
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type GormLoggerAdapter struct {
	logger Logger
}

func NewGormLoggerAdapter(logger Logger) *GormLoggerAdapter {
	return &GormLoggerAdapter{logger: logger}
}

func (g *GormLoggerAdapter) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *GormLoggerAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	g.logger.Infof(s, i...)
}

func (g *GormLoggerAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	g.logger.Warnf(s, i...)
}

func (g *GormLoggerAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	g.logger.Errorf(s, i...)
}

func (g *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.logger.GetLevel() > Debug {
		return
	}
	sql, rowsAffected := fc()
	g.logger.Debugf("begin: %s, sql: %s, rowsAffected: %d, err: %v", begin, sql, rowsAffected, err)
}
