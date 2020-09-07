package service

import (
	"context"
	"fmt"
	"time"
)

type XSMN struct {
	Date         time.Time `bson:"date"`
	Code         string    `bson:"code"`
	ProvinceCode string    `bson:"province_code"`
	Prize8       string    `bson:"prize_8"`
	Prize7       string    `bson:"prize_7"`
	Prize6       []string  `bson:"prize_6"`
	Prize5       string    `bson:"prize_5"`
	Prize4       []string  `bson:"prize_4"`
	Prize3       []string  `bson:"prize_3"`
	Prize2       string    `bson:"prize_2"`
	Prize1       string    `bson:"prize_1"`
	PrizeDB      string    `bson:"prize_db"`
}

// XosoService describes the service.
type XosoService interface {
	// Add your methods here
	Gets(ctx context.Context, date string) (rs XSMN, err error)
	Bar(ctx context.Context, s string) (rs string, err error)
}

type basicXosoService struct{}

func (b *basicXosoService) Gets(ctx context.Context, date string) (rs XSMN, err error) {
	// TODO implement the business logic of Gets
	fmt.Println(date)
	return rs, err
}
func (b *basicXosoService) Bar(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Bar
	return rs, err
}

// NewBasicXosoService returns a naive, stateless implementation of XosoService.
func NewBasicXosoService() XosoService {
	return &basicXosoService{}
}

// New returns a XosoService with all of the expected middleware wired in.
func New(middleware []Middleware) XosoService {
	var svc XosoService = NewBasicXosoService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
