package go_epay

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-epay/utils"
)

type Client struct {
	Params    *EPayInitParams
	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *EPayInitParams) *Client {
	if params.BaseURL == "" {
		params.BaseURL = DefaultBaseURL
	}
	if params.DepositPath == "" {
		params.DepositPath = DepositPath
	}
	if params.DepositViaPath == "" {
		params.DepositViaPath = DepositViaPath
	}
	if params.WithdrawPath == "" {
		params.WithdrawPath = WithdrawPath
	}
	return &Client{
		Params:    params,
		ryClient:  resty.New(),
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debug bool) {
	cli.debugMode = debug
}
