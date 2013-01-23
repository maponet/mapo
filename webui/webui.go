package webui

import (
    "net/http"
)

// questa funzione restituisce il contenuto della pagina / (root)
func Root(out http.ResponseWriter, in *http.Request) {

    http.ServeFile(out, in, "./webui/static/root.html")
}
