package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HttpClient struct {
	host string
}

func (h *HttpClient) SetHost(host string) {
	h.host = host
}

func (h *HttpClient) SendAspect(data *AspectFragment) {
	h.send(data)
}

func (h *HttpClient) send(data any) (err error) {
	jsonData, err := json.MarshalIndent(&data, "", "\t")
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

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	_, err = client.Do(req)
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
