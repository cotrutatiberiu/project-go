package signup

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cotrutatiberiu/project-go/dto"
)

func (controller *Controller) HandlePost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var payload dto.SignupPayload
	err = json.Unmarshal(b, &payload)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(payload)

	err = payload.Validate()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	account := payload.ToModel()

	controller.accountService.Create(context.Background(), account)

}
