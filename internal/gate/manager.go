package gate

import (
	"fmt"

	"wallet-service/internal/config"
)

type Manager struct {
	gates map[string]config.Gate
}

func New(cfg []config.Gate) *Manager {
	m := &Manager{
		gates: make(map[string]config.Gate),
	}

	for _, gate := range cfg {
		m.gates[gate.Name] = gate
	}

	return m
}

func (m *Manager) Mnemonic(
	name string,
) (string, error) {

	gate, ok := m.gates[name]
	if !ok {
		return "", fmt.Errorf(
			"gate not found: %s",
			name,
		)
	}

	return gate.Mnemonic, nil
}