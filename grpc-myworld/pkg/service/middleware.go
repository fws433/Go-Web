package service

import (
	"context"
	"errors"
	"golang.org/x/time/rate"
	"kit/endpoint"
)

var ErrLimitExceed = errors.New("Rate limit exceed!")

func TokenBucketLimitter(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}