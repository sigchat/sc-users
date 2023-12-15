package client

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestHTTPClient_GetUserByID(t *testing.T) {
	usersClientConf := &UsersService{
		Name:    "abobix",
		BaseUrl: "https://homaderaka.duckdns.org:8080",
	}
	cli := NewHTTPClient(usersClientConf)
	ctx := context.WithValue(context.Background(), "TOKEN", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXNlcnZpY2UiLCJzdWIiOiJhdXRoIiwiYXVkIjpbIioiXSwiZXhwIjoxNzAyNjU5NzgxLCJpYXQiOjE3MDI2NTYxODEsImp0aSI6ImU4YzBiZjZmLWU0N2EtNGY5MS04YjVhLTVmNzFmZDRjMjM5ZiIsImV4dHJhIjp7IlVzZXJJRCI6MX0sInJlZnJlc2hfZXhwaXJlc19hdCI6MH0.jf5vIRR8wUHcH9MOajVV4BJ-7P4kbJLZfUai-qsiTghInx-gDZn8CRStaBKky51bv455V0l_6VGOvlkk9CVrosi24k8NdkB6AwSmNO8tWGNArICXDO0QC24ysNYcrYg9T-RL_0YX3LDEX_5YJauEL4YiPZwF3Ll2CZrUlcMRpLjYTkHuTHge6FWRtPds9DnXEpT02UqgfvNC1AOX2dgZ6T_3FXWaqwYRe3N01Q_U9V3auMtA6D3lA3dDZCf2zoAWLgLmKsaXgbmL4vlKuixbWVkyY88AF54Z_6QhEw0ghqzYM5ZUKSc_ffhrGmY2pnkcnz5goUuWDNnFRwIF01TVJQ")
	info, err := cli.GetUserByID(ctx, 1)
	if err != nil {
		t.Fatal(fmt.Sprintf("failed to get user by id: %#v", err))
	}
	log.Println(info)
}
