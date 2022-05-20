package eventClient_order

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"github.com/gojek/heimdall/v7/httpclient"
	"io"
	"net/http"
)

type EventClient struct {
	HttpClient *httpclient.Client
	BaseUrl    string
}

func (e EventClient) convertToReader(model interface{}) (io.Reader, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}

func (e EventClient) convertToBytes(model interface{}) ([]byte, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e EventClient) getSize(stream io.Reader) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Len()
}

func (e EventClient) OrderBooked(data core_order.Booked) (*core_order.Schedule, error) {
	var request BookedRequest
	request.FromEntity(data)
	payload, err := e.convertToBytes(request)
	if err != nil {
		return nil, err
	}
	res, err := e.HttpClient.Post(
		fmt.Sprintf("%s/event/schedule/booked", e.BaseUrl), bytes.NewReader(payload), http.Header{
			"Content-Type":   {"application/json"},
			"Content-Length": {string(len(payload))},
		},
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.Status != "201 Created" {
		return nil, errors.New("error while fetch")
	}

	var body common.Response[ScheduleResponse]
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Data.ToEntity(), nil
}

func (e EventClient) GetScheduleById(scheduleId string) (*core_order.Schedule, error) {
	res, err := e.HttpClient.Get(fmt.Sprintf("%s/event/schedule/%s", e.BaseUrl, scheduleId), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body common.Response[ScheduleResponse]
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Data.ToEntity(), nil
}

func (e EventClient) GetEventById(eventId string) (*core_order.Event, error) {
	res, err := e.HttpClient.Get(fmt.Sprintf("%s/event/%s", e.BaseUrl, eventId), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body common.Response[EventResponse]
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Data.ToEntity(), nil
}

func NewEventClient(module EventClient) core_order.IEventRepository {
	return &module
}
