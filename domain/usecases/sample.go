package usecases

import (
	"Sample/domain/entities"
	"fmt"
)

type SampleServiceInterface interface {
	GetSampleInfoService(request entities.SampleReq)(response entities.SampleRes, err error)
	CreateSampleService(request entities.SampleReq)( err error)
	UpdateSampleService(request entities.SampleReq)( err error)
	DeleteSampleService(request entities.SampleReq)( err error)
}

type Sample struct{}

func(Sample) GetSampleService(request entities.SampleReq)(response entities.SampleRes, err error){
	fmt.Println("request: ", request)

	return entities.SampleRes{}, nil
}

func(Sample) CreateSampleService(request entities.SampleReq)( err error){
	fmt.Println("request: ", request)

	return  nil
}

func(Sample) UpdateSampleService(request entities.SampleReq)( err error){
	fmt.Println("request: ", request)

	return  nil
}

func(Sample) DeleteSampleService(request entities.SampleReq)( err error){
	fmt.Println("request: ", request)

	return  nil
}


