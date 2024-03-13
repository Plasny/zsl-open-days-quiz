package controller

import (
	"log"
	"net/http"
	"strconv"
	c "zsl-demo/components"
	"zsl-demo/util"

	"log/slog"
)

func Init() {
	port := util.EnvOr("PORT", "8080")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			c.Base(false, 0, c.PageNotFound()).Render(r.Context(), w)
			return
		}

		t := c.Base(true, 0, c.Start())
		t.Render(r.Context(), w)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("GET /quest/utk", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong id"))
			return
		}
		c.Base(true, id, c.QuestUTK()).Render(r.Context(), w)
	})
	mux.HandleFunc("GET /quest/sieci", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong id"))
			return
		}
		c.Base(true, id, c.QuestSieci()).Render(r.Context(), w)
	})
	mux.HandleFunc("GET /quest/vim", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong id"))
			return
		}
		c.Base(true, id, c.QuestVim()).Render(r.Context(), w)
	})
	mux.HandleFunc("GET /quest/tssi", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("wrong id"))
			return
		}
		c.Base(true, id, c.QuestTSSI()).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /scores", scores)

	mux.HandleFunc("/wait", wait)
	mux.HandleFunc("GET /wait/push", waitPush)

	mux.HandleFunc("GET /api/nextQuest", nextPush)

	mux.HandleFunc("POST /api/times", addTime)
	mux.HandleFunc("POST /api/users", addUser)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	slog.Info("Server running on port: " + port)
	log.Fatal(srv.ListenAndServe())
}
