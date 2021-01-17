package commands

import (
	"testing"
)

func Test_mockSmartContract(t *testing.T) {
	contract := mockSmartContract(&Order{}, "81a12a32d3f9aa4cccd881dabe341fab50aa3f5d5afa91a049b4bec8827a0ccc")
	t.Log(contract)
}
