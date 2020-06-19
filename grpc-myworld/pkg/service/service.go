package service
//service是具体业务的实现
import (
	"Demo-project/grpc-myworld/pkg/repository"
	"golang.org/x/net/context"
	"kit/log"
)

type Service interface{
	Get(ctx context.Context,key string)(val string,err error)
	Put(ctx context.Context,key,val string)(err error)
}
type service struct {
	logger     log.Logger
	repository repository.Repository
}

func (s *service) Put(ctx context.Context, key, val string) (err error) {
	return s.repository.Put(key, val)
}

func (s *service) Get(ctx context.Context, key string) (val string, err error) {
	res, err := s.repository.Get(key)
	if err != nil {
		return
	}
	return res.Val, err
}

func New(logger log.Logger, repository repository.Repository) Service {
	return &service{logger: logger, repository: repository}
}
