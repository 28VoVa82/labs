package exchange

import (
	"encoding/json"
	"log"
)

func ParseRequest(data []byte) (ClientRequest, bool) {
	if data == nil || len(data) == 0 {
		log.Printf("Request is empty.")
		return ClientRequest{}, false
	}

	var req ClientRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		log.Printf("Request parsing failed. Error: %s", err)
		return ClientRequest{}, false
	}

	return req, true
}

func SerializeRequest(request ClientRequest) []byte {
	data, err := json.Marshal(request)
	if err != nil {
		log.Printf("Request serialization failed. Error: %s", err)
		return nil
	}
	return data
}

func ParseRespond(data []byte) (ServerRespond, bool) {
	if data == nil || len(data) == 0 {
		log.Printf("Respond is empty.")
		return ServerRespond{}, false
	}

	var resp ServerRespond
	err := json.Unmarshal(data, &resp)
	if err != nil {
		log.Printf("Respond parsing failed. Error: %s", err)
		return ServerRespond{}, false
	}

	return resp, true
}

func SerializeRespond(respond ServerRespond) []byte {
	data, err := json.Marshal(respond)
	if err != nil {
		log.Printf("Respond serialization failed. Error: %s", err)
		return nil
	}
	return data
}