package response

import (
	"encoding/json"
	"net/http"
)

type Info struct {
	Status  int         `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

func Writer(w http.ResponseWriter, data Info) {
	w.Header().Set("Content-Type", "application/json")

	buff, err := json.Marshal(data)
	if err != nil {
		data.Status = http.StatusInternalServerError
		data.Message = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":500,"message":"Internal Server Error"}`))
		return
	}

	w.WriteHeader(data.Status)
	w.Write(buff)
}
