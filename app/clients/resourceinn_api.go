package clients

import (
	"github.com/rubyist/circuitbreaker"
	"time"
	"fmt"
	"net/http"
)

type ResourceinnAPI struct{
	client *circuit.HTTPClient
}

func (r *ResourceinnAPI) init(){
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	r.client = circuit.NewHTTPClient(time.Second * 5, 10, client)
}

func (r *ResourceinnAPI) call(url, method string) error{

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	// req.SetBasicAuth(username, password)

	response, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}