package tombstone

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/daoleno/higan/x/tombstone/internal/types"
)

// NewHandler creates an sdk.Handler for all the tombstone type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgRecord:
			return handleMsgRecord(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgRecord set record to note
func handleMsgRecord(ctx sdk.Context, k Keeper, msg MsgRecord) (*sdk.Result, error) {
	k.SetNote(ctx, msg)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Recorder.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
