package types

const (
	ModuleName = "ratelimit" // IBC at the end to avoid conflicts with the ibc prefix

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	PathKeyPrefix = KeyPrefix("path")
)