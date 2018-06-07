package util

import (
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) (bytes []byte, err error) {

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return

}
