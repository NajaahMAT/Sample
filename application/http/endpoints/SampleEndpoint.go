package endpoints

import (
	"Sample/domain/entities"
	"Sample/domain/usecases"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func GetSamplesEndpoint(svc usecases.SampleServiceInterface) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(entities.SampleReq)
		res, err := svc.GetSampleInfoService(req)
		return entities.Response{
			res,
			ctx,
		}, err
	}
}

func CreateSamplesEndpoint(svc usecases.SampleServiceInterface) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(entities.SampleReq)
		err = svc.CreateSampleService(req)
		return entities.Response{
			response,
			ctx,
		}, err
	}
}

func UpdateSamplesEndpoint(svc usecases.SampleServiceInterface) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(entities.SampleReq)
		err = svc.UpdateSampleService(req)
		return entities.Response{
			response,
			ctx,
		}, err
	}
}

func DeleteSamplesEndpoint(svc usecases.SampleServiceInterface) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(entities.SampleReq)
		err = svc.DeleteSampleService(req)
		return entities.Response{
			response,
			ctx,
		}, err
	}
}


