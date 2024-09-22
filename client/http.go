package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type (
	Request struct {
		URL                string            `json:"url" xml:"url"`
		Method             string            `json:"method" xml:"method"`
		Body               any               `json:"body" xml:"body"`
		Headers            map[string]string `json:"headers" xml:"headers"`
		TimeoutMiliseconds int               `json:"timeoutMiliseconds" xml:"timeoutMiliseconds"`
	}

	responseType    string
	Response[T any] struct {
		Success         bool    `json:"success" xml:"success"`
		StatusCode      int     `json:"statusCode" xml:"statusCode"`
		IsStatusCode2xx bool    `json:"isStatusCode2xx" xml:"isStatusCode2xx"`
		Message         string  `json:"message" xml:"message"`
		Data            T       `json:"data" xml:"data"`
		Request         Request `json:"request" xml:"request"`
	}
)

func HttpRequest[T any](request Request) Response[T] {
	var result T
	config := &http.Client{
		Timeout: time.Duration(request.TimeoutMiliseconds) * time.Millisecond,
	}

	jsonData, err := json.Marshal(request.Body)
	if err != nil {
		return response(reject, http.StatusInternalServerError, err.Error(), result, request)
	}

	req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if request.Headers != nil {
		for k, v := range request.Headers {
			req.Header.Set(k, v)
		}
	}

	if err != nil {
		return response(reject, http.StatusInternalServerError, err.Error(), result, request)
	}

	resp, err := config.Do(req)
	if err != nil {
		if resp != nil {
			message := fmt.Sprintf("error in calling api: %v", err)
			return response(reject, http.StatusInternalServerError, message, result, request)
		} else {
			message := fmt.Sprintf("error 2: %v", err)
			return response(reject, http.StatusInternalServerError, message, result, request)
		}
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		message := fmt.Sprintf("error in decoding response: %v", err)
		return response(reject, resp.StatusCode, message, result, request)
	}

	return response(success, resp.StatusCode, "success", result, request)
}

const (
	reject  = "reject"  // error
	success = "success" // success
)

func response[T any](respType responseType, statusCode int, message string, data T, request Request) Response[T] {
	is2xx := statusCode >= 200 && statusCode < 300
	if respType == success {
		return Response[T]{
			Success:         true,
			StatusCode:      statusCode,
			IsStatusCode2xx: is2xx,
			Message:         message,
			Data:            data,
			Request:         request,
		}
	}

	return Response[T]{
		Success:         false,
		StatusCode:      statusCode,
		IsStatusCode2xx: is2xx,
		Message:         message,
		Data:            data,
		Request:         request,
	}
}
