package agent

import (
	"fmt"
)

// Agent, bir API'yi temsil eder
type Agent struct {
	Name    string
	Execute func(params map[string]interface{}) (string, error)
}

// Registry, tüm agent'ları tutar
type Registry struct {
	agents map[string]Agent
}

// NewRegistry, yeni bir agent kayıt defteri oluşturur
func NewRegistry() *Registry {
	r := &Registry{
		agents: make(map[string]Agent),
	}
	r.registerAgents()
	return r
}

// registerAgents, örnek agent'ları kaydeder
func (r *Registry) registerAgents() {
	r.agents["weather-agent"] = Agent{
		Name: "weather-agent",
		Execute: func(params map[string]interface{}) (string, error) {
			city, ok := params["city"].(string)
			if !ok {
				return "", fmt.Errorf("Geçersiz şehir parametresi")
			}
			// İlgili hava durumu sorgulama işlemlerini burada yapın
			return fmt.Sprintf("%s için hava durumu agenti çalıştırıldı.", city), nil
		},
	}

	r.agents["translate-agent"] = Agent{
		Name: "translate-agent",
		Execute: func(params map[string]interface{}) (string, error) {
			text, ok := params["text"].(string)
			if !ok {
				return "", fmt.Errorf("Geçersiz metin parametresi")
			}
			toLang, ok := params["to"].(string)
			if !ok {
				return "", fmt.Errorf("Geçersiz dil parametresi")
			}
			// İlgili çeviri işlemlerini burada yapın
			return fmt.Sprintf("Çeviri: %s -> %s agenti çalıştırıldı.", text, toLang), nil
		},
	}
}

// GetAgent, isme göre agent döndürür
func (r *Registry) GetAgent(name string) (Agent, bool) {
	agent, ok := r.agents[name]
	return agent, ok
}
