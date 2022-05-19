package eventClient_order

import (
	"encoding/json"
	"fmt"
	"github.com/ac-kurniawan/loki-azure/pkg/common"
	core_order "github.com/ac-kurniawan/loki-azure/pkg/order/core"
	"github.com/gojek/heimdall/v7/httpclient"
)

type EventClient struct {
	HttpClient *httpclient.Client
	BaseUrl    string
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
