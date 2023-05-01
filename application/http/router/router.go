package router

import (
	"Sample/application/http/endpoints"
	jwtmiddleware "Sample/application/http/middleware"
	httpReq "Sample/application/http/request"
	httpRes "Sample/application/http/response"
	"Sample/domain/usecases"
	"fmt"
	httpTrans "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	trans "github.com/go-kit/kit/transport/http"
)

func Run(port int) {

	app := Handlers()
	//
	//jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
	//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//		return []byte("My Secret"), nil
	//	},
	//	SigningMethod: jwt.SigningMethodHS256,
	//	EnableAuthOnOptions: true,
	//})

	//app := jwtMiddleware.Handler(Handlers())

	err := http.ListenAndServe(fmt.Sprintf(`:%d`, port), app)
	if err != nil{
		log.Fatalln("Error occurred while starting web server: ", err)
	}
}

func Handlers() *mux.Router{
    r := mux.NewRouter()

    options := []trans.ServerOption{
    	trans.ServerBefore(jwtmiddleware.BindHttpTokenWithContext()),
	}

	var salesService usecases.GetSalesInterface
    {
    	salesService = usecases.GetSales{}
	}

	var productSalesService usecases.GetProductSalesInterface
    {
		productSalesService = usecases.GetProductSales{}
	}

	var sampleService usecases.SampleServiceInterface
    {
    	//salesService = usecases.Sample{}
	}

	r.Handle("/v1/sales", httpTrans.NewServer(
	   jwtmiddleware.AuthParser()(endpoints.GetSalesEndpoint(salesService)),
	   httpReq.DecodeGetSalesRequest,
	   httpRes.EncodeGetSalesResponse,
	   options...,
	 )).Methods("GET")

	r.Handle("/v1/product_sales", httpTrans.NewServer(
		jwtmiddleware.AuthParser()(endpoints.GetProductSalesEndpoint(productSalesService)),
		httpReq.DecodeGetProductSalesRequest,
		httpRes.EncodeGetProductSalesResponse,
		options...,
	)).Methods("GET")

	r.Handle("/v1/product_sales/new", httpTrans.NewServer(
		jwtmiddleware.AuthParser()(endpoints.GetProductSalesEndpoint(productSalesService)),
		httpReq.DecodeGetProductSalesRequest,
		httpRes.EncodeGetProductSalesResponse,
		options...,
	)).Methods("GET")

    //sample endpoints
	r.Handle("/v1/sample", httpTrans.NewServer(
		endpoints.GetSamplesEndpoint(sampleService),
		httpReq.DecodeGetSampleRequest,
		httpRes.EncodeGetSampleResponse,
	)).Methods("GET")

    return r
}