# Configure S3 Repository on Elasticsearch

## Environment Variables

 Name | Required | Default | Description |
|------|----------|---------|-------------|
| `ELASTICSEARCH_URL` | n | `http://127.0.0.1:9200` | Elasticsearch endpoint to target  |
| `APP_ENV` | n | `development` | Environment deployed against  |
| `BUCKET_NAME` | n | `sephora-k8s-logging` | Bucket Name to link to  |

## Usage

```
./elasticsearch-archiving-linux-amd64
```

Please remember to set the above environment variables

Sourcing in AWS keys as environment variables are no longer supported.
Please refer to https://www.elastic.co/guide/en/elasticsearch/plugins/current/repository-s3-client.html instead

## Contributing

To install dependencies:

```
dep ensure
```

All environment variables are housed in `config.go`
