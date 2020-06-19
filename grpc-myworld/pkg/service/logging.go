package service

import (
	"context"
	"kit/log"
	"kit/log/level"
	"time"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{level.Info(logger), s}
}

func (l *loggingService) Get(ctx context.Context, key string) (val string, err error) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Get",
			"key", key,
			"val", val,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return l.next.Get(ctx, key)
}

func (l *loggingService) Put(ctx context.Context, key, val string) (err error) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Put",
			"key", key,
			"val", val,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return l.next.Put(ctx, key, val)
}