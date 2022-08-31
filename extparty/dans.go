package extparty

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"go-api/config"
	"go-api/models"
	"io/ioutil"
	"net/http"
	"time"
)

type IDans interface {
	GetJobList(string) (*[]models.Job, error)
	GetJobListByID(id string) (*models.Job, error)
}

type DansConstruct struct {
}

func NewDansConstruct() *DansConstruct {
	return &DansConstruct{}
}

func transportDansClient() (*http.Client, error) {
	timeout := time.Duration(2 * time.Minute)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: timeout}

	return client, nil
}

func dansHeader(address string, requestBody []byte) (*http.Request, error) {
	conf := config.GetConfig()
	dansHost := conf.GetString("extparty.dans_host")

	request, err := http.NewRequest("GET", dansHost+address, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

func (dansConstruct *DansConstruct) sendDansPayload(address string, param interface{}) ([]byte, error) {
	requestBody, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	client, err := transportDansClient()
	if err != nil {
		return nil, err
	}

	request, err := dansHeader(address, requestBody)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}

	// dump, err := httputil.DumpResponse(response, true)
	// if err != nil {
	// 	fmt.Printf(" Dump response failed with error %s\n", err)
	// 	return nil, err
	// }

	// fmt.Printf("DUMP Dans::%q", dump)

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dansConstruct *DansConstruct) GetJobList(queryParams string) (*[]models.Job, error) {
	var res []models.Job
	data, err := dansConstruct.sendDansPayload("/api/recruitment/positions.json?"+queryParams, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(data), &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (dansConstruct *DansConstruct) GetJobListByID(id string) (*models.Job, error) {
	var res models.Job
	data, err := dansConstruct.sendDansPayload("/api/recruitment/positions/"+id, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(data), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
