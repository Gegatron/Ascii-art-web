package asciiartweb

import "net/http"

// HandlerCss serves CSS files and handles access restrictions.
func HandlerCss(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/css/" || r.URL.Path == "/css" {
		ErrorsHandler(w, " Acess Forbidden", http.StatusForbidden)
		return
	}
	if r.URL.Path == "/css/assets/" || r.URL.Path == "/css/assets" {
		ErrorsHandler(w, " Acess Forbidden", http.StatusForbidden)
		return
	}
	if r.URL.Path != "/css/style.css" && r.URL.Path != "/css/assets/background.gif" && r.URL.Path != "/css/result.css" && r.URL.Path != "/css/errors.css" {
		ErrorsHandler(w, "Page not found", http.StatusNotFound)
		return
	}

	HandlerCss := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
	HandlerCss.ServeHTTP(w, r)
}
