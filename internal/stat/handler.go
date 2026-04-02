package stat

import (
	"awesomeproject/configs"
	"awesomeproject/pkg/middleware"
	"awesomeproject/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByMonth = "month"
	GroupByDay   = "day"
)

type StatHandler struct {
	StatRepository *StatRepository
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func SetupStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.Handle("GET /stat", middleware.IsAuthed(handler.GetStat(), deps.Config))
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		layout := "2006-01-02"
		from, err := time.Parse(layout, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "invalid from param", http.StatusBadRequest)
			return
		}
		to, err := time.Parse(layout, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "invalid to param", http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")
		if by != GroupByMonth && by != GroupByDay {
			http.Error(w, "invalid by param", http.StatusBadRequest)
		}
		stats := handler.StatRepository.GetStats(by, from, to)
		res.Json(w, stats, http.StatusOK)
	}
}
