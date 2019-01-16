package main

import (

       "math/rand"
)
type Address struct{
	//declare empty struct
	account_addr []byte;
}
func NewAddress() Address {
	p := new(Address)
	token := make([]byte, 4)
	//generate random address
	rand.Read(token)
	p.account_addr  = token
	return *p
}
