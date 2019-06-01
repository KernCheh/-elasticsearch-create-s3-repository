# Configure S3 Repository on Elasticsearch

## Environment Variables

 Name | Required | Default | Description |
|------|----------|---------|-------------|
| `ELASTICSEARCH_URL` | n | `http://127.0.0.1:9200` | Elasticsearch endpoint to target  |
| `AWS_ACCESS_KEY_ID` | y | `-` | AWS Access Key used to communicate with S3 bucket  |
| `AWS_SECRET_ACCESS_KEY` | y | `-` | AWS Secret Key used to communicate with S3 bucket  |
| `APP_ENV` | n | `development` | Environment deployed against  |
| `BUCKET_NAME` | n | `sephora-k8s-logging` | Bucket Name to link to  |

## Usage

```
./elasticsearch-archiving-linux-amd64
```

Please remember to set the above environment variables

## Contributing

To install dependencies:

```
dep ensure
```

All environment variables are housed in `config.go`
