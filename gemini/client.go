package gemini

import (
	"agent-orchestrator/agent"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// AgentResponse, Gemini'nin döndürdüğü agent ve parametre yapısı
type AgentResponse struct {
	Agent      string                 `json:"agent"`
	Parameters map[string]interface{} `json:"parameters"`
}

// Client, Gemini API istemcisi
type Client struct {
	apiKey string
}

// NewClient, yeni bir Gemini istemcisi oluşturur
func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

// Query, Gemini'ye istem gönderir ve agent yanıtını döndürür
func (c *Client) Query(prompt string) (*AgentResponse, error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=" + c.apiKey
	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": c.buildPrompt(prompt)},
				},
			},
		},
	}
	bodyBytes, _ := json.Marshal(body)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("Gemini API hatası: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Yanıt ayrıştırma hatası: %v", err)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("Geçersiz Gemini yanıtı")
	}

	// Debug için yanıtı yazdır
	rawResponse := result.Candidates[0].Content.Parts[0].Text
	fmt.Printf("Gemini yanıtı: %s\n", rawResponse)

	// Markdown formatını temizle
	jsonStr := rawResponse
	jsonStr = strings.TrimPrefix(jsonStr, "```json")
	jsonStr = strings.TrimPrefix(jsonStr, "```")
	jsonStr = strings.TrimSuffix(jsonStr, "```")
	jsonStr = strings.TrimSpace(jsonStr)

	// Fazladan karakterleri temizle
	jsonStr = strings.ReplaceAll(jsonStr, "\n", "")
	jsonStr = strings.ReplaceAll(jsonStr, "  ", " ")
	jsonStr = strings.TrimSpace(jsonStr)

	// En sondaki ``` karakterini kaldır
	if strings.HasSuffix(jsonStr, "```") {
		jsonStr = jsonStr[:len(jsonStr)-3]
	}

	var agentResp AgentResponse
	if err := json.Unmarshal([]byte(jsonStr), &agentResp); err != nil {
		return nil, fmt.Errorf("Agent yanıtı ayrıştırma hatası: %v\nJSON: %s", err, jsonStr)
	}
	return &agentResp, nil
}

// buildPrompt, Gemini'ye gönderilecek istemi oluşturur
func (c *Client) buildPrompt(input string) string {
	return fmt.Sprintf(`Kullanıcı: "%s"

Görev: Aşağıdaki agent'lardan hangisi uygun? Seç ve parametreleri belirt.
- weather-agent: Hava durumu bilgisi, "city" parametresi gerekir.
- translate-agent: Metin çevirisi, "text" ve "to" parametreleri gerekir.

ÖNEMLİ: Yanıtını SADECE aşağıdaki JSON formatında ver, başka hiçbir metin ekleme:
{
    "agent": "agent-ismi",
    "parameters": {
       "parametre1": "değer1",
    }
}`, input)
}

// ExecuteAgent, seçilen agent için Gemini'yi kullanarak işi yapar
func (c *Client) ExecuteAgent(agentt *agent.Agent, params map[string]interface{}) (string, error) {

	register := agent.NewRegistry()

	// Mock yanıt (gerçek Gemini API çağrısı için değiştir)
	switch agentt.Name {
	case "weather-agent":
		agent, ok := register.GetAgent(agentt.Name)
		if !ok {
			return "", fmt.Errorf("agent bulunamadı: %s", agentt.Name)
		}
		return agent.Execute(params)
	case "translate-agent":
		agent, ok := register.GetAgent(agentt.Name)
		if !ok {
			return "", fmt.Errorf("agent bulunamadı: %s", agentt.Name)
		}
		return agent.Execute(params)
	default:
		return "", fmt.Errorf("desteklenmeyen agent: %s", agentt.Name)
	}
}
