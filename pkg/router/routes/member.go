package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/youssef-182/production-server/pkg/models"
)

func MembersRouter(r chi.Router) {
	r.Route("/members", func(r chi.Router) {
		r.Get("/", models.MemberIndex)
		r.Post("/create", models.MemberStore)
		r.Route("/member/{memberID}", func(r chi.Router) {
			r.Use(MemberCtx)
			r.Get("/", models.MemberShow)
			r.Put("/", models.MemberUpdate)
			r.Delete("/", models.MemberDelete)
		})
		// r.Get("/", member.Index)
	})
}

func MemberCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		memberID := chi.URLParam(r, "memberID")

		ctx := context.WithValue(r.Context(), "memberID", memberID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
