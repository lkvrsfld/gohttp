package sc

import (
	"fmt"
//	"log"
//	"net/http"
	"net/url"
)

func ParseUrl(scurl string, t string) (string, int){		

	if(t == "body"){
		fmt.Println("BDOY")

	}
	if(t == "params"){
		fmt.Println("PARAMS")

	}

// check if there is even anything
	if(scurl == ""){
		fmt.Errorf("URL IS EMPTY")
	}
//start our url parser
	parse, err := url.Parse(scurl)
	if(err !=  nil){
		return err.Error(), 400
	}

	// scheme is "https://, http://, ftp://"
	scheme := parse.Scheme
	//host 
	host := parse.Host

	//path
	path := parse.Path

	if(host == ""){
		return "URL is invalid", 400
	}
	if(scheme != "https"){
		return "URL is invalid, try HTTPS", 400
	}
	if(path == ""){
		return "URL is invalid, nothing after the /", 400
	}






	parsedUrl := parse.String()

	fmt.Println(parsedUrl)
	return parsedUrl, 200
}



