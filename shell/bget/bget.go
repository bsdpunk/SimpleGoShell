package bget

import "net/http"
import "io/ioutil"
import "log"

func Web(site string) string {
	resp, err := http.Get(site)

	if err != nil {
		log.Fatal(err)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(ret)
}
