package main

import (
	"net/http"
)

type Customer struct{
	Name string `json:"name"`
	Email string `json:"email"`
	City string `json:"city"`
}

type Customers []Customer

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
}


func (this *Response) SetStatus(status string){
	this.Status = status
}

func (this *Response) SetMessage(msg string){
	this.Message = msg
}