package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Hitesh3602/master_geography/internal/model"
	"github.com/Hitesh3602/master_geography/internal/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type getGeographiesResponse struct {
	Geographies []*model.Geography `json:"geographies,omitempty"`
	Geography   *model.Geography   `json:"geography,omitempty"`
}

type getGeographiesRequest struct {
	ID *int64 `json:"id,omitempty"`
}

//	type updateGeographyRequest struct {
//		ID   int64  `json:"id"`
//		Name string `json:"name"`
//	}
type updateGeographyRequest struct {
	ID       int64                  `json:"id"`
	Type     string                 `json:"type"`
	Name     string                 `json:"name"`
	Value    string                 `json:"value"`
	Metadata map[string]interface{} `json:"metadata"`
}

type deleteGeographyRequest struct {
	ID int64 `json:"id"`
}

//	type createGeographyRequest struct {
//		// Name string `json:"name"`
//		ID        int64     `json:"id"`
//		Name      string    `json:"name"`
//		Value     string    `json:"value"`
//		Metadata  string    `json:"metadata"`
//		CreatedAt time.Time `json:"created_at"`
//		UpdatedAt time.Time `json:"updated_at"`
//		Type      string    `json:"type"`
//	}
type createGeographyRequest struct {
	Type     string                 `json:"type"`
	Name     string                 `json:"name"`
	Value    string                 `json:"value"`
	Metadata map[string]interface{} `json:"metadata"`
}

func makeCreateGeographyEndpoint(svc service.GeographyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createGeographyRequest)
		metadata, err := json.Marshal(req.Metadata)
		if err != nil {
			return nil, err
		}

		geo := &model.Geography{
			Type:     req.Type,
			Name:     req.Name,
			Value:    req.Value,
			Metadata: metadata,
		}
		err = svc.CreateGeography(geo)
		if err != nil {
			return nil, err
		}
		return geo, nil
	}
}

func makeGetGeographiesEndpoint(svc service.GeographyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getGeographiesRequest)
		if req.ID != nil {
			geography, err := svc.GetGeographyByID(*req.ID)
			if err != nil {
				return nil, err
			}
			return getGeographiesResponse{Geography: geography}, nil
		} else {
			geographies, err := svc.GetGeographies()
			if err != nil {
				return nil, err
			}
			return getGeographiesResponse{Geographies: geographies}, nil
		}
	}
}

func makeUpdateGeographyEndpoint(svc service.GeographyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateGeographyRequest)
		metadata, err := json.Marshal(req.Metadata)
		if err != nil {
			return nil, err
		}

		geo := &model.Geography{
			ID:        req.ID,
			Type:      req.Type,
			Name:      req.Name,
			Value:     req.Value,
			Metadata:  metadata,
			UpdatedAt: time.Now(),
		}
		err = svc.UpdateGeography(geo)
		if err != nil {
			return nil, err
		}
		return geo, nil
	}
}

func makeDeleteGeographyEndpoint(svc service.GeographyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteGeographyRequest)
		err := svc.DeleteGeography(req.ID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}

func decodeCreateGeographyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// Print raw request body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Raw request body:", string(bodyBytes))

	// Reset the request body for further processing
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var req createGeographyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	// Debug print statements
	reqJSON, _ := json.Marshal(req)
	fmt.Println("Decoded request:", string(reqJSON))
	return req, nil
}

func decodeGetGeographiesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getGeographiesRequest
	vars := r.URL.Query()
	if id, ok := vars["id"]; ok && len(id) > 0 {
		id64, err := strconv.ParseInt(id[0], 10, 64)
		if err != nil {
			return nil, err
		}
		req.ID = &id64
	}
	return req, nil
}

func decodeUpdateGeographyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req updateGeographyRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	req.ID = id
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteGeographyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req deleteGeographyRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
