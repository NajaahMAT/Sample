package usecases

import "Sample/domain/interfaces/repositories"

type GetSalesInteractor struct{
	Salesrepo  repositories.SalesRepoInterface
}

type GetProductSalesInteractor struct{
	ProductSalesRepo  repositories.ProductSalesRepoInterface
}

type SampleInteractor struct{
	SampleRepo   repositories.SampleRepositoryInterface
}