package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"grpc-world/pkg/encode"
	"grpc-world/pkg/service"
)

//终端层，主要实现各种接口的handler，负责req/resp格式的转换，主要是将transport
//传过来的request内容进行类型转换并将数据传到service层及处理返回的内容和转换，简单来说它就是transport 和
//service 的桥梁


type GetRequest struct{
	Key string `json:"key"`
	Val string `json:"val"`
}
type Endpoints struct{
	GetEndpoint endpoint.Endpoint
	PutEndpoint endpoint.Endpoint
}
func NewEndpoint(s service.Service, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{
		GetEndpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(GetRequest)
			val, err := s.Get(ctx, req.Key)
			return encode.Response{
				Error: err,
				Data:  val,
			}, err
		},
		PutEndpoint: func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(GetRequest)
			err = s.Put(ctx, req.Key, req.Val)
			return encode.Response{
				Error: err,
			}, err
		},
	}

	for _, m := range mdw["Get"] {
		eps.GetEndpoint = m(eps.GetEndpoint)
	}
	for _, m := range mdw["Put"] {
		eps.PutEndpoint = m(eps.PutEndpoint)
	}

	return eps
}
