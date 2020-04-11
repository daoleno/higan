package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/daoleno/higan/x/tombstone/internal/types"
)

// NewQuerier creates a new querier for tombstone clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryRecord:
			return queryRecord(ctx, k, path[1])
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown tombstone query endpoint")
		}
	}
}

func queryRecord(ctx sdk.Context, k Keeper, recorder string) ([]byte, error) {
	recorderAcc, err := sdk.AccAddressFromBech32(recorder)
	if err != nil {
		return nil, err
	}
	note, err := k.GetNote(ctx, recorderAcc)
	if err != nil {
		return nil, err
	}

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, note)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
