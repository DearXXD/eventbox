package impl

import (
	"context"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/eventbox/conf"
	"github.com/infraboard/eventbox/pkg"
	"github.com/infraboard/eventbox/pkg/event"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger

	event.UnimplementedServiceServer
}

func (s *service) Config() error {
	db := conf.C().Mongo.GetDB()
	col := db.Collection("event")
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}
	_, err := col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}
	s.col = col
	s.log = zap.L().Named("Event")
	return nil
}

// HttpEntry todo
func (s *service) HTTPEntry() *http.EntrySet {
	return event.HttpEntry()
}

func init() {
	pkg.RegistryService("event", Service)
}
