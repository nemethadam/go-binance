package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetPositionMarginHistoryService get position margin history service
type GetPositionMarginHistoryService struct {
	c         *Client
	symbol    string
	_type     *int
	startTime *int64
	endTime   *int64
	limit     *int64
}

// Symbol set symbol
func (s *GetPositionMarginHistoryService) Symbol(symbol string) *GetPositionMarginHistoryService {
	s.symbol = symbol
	return s
}

// Type set type
func (s *GetPositionMarginHistoryService) Type(_type int) *GetPositionMarginHistoryService {
	s._type = &_type
	return s
}

// StartTime set startTime
func (s *GetPositionMarginHistoryService) StartTime(startTime int64) *GetPositionMarginHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetPositionMarginHistoryService) EndTime(endTime int64) *GetPositionMarginHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetPositionMarginHistoryService) Limit(limit int64) *GetPositionMarginHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetPositionMarginHistoryService) Do(ctx context.Context, opts ...RequestOption) (res []*PositionMarginHistory, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/positionMargin/history",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s._type != nil {
		r.setParam("type", *s._type)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*PositionMarginHistory, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// PositionMarginHistory define position margin history info
type PositionMarginHistory struct {
	Amount       string `json:"amount"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
	Type         int    `json:"type"`
	PositionSide string `json:"positionSide"`
}

// Do send request for OneWay account
func (s *GetPositionMarginHistoryService) DoOneWay(ctx context.Context, opts ...RequestOption) (res []*PositionMarginHistoryExtended, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/positionMargin/history",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s._type != nil {
		r.setParam("type", *s._type)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*PositionMarginHistoryExtended, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PositionMarginHistoryExtended struct {
	EntryPrice       string `json:"entryPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Notional         string `json:"notional"`
	IsolatedWallet   string `json:"isolatedWallet"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
	UpdateTime       string `json:"updateTime"`
}

// Do send request for Hedge account
func (s *GetPositionMarginHistoryService) DoHedge(ctx context.Context, opts ...RequestOption) (res [][]*PositionMarginHistoryExtended, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/positionMargin/history",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s._type != nil {
		r.setParam("type", *s._type)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([][]*PositionMarginHistoryExtended, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
