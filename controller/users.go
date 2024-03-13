package controller

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	// c "zsl-demo/components"
	"zsl-demo/model"
)

func addUser(w http.ResponseWriter, r *http.Request) {
	var sqlName sql.NullString
	username := r.FormValue("username")

	if username == "" {
		sqlName.Valid = false
	} else {
		sqlName.Valid = true
		sqlName.String = username
	}

	q := model.New(model.DB)
	id, err := q.AddUserReturnId(r.Context(), sqlName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database error"))
		return
	}

	slog.Info(fmt.Sprintf("New user with id=%v", id))
	http.Redirect(w, r, fmt.Sprintf("/wait?id=%v", id), http.StatusTemporaryRedirect)
	// c.Base(true, id, c.Wait()).Render(r.Context(), w)
}
