package client

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sigchat/sc-http/pkg/transport/errors"
	"github.com/sigchat/sc-http/pkg/transport/server/auth"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type HTTPClient struct {
	cli     *fasthttp.Client
	session *auth.Session
	config  *UsersService
}

func NewHTTPClient(config *UsersService) *HTTPClient {
	return &HTTPClient{
		cli: &fasthttp.Client{
			MaxIdleConnDuration: 10 * time.Second,
		},
		session: auth.NewSession(),
		config:  config,
	}
}

func (c *HTTPClient) GetUserByID(ctx context.Context, id int) (*dto.UserInfoDTO, error) {
	fullPath := fmt.Sprintf(`%s/api/v1/users/%d`, c.config.BaseUrl, id)

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.SetRequestURI(fullPath)
	req.Header.SetMethod(fasthttp.MethodGet)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	b, err := json.Marshal(id)
	if err != nil {
		return nil, fmt.Errorf("statusCode user by id: %v", err)
	}

	if err = c.session.SetToken(ctx, req); err != nil {
		return nil, fmt.Errorf("set token: %w", err)
	}

	req.SetBodyRaw(b)
	if err = c.cli.Do(req, resp); err != nil {
		return nil, errors.NewHttpError().
			WithCode(http.StatusServiceUnavailable).
			WithMessage(fmt.Sprintf("client get: %v", err))
	}

	if resp.StatusCode() > http.StatusBadRequest-1 {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode())
	}

	rBody := bufio.NewReader(bytes.NewBuffer(resp.Body()))

	var userInfo dto.UserInfoDTO
	if err = json.NewDecoder(rBody).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("json decode: %v", err)
	}
	return &userInfo, nil
}
