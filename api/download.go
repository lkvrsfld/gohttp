package api

import (
	"encoding/json"
	"fmt"
	sc "goscdl/sc"
	"io/ioutil"
	"net/http"
)
type Scurl struct {
	Url	string	`json:"url"`
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	
	// Read body
		
		body, err := getBody(w,r)

		if err != nil {
			fmt.Println("NO BODY DETECTED, LETS TRY PARAMS")
			param, errnew := getParams(w,r)
			if errnew != nil {
				http.Error(w, err.Error(), 500)
			}

			parsedString, statuscode := sc.ParseUrl(param, "params")
			if(statuscode == 200){
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(parsedString))
			}else{
				http.Error(w, parsedString, statuscode)
			}

	
			
		}else{

		parsedString, statuscode := sc.ParseUrl(body.Url, "body")
		if(statuscode == 200){
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(parsedString))
		}else{
			http.Error(w, parsedString, statuscode)
		}

		}
        return
	


}
func getBody(w http.ResponseWriter, r *http.Request) (Scurl, error) {
	var scurl Scurl
	b, err := ioutil.ReadAll(r.Body)


	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500) 
		return scurl, err
	}

	// Unmarshal
	
	err = json.Unmarshal(b, &scurl)
	if err != nil {
		return scurl, err
	}
	return scurl, nil
}

func getParams(w http.ResponseWriter, r *http.Request)(string, error)  {
	fmt.Println("HERE IS THE PARAM")
	fmt.Println(r.URL.Query()["url"][0])
	return r.URL.Query()["url"][0], nil

}

