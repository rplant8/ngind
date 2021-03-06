// +build !sputnikvm

package core

import (
	"math/big"

	"github.com/NginProject/ngind/core/state"
	"github.com/NginProject/ngind/core/types"
	evm "github.com/NginProject/ngind/core/vm"
)

const SputnikVMExists = false

var UseSputnikVM = "false"

func ApplyMultiVmTransaction(config *ChainConfig, bc *BlockChain, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, totalUsedGas *big.Int) (*types.Receipt, evm.Logs, *big.Int, error) {
	panic("not implemented")
}
