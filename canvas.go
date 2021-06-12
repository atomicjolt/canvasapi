package canvasapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type Canvas struct {
	AccessToken string
	CanvasURL   string
	UserAgent   string
}

func New(accessToken string, canvasURL string) Canvas {
	return Canvas{
		AccessToken: accessToken,
		CanvasURL:   canvasURL,
		UserAgent:   "AtomicJolt Go Agent v1.0",
	}
}

func (c *Canvas) SendRequest(canvasRequest CanvasRequest) (*http.Response, error) {
	err := canvasRequest.HasErrors()
	if err != nil {
		return nil, err
	}
	query, err := canvasRequest.GetQuery()
	if err != nil {
		return nil, err
	}

	request := http.Request{
		Method: canvasRequest.GetMethod(),
		Proto:  "HTTP/1.1",
		URL: &url.URL{
			Host:     c.CanvasURL,
			Scheme:   "https",
			Path:     path.Join("/api/v1", canvasRequest.GetURLPath()),
			RawQuery: query,
		},
		Header: http.Header{},
	}

	body, err := canvasRequest.GetBody()
	if err != nil {
		return nil, err
	}

	if body != "" {
		reqBody := ioutil.NopCloser(strings.NewReader(body))
		request.Body = reqBody
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	request.Header.Add("User-Agent", c.UserAgent)

	response, err := http.DefaultClient.Do(&request)
	if err != nil {
		return nil, err
	}

	var e error
	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return response, nil
	case http.StatusForbidden:
		response.Body.Close()
		return nil, ErrRateLimitExceeded
	case http.StatusUnprocessableEntity:
		return nil, fmt.Errorf("HTTP status: %v. %w", response.Status, response.Body.Close())
	case http.StatusNotFound, http.StatusUnauthorized:
		e = fmt.Errorf("HTTP status: %v. %w", response.Status, response.Body.Close())
	case http.StatusBadRequest, http.StatusInternalServerError:
		e = fmt.Errorf("HTTP status: %v. %w", response.Status, response.Body.Close())
	default:
		e = fmt.Errorf("HTTP status: %v. %w", response.Status, response.Body.Close())
	}
	return nil, fmt.Errorf("%v %v %w", json.NewDecoder(response.Body).Decode(&e), response.Body.Close(), e)
}
