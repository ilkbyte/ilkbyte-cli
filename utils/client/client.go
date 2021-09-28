package client

import ilkbyte "github.com/ilkbyte/api-go"

type Option struct {
	AccessKey string
	SecretKey string
}

func New(option *Option) *ilkbyte.Client {
	return ilkbyte.NewClient(option.AccessKey, option.SecretKey)
}
