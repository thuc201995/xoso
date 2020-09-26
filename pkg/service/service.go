package service

import (
	"context"
	"fmt"
	"time"

	"github.com/thuc201995/xoso/pkg/config"
	"github.com/thuc201995/xoso/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

type XSMN struct {
	ID           string    `bson:"_id" json:"id"`
	Date         time.Time `bson:"date" json:"date"`
	Code         string    `bson:"code" json:"code"`
	ProvinceCode string    `bson:"province_code" json:"province_code"`
	Prize8       string    `bson:"prize_8" json:"prize_8"`
	Prize7       string    `bson:"prize_7" json:"prize_7"`
	Prize6       []string  `bson:"prize_6" json:"prize_6"`
	Prize5       string    `bson:"prize_5" json:"prize_5"`
	Prize4       []string  `bson:"prize_4" json:"prize_4"`
	Prize3       []string  `bson:"prize_3" json:"prize_3"`
	Prize2       string    `bson:"prize_2" json:"prize_2"`
	Prize1       string    `bson:"prize_1" json:"prize_1"`
	PrizeDB      string    `bson:"prize_db" json:"prize_db"`
}
type XSMT struct {
	ID           string    `bson:"_id" json:"id"`
	Date         time.Time `bson:"date" json:"date"`
	ProvinceCode string    `bson:"province_code" json:"province_code"`
	Prize8       string    `bson:"prize_8" json:"prize_8"`
	Prize7       string    `bson:"prize_7" json:"prize_7"`
	Prize6       []string  `bson:"prize_6" json:"prize_6"`
	Prize5       string    `bson:"prize_5" json:"prize_5"`
	Prize4       []string  `bson:"prize_4" json:"prize_4"`
	Prize3       []string  `bson:"prize_3" json:"prize_3"`
	Prize2       string    `bson:"prize_2" json:"prize_2"`
	Prize1       string    `bson:"prize_1" json:"prize_1"`
	PrizeDB      string    `bson:"prize_db" json:"prize_db"`
}

// XosoService describes the service.
type XosoService interface {
	// Add your methods here
	GetByDate(ctx context.Context, date string) (rs []XSMN, err error)
	GetByProvince(ctx context.Context, date string, province string) (rs XSMN, err error)
	GetXSMTByDate(ctx context.Context, date string) (rs []XSMT, err error)
	GetMTByProvince(ctx context.Context, date string, province string) (rs XSMT, err error)
}

type basicXosoService struct{}

func (b *basicXosoService) GetByDate(ctx context.Context, date string) (rs []XSMN, err error) {
	// TODO implement the business logic of GetByDate
	var t time.Time
	client, err := db.GetMongoClient()
	if err != nil {
		return rs, err
	}
	collection := client.Database("test").Collection("xsmn")
	t, err = time.Parse("02-01-2006", date)
	if err != nil {
		return rs, config.DateTimeError
	}
	data, err1 := collection.Find(ctx, bson.M{"date": t})
	if err1 != nil {
		return rs, err
	}
	if err = data.All(ctx, &rs); err != nil {
		return rs, err
	}
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

func (b *basicXosoService) GetByProvince(ctx context.Context, date string, province string) (rs XSMN, err error) {
	var t time.Time
	client, err := db.GetMongoClient()
	if err != nil {
		return rs, err
	}
	collection := client.Database("test").Collection("xsmn")
	t, err = time.Parse("02-01-2006", date)
	if err != nil {
		return rs, config.DateTimeError
	}
	err = collection.FindOne(ctx, bson.M{"date": t, "province_code": province}).Decode(&rs)
	if err != nil {
		return rs, err
	}

	return rs, err
}

func (b *basicXosoService) GetXSMTByDate(ctx context.Context, date string) (rs []XSMT, err error) {
	var t time.Time
	client, err := db.GetMongoClient()
	fmt.Println(5454545)
	if err != nil {
		return rs, err
	}
	collection := client.Database("test").Collection("xsmt")
	t, err = time.Parse("02-01-2006", date)
	if err != nil {
		return rs, config.DateTimeError
	}
	data, err1 := collection.Find(ctx, bson.M{"date": t})
	if err1 != nil {
		return rs, err
	}
	if err = data.All(ctx, &rs); err != nil {
		return rs, err
	}
	return rs, err
}

func (b *basicXosoService) GetMTByProvince(ctx context.Context, date string, province string) (rs XSMT, err error) {
	var t time.Time
	client, err := db.GetMongoClient()
	if err != nil {
		return rs, err
	}
	collection := client.Database("test").Collection("xsmt")
	t, err = time.Parse("02-01-2006", date)
	if err != nil {
		return rs, config.DateTimeError
	}
	err = collection.FindOne(ctx, bson.M{"date": t, "province_code": province}).Decode(&rs)
	if err != nil {
		return rs, err
	}

	return rs, err
}
