package genesis

import (
	"encoding/json"
	"os"
	"time"
)

type Genesis struct {
	Date          time.Time         `json:"date"`
	ChainId       uint16            `json:"chain_id"`
	TransPerBlock uint16            `json:"trans_per_block"`
	Difficulty    uint16            `json:"difficulty"`
	MiningReward  uint64            `json:"mining_reward"`
	GasPrice      uint64            `json:"gas_price"`
	Balances      map[string]uint64 `json:"balances"`
}

// =========================================================================================================================

func Load() (Genesis, error) {
	// dir,_:= os.Getwd()
	// fmt.Println("Load() called" , dir)
	path := "zblock/genesis.json"
	content, err := os.ReadFile(path)
	if err != nil {
		return Genesis{}, err
	}

	var genesis Genesis
	err = json.Unmarshal(content, &genesis)
	if err != nil {
		return Genesis{}, err
	}

	return Genesis{}, nil
}
