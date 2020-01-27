package main

import (
	"fmt"
	erpc2 "github.com/vahanerevan/erpc"
)

type ttt struct {
	Name    string  `json:"name"`
	Id      int     `json:"id"`
	Balance float64 `json:"balance"`
}


type ttt1 struct {
	Balance string `json:"balance"`
}


func main() {

	config := erpc2.RequestConfig{
		URL:    "http://localhost:3000",
		Secret: "testest",
	}

	r := erpc2.NewRequest(config)

	r.SetRequestObject(ttt{
		Name:    "user",
		Id:      212,
		Balance: 33.2,
	})

	err := r.Call()

	if nil != err {
		fmt.Println(err.Error())
	}


	var tp  ttt1;
	r.Response().ToJSON(&tp)

	fmt.Println(tp.Balance)

	fmt.Println(r.Response().Status)

}
