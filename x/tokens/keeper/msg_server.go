package keeper

import (
	"context"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	customgovtypes "github.com/KiraCore/sekai/x/gov/types"
	"github.com/KiraCore/sekai/x/tokens/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	keeper Keeper
	cgk    types.CustomGovKeeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, cgk types.CustomGovKeeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
		cgk:    cgk,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ProposalUpsertTokenRates(goCtx context.Context, msg *types.MsgProposalUpsertTokenRates) (*types.MsgProposalUpsertTokenRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := k.cgk.CheckIfAllowedPermission(ctx, msg.Proposer, customgovtypes.PermCreateUpsertTokenRateProposal)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermCreateUpsertTokenRateProposal.String())
	}

	proposalID, err := k.CreateAndSaveProposalWithContent(ctx,
		msg.Description,
		types.NewProposalUpsertTokenRates(
			msg.Denom,
			msg.Rate,
			msg.FeePayments,
		))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			customgovtypes.EventTypeSubmitProposal,
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalId, fmt.Sprintf("%d", proposalID)),
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalType, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyRate, msg.Rate.String()),
			sdk.NewAttribute(types.AttributeKeyFeePayments, fmt.Sprintf("%t", msg.FeePayments)),
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer.String()),
		),
	)

	return &types.MsgProposalUpsertTokenRatesResponse{
		ProposalID: proposalID,
	}, err
}

func (k msgServer) ProposalUpsertTokenAlias(goCtx context.Context, msg *types.MsgProposalUpsertTokenAlias) (*types.MsgProposalUpsertTokenAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := k.cgk.CheckIfAllowedPermission(ctx, msg.Proposer, customgovtypes.PermCreateUpsertTokenAliasProposal)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermCreateUpsertTokenAliasProposal.String())
	}

	proposalID, err := k.CreateAndSaveProposalWithContent(
		ctx,
		msg.Description,
		types.NewProposalUpsertTokenAlias(
			msg.Symbol,
			msg.Name,
			msg.Icon,
			msg.Decimals,
			msg.Denoms,
		))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			customgovtypes.EventTypeSubmitProposal,
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalId, fmt.Sprintf("%d", proposalID)),
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalType, msg.Type()),
			sdk.NewAttribute(types.AttributeKeySymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyIcon, msg.Icon),
			sdk.NewAttribute(types.AttributeKeyDecimals, fmt.Sprintf("%d", msg.Decimals)),
			sdk.NewAttribute(types.AttributeKeyDenoms, strings.Join(msg.Denoms, ",")),
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer.String()),
		),
	)
	return &types.MsgProposalUpsertTokenAliasResponse{
		ProposalID: proposalID,
	}, err
}

func (k msgServer) ProposalTokensWhiteBlackChange(goCtx context.Context, msg *types.MsgProposalTokensWhiteBlackChange) (*types.MsgProposalTokensWhiteBlackChangeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := k.cgk.CheckIfAllowedPermission(ctx, msg.Proposer, customgovtypes.PermCreateTokensWhiteBlackChangeProposal)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermCreateTokensWhiteBlackChangeProposal.String())
	}

	proposalID, err := k.CreateAndSaveProposalWithContent(
		ctx,
		msg.Description,
		types.NewProposalTokensWhiteBlackChange(
			msg.IsBlacklist,
			msg.IsAdd,
			msg.Tokens,
		))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			customgovtypes.EventTypeSubmitProposal,
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalId, fmt.Sprintf("%d", proposalID)),
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalType, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer.String()),
			sdk.NewAttribute(types.AttributeKeyIsBlacklist, fmt.Sprintf("%t", msg.IsBlacklist)),
			sdk.NewAttribute(types.AttributeKeyIsAdd, fmt.Sprintf("%t", msg.IsAdd)),
			sdk.NewAttribute(types.AttributeKeyTokens, strings.Join(msg.Tokens, ",")),
			sdk.NewAttribute(types.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer.String()),
		),
	)
	return &types.MsgProposalTokensWhiteBlackChangeResponse{
		ProposalID: proposalID,
	}, err
}

func (k msgServer) UpsertTokenAlias(
	goCtx context.Context,
	msg *types.MsgUpsertTokenAlias,
) (*types.MsgUpsertTokenAliasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := k.cgk.CheckIfAllowedPermission(ctx, msg.Proposer, customgovtypes.PermUpsertTokenAlias)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermUpsertTokenAlias.String())
	}

	err := k.keeper.UpsertTokenAlias(ctx, *types.NewTokenAlias(
		msg.Symbol,
		msg.Name,
		msg.Icon,
		msg.Decimals,
		msg.Denoms,
	))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpsertTokenAlias,
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer.String()),
			sdk.NewAttribute(types.AttributeKeySymbol, msg.Symbol),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyIcon, msg.Icon),
			sdk.NewAttribute(types.AttributeKeyDecimals, fmt.Sprintf("%d", msg.Decimals)),
			sdk.NewAttribute(types.AttributeKeyDenoms, strings.Join(msg.Denoms, ",")),
		),
	)
	return &types.MsgUpsertTokenAliasResponse{}, err
}

func (k msgServer) UpsertTokenRate(goCtx context.Context, msg *types.MsgUpsertTokenRate) (*types.MsgUpsertTokenRateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.ValidateBasic()
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	isAllowed := k.cgk.CheckIfAllowedPermission(ctx, msg.Proposer, customgovtypes.PermUpsertTokenRate)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermUpsertTokenRate.String())
	}

	err = k.keeper.UpsertTokenRate(ctx, *types.NewTokenRate(
		msg.Denom,
		msg.Rate,
		msg.FeePayments,
	))

	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpsertTokenRate,
			sdk.NewAttribute(types.AttributeKeyProposer, msg.Proposer.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyRate, msg.Rate.String()),
			sdk.NewAttribute(types.AttributeKeyFeePayments, fmt.Sprintf("%t", msg.FeePayments)),
		),
	)

	return &types.MsgUpsertTokenRateResponse{}, nil
}

func (k msgServer) CreateAndSaveProposalWithContent(ctx sdk.Context, description string, content customgovtypes.Content) (uint64, error) {
	blockTime := ctx.BlockTime()
	proposalID := k.cgk.GetNextProposalIDAndIncrement(ctx)
	properties := k.cgk.GetNetworkProperties(ctx)

	proposal, err := customgovtypes.NewProposal(
		proposalID,
		content,
		blockTime,
		blockTime.Add(time.Second*time.Duration(properties.ProposalEndTime)),
		blockTime.Add(time.Second*time.Duration(properties.ProposalEndTime)+
			time.Second*time.Duration(properties.ProposalEnactmentTime),
		),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
		description,
	)

	if err != nil {
		return proposalID, err
	}

	k.cgk.SaveProposal(ctx, proposal)
	k.cgk.AddToActiveProposals(ctx, proposal)

	return proposalID, nil
}
