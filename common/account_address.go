package common

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type AccountAddress [32]byte

var (
	AccountAddressZero  AccountAddress
	AccountAddressOne   AccountAddress
	AccountAddressTwo   AccountAddress
	AccountAddressThree AccountAddress
	AccountAddressFour  AccountAddress
)

func init() {
	AccountAddressZero, _ = HexToAccountAddress("0x0")
	AccountAddressOne, _ = HexToAccountAddress("0x1")
	AccountAddressTwo, _ = HexToAccountAddress("0x2")
	AccountAddressThree, _ = HexToAccountAddress("0x3")
	AccountAddressFour, _ = HexToAccountAddress("0x4")
}

func (addr AccountAddress) PrefixZeroTrimmedHex() string {
	nonZeroIndex := 0
	for nonZeroIndex < 32 && addr[nonZeroIndex] == 0 {
		nonZeroIndex++
	}
	if nonZeroIndex == 32 {
		return "0x0"
	}

	h := hex.EncodeToString(addr[nonZeroIndex:])
	return "0x" + strings.TrimPrefix(h, "0")
}

func (addr AccountAddress) ToHex() string {
	return "0x" + hex.EncodeToString(addr[:])
}

func HexToAccountAddress(addr string) (AccountAddress, error) {
	addr = strings.TrimPrefix(addr, "0x")
	if len(addr)%2 == 1 {
		addr = "0" + addr
	}
	addrBytes, err := hex.DecodeString(addr)
	if err != nil {
		return [32]byte{}, err
	}

	length := len(addrBytes)
	if length > 32 {
		return [32]byte{}, fmt.Errorf("unexpected addr length: %d", length)
	}

	paddingBytes := make([]byte, 32-length)
	addrBytes = append(paddingBytes, addrBytes...)
	return *(*[32]byte)(addrBytes), nil
}
