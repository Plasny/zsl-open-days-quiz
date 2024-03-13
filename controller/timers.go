package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	// c "zsl-demo/components"
	"zsl-demo/model"
)

type timeJSON struct {
	Id   int64  `json:"id"`
	Ms   int64  `json:"ms"`
	Task string `json:"task"`
}

func addTime(w http.ResponseWriter, r *http.Request) {
	var data timeJSON

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect JSON data"))
		return
	}

	q := model.New(model.DB)
	err = q.AddTime(r.Context(), model.AddTimeParams{
		UserID:   data.Id,
		UserID_2: data.Id,
		Time:     data.Ms,
		Type:     data.Task,
		Type_2:   data.Task,
	})
	if err != nil {
		slog.Info("error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database error"))
		return
	}

	slog.Info(fmt.Sprintf("New time id=%d ms=%d type=%s", data.Id, data.Ms, data.Task))
	w.Write([]byte("ok"))
	// c.Base(true, data.Id, c.Wait()).Render(r.Context(), w)
}
