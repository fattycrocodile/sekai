package types

import (
	functionmeta "github.com/KiraCore/sekai/function_meta"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec register codec and metadata
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpsertTokenAlias{}, "kiraHub/MsgUpsertTokenAlias", nil)
	cdc.RegisterConcrete(&MsgUpsertTokenRate{}, "kiraHub/MsgUpsertTokenRate", nil)

	functionmeta.AddNewFunction((&MsgUpsertTokenAlias{}).Type(), `{
		"function_id": 0,
		"description": "MsgUpsertTokenAlias represents a message to register token alias.",
		"parameters": {
			"expiration": {
				"type":        "uint32",
				"description": "expiration time (uint32) - seconds since submission."
			},
			"enactment": {
				"type":        "uint32",
				"description": "enactment time (uint32) - seconds since expiration."
			},
			"allowed_vote_types": {
				"type":        "array<VoteType>",
				"description": "Allowed vote types: yes, no, veto, abstain."
			},
			"symbol": {
				"type":        "string",
				"description": "Ticker (eg. ATOM, KEX, BTC)"
			},
			"name": {
				"type":        "string",
				"description": "Token Name (e.g. Cosmos, Kira, Bitcoin)"
			},
			"icon": {
				"type":        "string",
				"description": "Graphical Symbol (url link to graphics)"
			},
			"decimals": {
				"type":        "int",
				"description": "Integer number of max decimals"
			},
			"denoms": {
				"type":        "array<string>",
				"description": "An array of token denoms to be aliased"
			},
			"status": {
				"type":        "enum",
				"description": "token alias voting status: undefined, active, rejected, passed, enacted"
			},
			"proposer": {
				"type":        "address",
				"description": "proposer of this message"
			}
		}
	}`)

	functionmeta.AddNewFunction((&MsgUpsertTokenRate{}).Type(), `{
		"function_id": 0,
		"description": "MsgUpsertTokenRate represents a message to register token rate.",
		"parameters": {
			"denom": {
				"type":        "string",
				"description": "denomination target."
			},
			"rate": {
				"type":        "float",
				"description": "Exchange rate in terms of KEX token. e.g. 0.1, 10.5"
			},
			"fee_payments": {
				"type":        "bool",
				"description": "defining if it is enabled or disabled as fee payment method."
			},
			"proposer": {
				"type":        "address",
				"description": "proposer of this message"
			}
		}
	}`)
}

// RegisterInterfaces register Msg and structs
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpsertTokenRate{},
		&MsgUpsertTokenAlias{},
	)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
