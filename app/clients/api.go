package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	circuit "github.com/rubyist/circuitbreaker"
)

type API struct{
	client *circuit.HTTPClient
}

func (r *API) Init() *API {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	r.client = circuit.NewHTTPClient(time.Second * 5, 10, client)
	return r
}

func (r *API) Call(url, method string) (*http.Response, error){
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
		return nil, err
	
	}
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// req.SetBasicAuth(username, password)

	response, err := r.client.Do(req)
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
		return nil, err
	}
	// body, err := io.ReadAll(response.Body)

	// defer response.Body.Close()
	return response,err
}

func (r *API) JSONCall(url, method string, target interface{} ) (error){
	res, err := r.Call(url, method)
	if err != nil {
		return err
	}
	defer res.Body.Close()
    return json.NewDecoder(res.Body).Decode(target)
}
