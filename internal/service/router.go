package service

import (
	"blob-base/internal/data/postgres"
	"blob-base/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBlobsQ(postgres.NewBlobsTable(s.cfg.DB())),
		),
	)
	r.Route("/blobs", func(r chi.Router) {
		r.Post("/", handlers.AddBlob)
		r.Get("/", handlers.GetBlobsList)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.GetBlob)
			r.Delete("/", handlers.DeleteBlob)
		})
	})

	r.Post("/accounts/", handlers.CreateAccount)

	return r
}
