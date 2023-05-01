package repositories

import "Sample/domain/entities"

type SampleRepositoryInterface  interface{
	Get()(entities.Sample, error)
	Create(entities.Sample) (err error)
	Update(entities.Sample)(err error)
	Delete(entities.Sample)(err error)
}