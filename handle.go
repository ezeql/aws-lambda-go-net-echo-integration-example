package main

import (
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Listener = net.Listen() //use a aws-lambda-go-net listener

	e.POST("/users", post)

	Handle = apigatewayproxy.New(e.Listener, nil).Handle //proxy using custom listener

	go e.Start("")

}

func post(c echo.Context) error {
	var u user

	if err := c.Bind(&u); err != nil {
		return err
	}
	return c.JSON(200, u)
}

type user struct {
	Name string
	Age  int
}
