package handlers

import (
	auth "github.com/fffnite/go-oneroster/internal/auth"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info(r)
		u := r.FormValue("clientid")
		p := r.FormValue("clientsecret")
		log.Infof("Attempting login; username: %v", u)
		t, err := auth.Login(u, p)
		if err != nil {
			// TODO: 401
			render.JSON(w, r, err)
			return
		}
		render.JSON(w, r, t)
	}
}
