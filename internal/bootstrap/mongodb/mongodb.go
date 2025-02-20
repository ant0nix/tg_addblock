package mongodb

import (
	"context"
	"github.com/ant0nix/tg_addblock/internal/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/v2/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/connstring"
	"time"
)

type DB struct {
	ctx      context.Context
	client   *mongo.Client
	database *mongo.Database
}

func (db *DB) Close() error {
	return db.client.Disconnect(db.ctx)
}

func (db *DB) New(ctx context.Context, cfg *config.Config) (*mongo.Database, error) {
	connString, err := connstring.ParseAndValidate(cfg.MongoDBURI)
	if err != nil {
		return nil, errors.Wrap(err, "parsing mongodb connection string")
	}
	if connString.Database == "" {
		return nil, errors.New("mongodb connection string is empty")
	}
	cmdMonitor := &event.CommandMonitor{}

	clientOpts := options.Client().
		SetTimeout(cfg.Timeout).
		SetHeartbeatInterval(cfg.PingInterval).
		ApplyURI(cfg.MongoDBURI).
		SetMonitor(cmdMonitor)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, errors.Wrap(err, "can't connecting to mongodb")
	}

	go func(ctx context.Context) {
		t := time.NewTicker(cfg.PingInterval)
		for range t.C {
			select {
			case <-ctx.Done():
				logrus.Info("shutting down mongodb")
				return
			default:
				if err := client.Ping(ctx, nil); err != nil {
					logrus.WithError(err).Error("ping mongodb")
				}
			}
		}
	}(ctx)

	db := &DB{
		ctx:      ctx,
		client:   client,
		database: client.Database(connString.Database),
	}
	//todo add closer
	return db.database, nil
}
