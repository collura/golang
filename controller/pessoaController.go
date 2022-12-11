package controller

import (
	"github/collura/service"
	"io"
	"net/http"
	"strconv"
)

func init() {
	println("INIT")
}

func HandlerPessoa(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	val, _ := strconv.ParseInt(id, 0, 8)

	io.WriteString(w, service.Handler(val))
}
