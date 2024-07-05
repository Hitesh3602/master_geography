package http

import (
	"net/http"

	"github.com/Hitesh3602/master_geography/internal/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(svc service.GeographyService) http.Handler {
	router := mux.NewRouter()

	router.Methods("POST").Path("/rms/master_geography").Handler(httptransport.NewServer(
		makeCreateGeographyEndpoint(svc),
		decodeCreateGeographyRequest,
		encodeResponse,
	))

	router.Methods("GET").Path("/rms/master_geography").Handler(httptransport.NewServer(
		makeGetGeographiesEndpoint(svc),
		decodeGetGeographiesRequest,
		encodeResponse,
	))

	router.Methods("PUT").Path("/rms/master_geography/{id}").Handler(httptransport.NewServer(
		makeUpdateGeographyEndpoint(svc),
		decodeUpdateGeographyRequest,
		encodeResponse,
	))

	router.Methods("DELETE").Path("/rms/master_geography/{id}").Handler(httptransport.NewServer(
		makeDeleteGeographyEndpoint(svc),
		decodeDeleteGeographyRequest,
		encodeResponse,
	))

	return router
}
