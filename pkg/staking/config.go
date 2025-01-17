package staking

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func GetRPCURL() (string, error) {
	url := os.Getenv("RPC_URL")
	if url == "" {
		return "", errors.New("RPC_URL environment variable is not set")
	}
	return url, nil
}

func LoadContractAddresses() (*ContractAddresses, error) {
	path := filepath.Join("contracts", "node_modules", "@masa-finance", "masa-contracts-oracle", "addresses.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var addresses ContractAddresses
	err = json.Unmarshal(data, &addresses)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}
