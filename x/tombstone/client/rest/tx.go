package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/daoleno/higan/x/tombstone/internal/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/record", storeName), SetRecordHandlerFn(cliCtx)).Methods("POST")
}

// SetRecordReq TX body
type SetRecordReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Name    string       `json:"name"`
	Born    time.Time    `json:"born"`
	Died    time.Time    `json:"died"`
	Memo    string       `json:"memo"`

	Recorder sdk.AccAddress `json:"recorder"`
}

// SetRecordHandlerFn set record
func SetRecordHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SetRecordReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSetRecord(req.Name, req.Born, req.Died, req.Memo, req.Recorder)
		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
