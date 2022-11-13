package example

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orc-system/pkg/utils"
	"time"
)

type exampleService struct {
	endpoint   string
	httpClient *http.Client
}

func NewExampleService(endPoint string) IExample {
	return &exampleService{
		endpoint: endPoint,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (i *exampleService) GetList(param *ExpInput) (*ExpOutPut, error) {
	url := fmt.Sprintf("%v/example", i.endpoint)
	if param != nil {
		url = fmt.Sprintf("%v/example?id=%s", i.endpoint, param.ParamId)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := i.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if !utils.IsHTTPSuccess(resp.StatusCode) {
		// chinh sua thanh struct tuong ung voi error server tra ve
		var errHttp interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errHttp); err != nil {
			return nil, err
		}
		// TODO:...
		return nil, errHttp.(error)
	}
	var data ExpOutPut
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
