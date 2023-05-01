package entities

import "context"

type Response struct{
	Response  interface{}
	Ctx       context.Context
}