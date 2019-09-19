package gate

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Gate interface {
	Open()
}

type gate struct {
	router *httprouter.Router
}

func NewGate() (Gate, error) {
	g := &gate{
		router: httprouter.New(),
	}
	return g, nil
}

func (gate *gate) Open() {
	gate.router.POST("/mine", Miner)
	gate.router.OPTIONS("/mine", slient)
	log.Fatal(http.ListenAndServe(":8080", gate.router))
}