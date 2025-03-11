package main

import (
	"context"
	"net/http"
	"os"

	"github.com/dwethmar/berserk/sagemcomfast5359"
)

func main() {
	ctx := context.Background()
	c := &http.Client{}
	host := os.Getenv("HOST")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	routerClient := sagemcomfast5359.NewClient(c, host)
	if err := routerClient.Login(ctx, username, password); err != nil {
		panic(err)
	}
	d, err := routerClient.LanDevices(ctx)
	if err != nil {
		panic(err)
	}
	sagemcomfast5359.Table(os.Stdout, d)
}
