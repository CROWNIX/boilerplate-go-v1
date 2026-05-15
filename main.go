package main

import (
	"log"
	"sync"

	util_provider "gitea.qwertysystem.net/BETS/ts-utils/provider"
	"github.com/CROWNIX/boilerplate-go-v1/app"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	serviceProvider := util_provider.ServiceProviderImpl{}

	wg := sync.WaitGroup{}
	a := app.MakeApp()
	wg.Add(1)

	go func() {
		defer wg.Done()
		a.Run(&serviceProvider)
	}()

	wg.Wait()
}
