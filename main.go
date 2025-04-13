package main

import (
	"bufio"
	"fmt"
	"os"

	"agent-orchestrator-gemini/gemini"
	"agent-orchestrator-gemini/orchestrator"
)

func main() {
	// Gemini API istemcisi (gerçek bir API anahtarı ile değiştir)
	geminiClient := gemini.NewClient("your-api-key")

	// Orchestrator oluştur
	orch := orchestrator.New(geminiClient)

	// Kullanıcıdan giriş al
	fmt.Println("İsteminizi girin (çıkmak için 'exit'):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}

		// Orchestrator ile sonucu al
		result, err := orch.Process(input)
		if err != nil {
			fmt.Printf("Hata: %v\n", err)
		} else {
			fmt.Printf("Sonuç: %s\n", result)
		}

		fmt.Println("\nYeni istem girin:")
	}
}
