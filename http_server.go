package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/coin.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = t.Execute(w, t); err != nil {
		fmt.Println(err)
		return
	}
}

func coinHandler(w http.ResponseWriter, r *http.Request) {
	coin := r.PostFormValue("coin")
	infolevel := r.FormValue("infolevel")

	getCryptoInfo(coin, infolevel)

        t, err := template.ParseFiles("templates/ack.html")
        if err != nil {
                fmt.Println(err)
                return
        }

        if err = t.Execute(w, t); err != nil {
                fmt.Println(err)
                return
        }
}
