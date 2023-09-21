package logger

import (
	"context"
	"gorm.io/gorm/logger"
	"log/slog"
	"time"
)

type GormLoggerAdapter struct {
	logger *slog.Logger
}

func NewGormLoggerAdapter(logger *slog.Logger) *GormLoggerAdapter {
	return &GormLoggerAdapter{logger: logger}
}

func (g *GormLoggerAdapter) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g *GormLoggerAdapter) Info(ctx context.Context, s string, i ...interface{}) {
	g.logger.InfoContext(ctx, s, i...)
}

func (g *GormLoggerAdapter) Warn(ctx context.Context, s string, i ...interface{}) {
	g.logger.WarnContext(ctx, s, i...)
}

func (g *GormLoggerAdapter) Error(ctx context.Context, s string, i ...interface{}) {
	g.logger.ErrorContext(ctx, s, i...)
}

func (g *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rowsAffected := fc()
	g.logger.DebugContext(ctx,
		"gorm",
		"begin", begin,
		"sql", sql,
		"rowsAffected", rowsAffected,
		"err", err,
	)
}
