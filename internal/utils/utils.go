package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, rawData any) {
	parsedJson, err := json.Marshal(rawData)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(parsedJson)
}
