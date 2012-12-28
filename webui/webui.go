package webui

import (
    "net/http"
)

func Root(out http.ResponseWriter, in *http.Request) {

    http.ServeFile(out, in, "./webui/static/root.html")
}
