package service

import (
	"fmt"

	"wallet-service/internal/config"
)

type GateManager struct {
	gates map[string]config.Gate
}

func NewGateManager(
	gates []config.Gate,
) *GateManager {

	items := make(
		map[string]config.Gate,
		len(gates),
	)

	for _, gate := range gates {
		items[gate.Name] = gate
	}

	return &GateManager{
		gates: items,
	}
}

func (m *GateManager) Get(
	name string,
) (config.Gate, error) {

	gate, ok := m.gates[name]
	if !ok {
		return config.Gate{},
			fmt.Errorf(
				"gate not found: %s",
				name,
			)
	}

	return gate, nil
}

func (m *GateManager) Mnemonic(
	name string,
) (string, error) {

	gate, err := m.Get(name)
	if err != nil {
		return "", err
	}

	return gate.Mnemonic, nil
}