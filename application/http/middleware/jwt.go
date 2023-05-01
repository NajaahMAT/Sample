package middleware

import (
	"Sample/application/config"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	kit "github.com/go-kit/kit/transport/http"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"log"
	"net/http"
	"strings"
)

var conf = config.ParseConfig("./configuration")
var secretKey = []byte(conf.AppConfig.SecretKey)

type AccessInfo struct {
	AccessUuid   string  		`json:"access_uuid"`
	UserId       string  		`json:"userid"`
	Exp          interface{}    `json:"exp"`
	Authorized   bool   		`json:"authorized"`
}

func BindHttpTokenWithContext() kit.RequestFunc{
	return func(ctx context.Context, req *http.Request) context.Context{
		var tokenForContext string

		c := context.Background()
		uuid, _:= uuid.New()
		ct := context.WithValue(c,"uuid" , uuid)
		header := req.Header.Get("Authorization")

		token := strings.Split(header, " ")

		if len(token) != 2 && strings.ToLower(token[0]) != "bearer"{
			log.Print("Bearer token not exist")
		}
		if len(token) != 2  && strings.ToLower(token[0]) == "bearer"{
			log.Print("token part missing")
		}
		if len(token) != 2{
			log.Print("not a valid token")
		}

		if len(token) == 2 && strings.ToLower(token[0]) == "bearer"{
			tokenForContext = token[1]
		}else{
			tokenForContext = ""
		}

		return context.WithValue(ct, "BearerToken", tokenForContext)
	}
}

func AuthParser()endpoint.Middleware{
	return func(next endpoint.Endpoint) endpoint.Endpoint{
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			//extracting token from context
			token,ok := ctx.Value("BearerToken").(string)
			if !ok{
				errorMsg := "No auth token in the request"
				return nil, fmt.Errorf(errorMsg)
			}

			// check if the token is empty
			if token == "" {
	    		// If we get here, the required token is missing
				errorMsg := "Required authorization token not found"
				return nil, fmt.Errorf(errorMsg)
			}

			parsedToken, err := jwt.Parse(token, func(token *jwt.Token)(interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
					errorMsg := "Invalid signing method"
					return nil, fmt.Errorf(errorMsg)
				}

				if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
					errorMsg := "Token expired"
					return nil, fmt.Errorf(errorMsg)
				}

				return secretKey, nil
			})
			if err != nil{
				return nil, fmt.Errorf("error parsing token: %w", err)
			}

			tokenDetails, err := ExtractTokenInfo(parsedToken)
			if err != nil{
				errorMsg := "error occurred while extracting token meta data"
				return nil, fmt.Errorf(errorMsg)
			}

			tokenRemainingTimeValidity := getTokenRemainingValidity(tokenDetails.Exp)
			if tokenRemainingTimeValidity <= 0 {
				errorMsg := "Token has expired"
				return nil, fmt.Errorf(errorMsg)
			}

			if !tokenDetails.Authorized {
				errorMsg := "Token not authorized"
				return nil, fmt.Errorf(errorMsg)
			}

			if len(tokenDetails.AccessUuid) == 0 || len(tokenDetails.UserId) == 0{
				errorMsg := "Invalid payload in the token"
				return nil, fmt.Errorf(errorMsg)
			}

			return next(ctx, request)
		}
	}
}

func ExtractTokenInfo(token *jwt.Token) (tokenDetails *AccessDetails, err error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, ok :=  claims["userid"].(string)
		if !ok {
			return nil, err
		}

		isAuthorized, ok := claims["authorized"].(bool)
		if !ok {
			return nil, err
		}

		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:   userId,
			Exp   : claims["exp"],
			Authorized: isAuthorized,
		}, nil
	}

	return nil, err
}