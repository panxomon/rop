package railway

import "github.com/go-kit/kit/endpoint"

func (r *Railway) Add(name string, service Service) {
	r.services[name] = service
}

type Railway struct {
	services map[string]Service

	endpoints map[string]endpoint.Endpoint

	middlewares []Middleware
}

func New() *Railway {
	return &Railway{
		services:  make(map[string]Service),
		endpoints: make(map[string]endpoint.Endpoint),
	}
}

type Middleware func(Service) Service

func Chain(middlewares ...Middleware) Middleware {
	return func(next Service) Service {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
