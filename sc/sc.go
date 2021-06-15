package sc

import (
//	"log"
//	"net/http"
	"net/url"
	"errors"
)

func ParseUrl(scurl string) (string, error){		
//start our url parser
	parse, err := url.Parse(scurl)
	if(err !=  nil){
		return "", errors.New("couldn't parse!")
	}

	scheme := parse.Scheme 
	host := parse.Host
	path := parse.Path


	//here we check for an existent host (no matter what)
	if(host == ""){
		return "", errors.New("URL is invalid")
	}

	//here we check if the url is https
	if(scheme != "https"){
		return "", errors.New("URL is invalid, try HTTPS")
	}

	//here we check for an existent directory/parameters
	if(path == ""){
		return "", errors.New("URL is invalid, nothing after the /") 
	}

// ToDo - implement soundcloud valid parsing


	parsedUrl := parse.String()

	return parsedUrl, nil
}



func getDownloadType(url string) string {

	


/*



// SEARCH A SINGLE SONG
https://soundcloud.com/tiapyne/carta-de-f-banda-fly

// THIS IS AN ARTIST
https://soundcloud.com/djneck

// PLAYLIST
https://soundcloud.com/nemobaert/sets/industrial-techno



// FROM FEATURED / HOME
https://soundcloud.com/discover/sets/personalized-tracks::luke-eversfield:1007363863
*/

}