package types

const (
	// ModuleName is the name of the module
	ModuleName = "marketplace"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey string = ModuleName
)

var (
	PrefixAuction = []byte{0x01}
)
