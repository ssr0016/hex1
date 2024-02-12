package handler

import (
	"fmt"
	"net/http"

	"github.com/ssr0016/goHex/errors"
)

func handleError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case errors.AppError:
		w.WriteHeader(e.Code)
		fmt.Fprintln(w, e)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, e)
	}
}
