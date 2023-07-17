package client

import (
	"cliclient/pkg/logging"
	"net/http"
	"net/url"
)

type Client struct {
	logger logging.Logger
	//client - сам клиент
	client *http.Client
	//config - параметры клиента
	//config *config.Agent
	url *url.URL
}
