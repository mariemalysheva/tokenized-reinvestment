package handler

import (
	"encoding/json"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/response"
	"github.com/mariemalysheva/tokenized-reinvestment/internal/handler/wrapper"
	"net/http"
)

func (i *Implementation) PostSetPrice(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req wrapper.PostSetPriceReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	txHash, err := i.adminSvc.PostSetPrice(ctx, req.Price)
	if err != nil {
		handleErrResponse(rw, err)
		return
	}

	response.OK(rw, wrapper.RepackTxHashResp(txHash))
}
