package controller

import (
	"net/http"
	"strconv"
	c "zsl-demo/components"
	"zsl-demo/model"
)

func scores(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		id = 0
	}

	q := model.New(model.DB)
	arr, err := q.GetAllTotalTimes(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong with database"))
		return
	}

	c.Base(true, id, c.Scores(arr, id)).Render(r.Context(), w)
}
