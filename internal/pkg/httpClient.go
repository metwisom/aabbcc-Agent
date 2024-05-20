package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HttpClient struct {
	host  string
	token string
}

func (h *HttpClient) SetHost(host string) {
	h.host = host
}

func (h *HttpClient) SetToken(token string) {
	h.token = token
}

func (h *HttpClient) SendAspect(data []AspectFragment) {
	items := map[string]interface{}{"items": &data}
	h.send(items)
}

func (h *HttpClient) send(data map[string]interface{}) (err error) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(jsonData)
	requestURL := fmt.Sprintf(h.host + "/data")
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", h.token)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	var response *http.Response
	response, err = client.Do(req)
	if err != nil {
		return err
	}

	client.CloseIdleConnections()
	err = response.Body.Close()
	if err != nil {
		return err
	}
	return nil
	//bodyBytes, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
}
