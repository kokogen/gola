package apiserver

import (
	"log"

	"github.com/gola/internal/conf"
	"github.com/gola/internal/storage"
)

func Start(config *conf.Config) error {

	r, err := storage.New(config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	srv := NewServer(&r, config.Addr)
	srv.Run()

	return nil
}
