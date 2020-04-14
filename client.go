package picpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

// APIClient - struct client
type APIClient struct {
	client *http.Client
	Env    string
	Token  string
}

// Error - struct error
type Error struct {
	Message string       `json:"message"`
	Errors  []*ErrorItem `json:"errors"`
	Data    string       `json:"data"`
}

type ErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

//New - create a new client
func New(token, env string) *APIClient {
	return &APIClient{
		client: &http.Client{Timeout: 60 * time.Second},
		Env:    env,
		Token:  token,
	}
}

//Request - In webservice
func (client *APIClient) Request(method, action string, body interface{}, query interface{}, out interface{}) (error, *Error) {
	if client.client == nil {
		client.client = &http.Client{Timeout: 60 * time.Second}
	}

	dataData, err := json.Marshal(body)
	if err != nil {
		return err, nil
	}
	endpoint := fmt.Sprintf("%s%s", client.devProd(), action)
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(dataData))
	if err != nil {
		return err, nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-picpay-token", client.Token)
	if query != nil {
		q := url.Values{}
		queryStruct := structToMap(query)
		for k, v := range queryStruct {
			q.Add(k, fmt.Sprintf("%#v", v))
		}
		req.URL.RawQuery = q.Encode()
	}
	res, err := client.client.Do(req)
	if err != nil {
		return err, nil
	}

	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
		var errAPI Error
		err = json.Unmarshal(bodyResponse, &errAPI)
		if err != nil {
			return err, nil
		}
		errAPI.Data = string(bodyResponse)
		return nil, &errAPI
	}
	err = json.Unmarshal(bodyResponse, out)
	if err != nil {
		return err, nil
	}
	return nil, nil
}

//devProd - check type Env
func (client *APIClient) devProd() string {
	if client.Env == "develop" {
		return "https://appws.picpay.com"
	}
	return "https://appws.picpay.com"
}

func structToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
