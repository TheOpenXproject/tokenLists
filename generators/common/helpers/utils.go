package helpers

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// GetKey returns the key of a token in a specific format to make it sortable
func GetKey(chainID uint64, address common.Address) string {
	chainIDStr := strconv.FormatUint(chainID, 10)
	chainIDStr = strings.Repeat("0", 18-len(chainIDStr)) + chainIDStr
	return chainIDStr + `_` + address.Hex()
}

// SafeString returns the provided variable or a fallback if it is empty
func SafeString(value string, fallback string) string {
	if value == `` {
		return fallback
	}
	return value
}

// SafeInt returns the provided variable or a fallback if it is empty
func SafeInt(value int, fallback int) int {
	if value == 0 {
		return fallback
	}
	return value
}

// ToAddress convert a string to a checksummed address
func ToAddress(value string) string {
	return common.HexToAddress(value).Hex()
}

func IncludesAddress(slice []string, value common.Address) bool {
	for _, item := range slice {
		if common.HexToAddress(item).Hex() == value.Hex() {
			return true
		}
	}
	return false
}

// Includes returns true if the provided T is in the provided slice
func Includes[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Contains returns true if value exists in arr
func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
