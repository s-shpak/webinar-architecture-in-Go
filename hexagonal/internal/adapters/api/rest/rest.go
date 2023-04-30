package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"hexagonal/internal/core/domain"
)

type FoobarService interface {
	GetFoobar(req *domain.FoobarRequest) (*domain.FoobarResponse, error)
}

type API struct {
	srv *http.Server
}

func NewAPI(foobar FoobarService) *API {
	h := &handler{
		foobar: foobar,
	}
	router := httprouter.New()
	router.GET("/:n", h.getFoobarHandler)
	return &API{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}

func (a *API) Run() error {
	return a.srv.ListenAndServe()
}

type handler struct {
	foobar FoobarService
}

func (a *handler) getFoobarHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	nParam := ps.ByName("n")
	n, err := strconv.ParseInt(nParam, 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := a.foobar.GetFoobar(&domain.FoobarRequest{
		N: int(n),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	respBody, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(respBody); err != nil {
		return
	}
}
