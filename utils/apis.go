package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mesh-uat-go-two/config"
	_type "mesh-uat-go-two/types"
	"net/http"
)

func MeshUatOneAppApiMeshOne() (*_type.ResponseDto, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, config.MeshUatGoOneEndpoint+"/api/v1/hello/one", nil)
	if req != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in connecting '%s': %s", config.MeshUatGoOneEndpoint, err.Error()))
	}

	defer resp.Body.Close()
	if resp.Body != nil {
		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error in getting response body: %s", err.Error()))
		}
		responseDto := _type.ResponseDto{}
		err = json.Unmarshal([]byte(jsonDataFromHttp), &responseDto) // here!
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error in parsing response body: %s", err.Error()))
		}

		return &responseDto, nil

	} else {
		return nil, errors.New(fmt.Sprintf("empty response body"))
	}
}

func MeshUatOneAppApiMeshTwo() (*_type.ResponseDto, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, config.MeshUatGoOneEndpoint+"/api/v1/hello/two", nil)
	if req != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error in connecting '%s': %s", config.MeshUatGoOneEndpoint, err.Error()))
	}

	defer resp.Body.Close()
	if resp.Body != nil {
		jsonDataFromHttp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error in getting response body: %s", err.Error()))
		}
		responseDto := _type.ResponseDto{}
		err = json.Unmarshal([]byte(jsonDataFromHttp), &responseDto) // here!
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error in parsing response body: %s", err.Error()))
		}

		return &responseDto, nil

	} else {
		return nil, errors.New(fmt.Sprintf("empty response body"))
	}
}
