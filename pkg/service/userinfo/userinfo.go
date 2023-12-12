package userinfo

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type HTTPClient struct {
	*fasthttp.Client
}

func NewHTTPClient(serviceName string, clientName string) *HTTPClient {
	return &HTTPClient{
		Client: &fasthttp.Client{
			MaxIdleConnDuration: 10 * time.Second,
		},
	}
}

func (c *HTTPClient) GetUserByID(ctx context.Context, id int) (*dto.UserInfoDTO, error) {
	b, err := json.Marshal(id)
	if err != nil {
		return nil, fmt.Errorf("statusCode user by id: %v", err)
	}
	statusCode, body, err := c.Client.Get(b, fmt.Sprintf(`/api/v1/users/%d`, id))
	if err != nil {
		return nil, fmt.Errorf("client get: %v", err)
	}
	if statusCode > http.StatusBadRequest-1 {
		return nil, fmt.Errorf("wrong status code: %d", statusCode)
	}

	rBody := bufio.NewReader(bytes.NewBuffer(body))

	var userInfo dto.UserInfoDTO
	if err = json.NewDecoder(rBody).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("json decode: %v", err)
	}

	return &userInfo, nil
}
