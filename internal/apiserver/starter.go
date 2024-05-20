package apiserver

import (
	"log"
	"net/http"

	"github.com/gola/internal/conf"
	"github.com/gola/internal/storage"
)

func Start(config *conf.Config) error {

	r, err := storage.New(config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	srv := NewServer(&r)
	http.ListenAndServe(config.Addr, srv)

	return nil
}
