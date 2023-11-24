package httpclient

import (
	"fmt"
	"net/http"
	"os"
	"time"
)
type httpClient struct {
	request *http.Request
	url string
	header string
	method string
	timout int
}

func (h* httpClient) New(requestURL string, headers string){
	h.header = headers
	h.url = requestURL
}

func (h* httpClient) Get(){
	h.method = http.MethodGet
}
func (h* httpClient) Exec() *http.Response{
	req, err := http.NewRequest(h.url, h.method, nil)
	if err != nil {
		panic("Error while making request");
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	return res
}







