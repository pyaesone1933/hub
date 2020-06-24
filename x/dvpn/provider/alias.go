// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/provider/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/provider/querier
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/provider/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/provider/client/cli
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/provider/client/rest
package provider

import (
	"github.com/sentinel-official/hub/x/dvpn/provider/client/cli"
	"github.com/sentinel-official/hub/x/dvpn/provider/client/rest"
	"github.com/sentinel-official/hub/x/dvpn/provider/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider/querier"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

const (
	Codespace               = types.Codespace
	EventTypeSetProvider    = types.EventTypeSetProvider
	EventTypeUpdateProvider = types.EventTypeUpdateProvider
	AttributeKeyAddress     = types.AttributeKeyAddress
	ModuleName              = types.ModuleName
	QuerierRoute            = types.QuerierRoute
	QueryKeyProvider        = types.QueryKeyProvider
	QueryKeyProviders       = types.QueryKeyProviders
)

var (
	// functions aliases
	NewKeeper              = keeper.NewKeeper
	Querier                = querier.Querier
	RegisterCodec          = types.RegisterCodec
	ErrorMarshal           = types.ErrorMarshal
	ErrorUnmarshal         = types.ErrorUnmarshal
	ErrorUnknownMsgType    = types.ErrorUnknownMsgType
	ErrorUnknownQueryType  = types.ErrorUnknownQueryType
	ErrorInvalidField      = types.ErrorInvalidField
	ErrorDuplicateProvider = types.ErrorDuplicateProvider
	ErrorNoProviderFound   = types.ErrorNoProviderFound
	NewGenesisState        = types.NewGenesisState
	DefaultGenesisState    = types.DefaultGenesisState
	ProviderKey            = types.ProviderKey
	NewMsgRegisterProvider = types.NewMsgRegisterProvider
	NewMsgUpdateProvider   = types.NewMsgUpdateProvider
	NewQueryProviderParams = types.NewQueryProviderParams
	GetQueryCommands       = cli.GetQueryCommands
	GetTxCommands          = cli.GetTxCommands
	RegisterRoutes         = rest.RegisterRoutes

	// variable aliases
	ModuleCdc          = types.ModuleCdc
	RouterKey          = types.RouterKey
	StoreKey           = types.StoreKey
	ProviderKeyPrefix  = types.ProviderKeyPrefix
	QueryProviderPath  = types.QueryProviderPath
	QueryProvidersPath = types.QueryProvidersPath
)

type (
	Keeper              = keeper.Keeper
	GenesisState        = types.GenesisState
	MsgRegisterProvider = types.MsgRegisterProvider
	MsgUpdateProvider   = types.MsgUpdateProvider
	Provider            = types.Provider
	Providers           = types.Providers
	QueryProviderParams = types.QueryProviderParams
)
