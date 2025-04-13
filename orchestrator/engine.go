package orchestrator

import (
	"fmt"

	"agent-orchestrator/agent"
	"agent-orchestrator/gemini"
)

// Orchestrator, Gemini ve agent'ları koordine eder
type Orchestrator struct {
	geminiClient  *gemini.Client
	agentRegistry *agent.Registry
}

// New, yeni bir orchestrator oluşturur
func New(geminiClient *gemini.Client) *Orchestrator {
	return &Orchestrator{
		geminiClient:  geminiClient,
		agentRegistry: agent.NewRegistry(),
	}
}

// Process, kullanıcı istemini işler ve sonucu döndürür
func (o *Orchestrator) Process(input string) (string, error) {
	// Gemini'den agent ve parametreleri al
	resp, err := o.geminiClient.Query(input)
	if err != nil {
		return "", fmt.Errorf("Gemini sorgu hatası: %v", err)
	}

	// Agent'ı bul
	agent, ok := o.agentRegistry.GetAgent(resp.Agent)
	if !ok {
		return "", fmt.Errorf("Bilinmeyen agent: %s", resp.Agent)
	}

	// Gemini'yi kullanarak agent'ı çalıştır
	result, err := o.geminiClient.ExecuteAgent(&agent, resp.Parameters)
	if err != nil {
		return "", fmt.Errorf("%s hatası: %v", agent.Name, err)
	}

	return result, nil
}
