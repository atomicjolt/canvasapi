package canvasapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
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

	body, err := canvasRequest.GetBody()
	if err != nil {
		return nil, err
	}

	canvasUrl := &url.URL{
		Host:     c.CanvasURL,
		Scheme:   "https",
		Path:     path.Join("/api/v1", canvasRequest.GetURLPath()),
		RawQuery: query,
	}

	return c.Send(canvasUrl, canvasRequest.GetMethod(), &body)
}

func (c *Canvas) Send(canvasUrl *url.URL, method string, body *url.Values) (*http.Response, error) {

	request := http.Request{
		Method: method,
		Proto:  "HTTP/1.1",
		URL:    canvasUrl,
		Host:   canvasUrl.Host,
		Header: http.Header{},
	}

	if body != nil {
		encodedBody := body.Encode()
		reqBody := bytes.NewBuffer([]byte(encodedBody))
		request.Body = ioutil.NopCloser(reqBody)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Add("Content-Length", strconv.Itoa(len(encodedBody)))
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
