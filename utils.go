package gopyload

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func addMapToHeaders(req *http.Request, m map[string]string) {
	for k, v := range m {
		req.Header.Set(k, v)
	}
}

func endpointUrl(address string, endpoint endpoint) string {
	return fmt.Sprintf("%v%v", address, endpoint)
}


func StructureResBody[T any](data io.Reader) (*T, error) {
	var newData T
	bytes, err := io.ReadAll(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &newData)
	if err != nil {
		return nil, err
	}
	return &newData, nil
}

func DeStructure(data any) ([]byte, error) {
	newData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func SaveResToJson(data io.Reader, path string) error {
	b, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%v/%v.json", path, time.Now().UnixNano())
	err = os.WriteFile(fileName, b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
func SaveAnyToJson(data any, path, name string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%v/[%v] %v.json", path, name, time.Now().UnixNano())
	err = os.WriteFile(fileName, b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func GeneralPostEndpoint[T any](p *Pyload, e endpoint, v url.Values) (T, error) {
	res, err := p.client.PostForm(endpointUrl(p.Address, e), v)
	if err != nil {
		return *new(T), err
	}
	err, ok := HttpMapToStatusCode[res.StatusCode]
	if !ok {
		return *new(T), fmt.Errorf("status code %v not found", res.StatusCode)
	}
	if err != nil {
		return *new(T), err
	}
	val, err := StructureResBody[T](res.Body)
	if err != nil {
		return *new(T), err
	}
	return *val, nil
}
func GeneralGetEndpoint[T any](p *Pyload, e endpoint) (T, error) {
	res, err := p.client.Get(endpointUrl(p.Address, e))
	if err != nil {
		return *new(T), err
	}
	err, ok := HttpMapToStatusCode[res.StatusCode]
	if !ok {
		return *new(T), fmt.Errorf("status code %v not found", res.StatusCode)
	}
	if err != nil {
		return *new(T), err
	}
	val, err := StructureResBody[T](res.Body)
	if err != nil {
		return *new(T), err
	}
	return *val, nil
}
