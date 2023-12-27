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
	ctx := context.WithValue(context.Background(), "TOKEN", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhdXRoLXNlcnZpY2UiLCJzdWIiOiJhdXRoIiwiYXVkIjpbIioiXSwiZXhwIjoxNzAyOTIwODE5LCJpYXQiOjE3MDI5MTcyMTksImp0aSI6IjQ1YTI0NGJkLWFhODMtNGQxYy04ZjBlLTJjNGM3MmJkZTc5NyIsImV4dHJhIjp7IlVzZXJJRCI6MX0sInJlZnJlc2hfZXhwaXJlc19hdCI6MH0.ayJPdsnxie8_vhbfSvBHFHjruDAYXIWVQkGtbkL8vUHXieJV4t10eNTIjc4Ac_9fBy3tG_TfpQ3D9lq3MTknmwlBQLfdqp2BTcDUI7DzjyyXWizq6t16ssycDZOKs3tB5cdjsGTIkFFFijcrb4rJUEKluWbCXyXQFIY3pn2G7y5R5ZTkjfqxcUkrxi9DJtOfQUzUujU7WUuQ9IBofoq43M83dytGUWUcyENjaVJtJZ1LwEhvpmcE4QsuWsx15KBGGauy4Rx1dEv6cB_LYKx50J8NBa8BzpIhxfPptCVYWrcDWcspiZdeZVfr2OIdgojT4jn0JsZYC1BYhuVQ7WifOw")
	info, err := cli.GetUserByID(ctx, 1)
	if err != nil {
		t.Fatal(fmt.Sprintf("failed to get user by id: %#v", err))
	}
	log.Println(info)
}
