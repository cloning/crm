package api

import (
	"../../../src/app/api"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	PORT = 4001
)

type apifn func(*api.Api)

func Get(path string) ([]byte, *http.Response, error) {
	apiUrl := fmt.Sprintf("http://localhost:%d", PORT)
	resp, err := http.Get(apiUrl + path)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, resp, err
}

func WithApi(fn apifn) {
	var wg sync.WaitGroup
	api := api.NewApi(PORT, wg)
	defer func() {
		api.Stop()
		wg.Wait()
	}()
	go api.Run()
	fn(api)
}
