package api

import (
	"errors"
	"fmt"
	sc "goscdl/sc"
	"net/http"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	
	//here we check if our request has valid parameters.
	param, err := checkUrlParam(w,r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//Here we will parse it for Soundclouid
	parsedString, err := sc.ParseUrl(param)
	if(err != nil){
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(parsedString))
			

}







func checkUrlParam(w http.ResponseWriter, r *http.Request)(string, error)  {
	//fmt.Println(r.URL.Query()["url"][0])
	
	if(len(r.URL.Query()["url"]) < 1){
		fmt.Println("KEY IS NONEXISTENT")
		err := errors.New("no url parameter is given")
		return "", err
	}

	
	if(len(r.URL.Query()["url"][0]) < 1){
		fmt.Println("URLPARAM IS NULL")
		err := errors.New("no url was given")
		
		return "", err
	}

	return r.URL.Query()["url"][0], nil
}


