package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SalesModel struct{
	ID     primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Sales  Sales           `json:"sales" bson:"sales"`
	CustomerInfo  []CustomerInfo   `json:"CustomerInfo"  bson:"CustomerInfo"`
}

type CustomerInfo struct{
	//ID     primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Customer    Customer        `json:"Customer" bson:"Customer"`
}

type Customer struct{
	Name     string    `json:"Name"  bson:"Name"`
}

type Sales struct {
	CustomerID         int          `json:"CustomerID" bson:"CustomerID,omitempty"`
	InvoiceNumber      string       `json:"InvoiceNo" bson:"InvoiceNo,omitempty"`
	TotalInvoiceValue  float64      `json:"TotalInvoiceValue" bson:"TotalInvoiceValue,omitempty"`
	SalesDate          time.Time    `json:"SalesDate" bson:"SalesDate,omitempty"`
	SalesOrderID       int          `json:"SalesOrderID" bson:"SalesOrderID,omitempty"`
	LocationID         int          `json:"LocationID" bson:"LocationID,omitempty"`
	OrganizationID     string       `json:"OrgID" bson:"OrgID,omitempty"`
}

type ProductSalesModel struct{
	ID                primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Sales             ProductSales         `json:"Sales" bson:"Sales"`
	SalesOrderDetails SalesOrderDetails    `json:"SalesOrderDetails" bson:"SalesOrderDetails"`
	Product           Product              `json:"ProductInfo" bson:"ProductInfo"`
}

type ProductSales struct{
	SalesDate          time.Time    `json:"SalesDate" bson:"SalesDate,omitempty"`
	LocationID         int          `json:"LocationID" bson:"LocationID,omitempty"`
	OrganizationID     string       `json:"OrgID" bson:"OrgID,omitempty"`
}

type SalesOrderDetails  struct{
   // Product           Product           `json:"ProductInfo" bson:"ProductInfo,omitempty"`
    ProductCategory   ProductCategory   `json:"ProductCategoryInfo" bson:"ProductCategoryInfo"`
}

type SellingPrices  struct{
	Sellingprice    float64    `json:"Sellingprice" bson:"Sellingprice,omitempty"`
}

type Product struct{
	ProductID    int       `json:"ProductID"  bson:"ProductID"`
	Name         string    `json:"Name"  bson:"Name"`
	SellingPricesArray   []SellingPrices  `json:"SellingPricesArray" bson:"SellingPricesArray"`
}

type ProductCategory struct{
	ProductCategory   PCategory  `json:"ProductCategory" bson:"ProductCategory"`
}

type PCategory struct{
	Category    string    `json:"Category"  bson:"Category,omitempty"`
}

