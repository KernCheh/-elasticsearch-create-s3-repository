package main

import (
	"context"

	"github.com/kerncheh/elasticsearch-create-s3-repository/config"
	"github.com/olivere/elastic"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()
var sugar = logger.Sugar()

func main() {
	defer logger.Sync()
	ctx := context.Background()

	client, err := elastic.NewClient(
		elastic.SetURL(config.GetInstance().ELASTICSEARCH_URL),
	)

	if err != nil {
		sugar.Panic(err)
	}

	sns := client.SnapshotCreateRepository(config.GetInstance().BUCKET_NAME)
	sns.Type("s3")
	sns.Settings(map[string]interface{}{
		"bucket": config.GetInstance().BUCKET_NAME,
		"region": "ap-southeast-1",
	})

	resp, err := sns.Do(ctx)
	if err != nil {
		sugar.Panic(err)
	}

	sugar.Info(resp)
}
