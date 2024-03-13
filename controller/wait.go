package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	c "zsl-demo/components"
)

var currentUrl = -1
var urls = [...]string{
	"/quest/utk",
	"/quest/sieci",
	"/quest/vim",
	"/quest/tssi",
}

func wait(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong id"))
		return
	}

	c.Base(true, id, c.Wait()).Render(r.Context(), w)
}

func waitPush(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// TODO remake using channels

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	tmpNext := currentUrl

	for {
		select {
		case <-ticker.C:
			if tmpNext != currentUrl {
				if currentUrl >= len(urls) {
					fmt.Fprint(w, "data: /scores\n\n")
				} else {
					fmt.Fprintf(w, "data: %s\n\n", urls[currentUrl])
				}
			}

			w.(http.Flusher).Flush()
		}
	}
}

func nextPush(w http.ResponseWriter, r *http.Request) {
	currentUrl++
	w.Write([]byte("next"))
}
