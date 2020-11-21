package signup

import (
	"context"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cotrutatiberiu/project-go/dto"
	"github.com/cotrutatiberiu/project-go/utils/httpresponse"
)

func (controller *Controller) HandlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var payload dto.SignupPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(payload)

	errorTags := payload.Validate()
	if errorTags != nil {
		response := httpresponse.Reject{Status: 400, Location: "SIGNUP_CONTROLLER-ACCOUNT-VALIDATION-ERROR", Description: "Invalid data", Payload: errorTags, Error: nil}
		// http.Error(w, err.Error(), response.Status)
		payload, err := json.Marshal(response)
		if err != nil {

		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(payload)
	body, err := ioutil.ReadAll(resp.Body)

	}

	account := payload.ToModel()

	accountCreateError := controller.accountService.Create(context.Background(), account)
	if accountCreateError != nil {
		response := httpresponse.Reject{Status: 400, Location: "SIGNUP_CONTROLLER-ACCOUNT-CREATION-ERROR", Description: "Invalid data", Payload: accountCreateError, Error: nil}
		http.Error(w, err.Error(), response.Status)
		return
	}

	// response := httpresponse.Success{Status: 201, Location: "SIGNUP_CONTROLLER-ACCOUNT-CREATION", Description: "Success", Payload: errorTags}
	// w.Write([]byte)
}
