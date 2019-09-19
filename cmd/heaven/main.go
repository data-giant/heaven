package main

import (
	"heaven/gate"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main()  {
	err := godotenv.Load("./cmd/heaven/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	g, err := gate.NewGate()
	if err != nil{
		fmt.Errorf("Start gate error:", err)
		return
	}
	g.Open()
}


