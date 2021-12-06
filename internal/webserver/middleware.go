package webserver

import (
	"net/http"
	"path"
)

func (h *handler) sessionValidateRedirect(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if err := h.validateSession(r); err != nil {
			newPath := path.Join(h.RootPath, "/login")
			redirectURL := createRedirectURL(newPath, r.URL.String())
			redirectPage(w, r, redirectURL)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
