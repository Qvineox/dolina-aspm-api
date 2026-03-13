package db

import (
	"log/slog"
	"time"

	"github.com/orandin/slog-gorm"
	"gitlab.domsnail.ru/dolina/dolina-aspm-api/pkg/cfg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ref: https://gorm.io/docs/gorm_config.html

func NewOrmClient(config cfg.DatabaseConfig) (orm *gorm.DB, err error) {
	var opts = []slogGorm.Option{
		slogGorm.WithHandler(slog.Default().Handler()),
		slogGorm.WithSlowThreshold(time.Second * 5),
	}

	if config.TraceAllMessages {
		opts = append(opts, slogGorm.WithTraceAll())
	} else {
		opts = append(opts, slogGorm.WithIgnoreTrace())
	}

	connCfg := gorm.Config{
		SkipDefaultTransaction:                   false,
		DefaultTransactionTimeout:                time.Duration(300) * time.Second,
		FullSaveAssociations:                     false,
		Logger:                                   slogGorm.New(opts...),
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          100,
		TranslateError:                           true,
		PropagateUnscoped:                        true,
	}

	return gorm.Open(postgres.Open(config.Postgres()), &connCfg)
}
