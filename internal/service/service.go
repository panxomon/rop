package service

import "context"

type Service interface {
	Call(ctx context.Context, request interface{}) (response interface{}, err error)
}


type service struct {
	client client.CartolaSoap
}

type service struct {
	client client.CartolaSoap
}

func MakeService(client client.CartolaSoap) Service {
	return &service{
		client: client,
	}
}

func (s *service) Call(ctx context.Context, req *client.TransaccionesHistoricas) (*client.TransaccionesHistoricasResponse, error) {
	//DO NOTHING
	return s.client.TransaccionesHistoricas(req)
}
