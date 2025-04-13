# Agent Orchestrator Gemini

Bu proje, Google Gemini API'sini kullanarak çeşitli agent'ları (ajanları) yöneten ve koordine eden bir sistemdir.

## Özellikler

- Google Gemini API entegrasyonu
- Çoklu agent desteği
- Dinamik agent seçimi ve parametre yönetimi
- Genişletilebilir agent mimarisi

## Mevcut Agent'lar

1. **Hava Durumu Agent'ı (weather-agent)**
   - Şehir bazlı hava durumu bilgisi sağlar
   - Gerekli parametre: `city`

2. **Çeviri Agent'ı (translate-agent)**
   - Metin çevirisi yapar
   - Gerekli parametreler: `text`, `to`

## Kurulum

1. Projeyi klonlayın:
```bash
git clone https://github.com/kullanici/agent-orchestrator-gemini.git
cd agent-orchestrator-gemini
```

2. Bağımlılıkları yükleyin:
```bash
go mod download
```

3. Google Gemini API anahtarınızı ayarlayın:
   - `main.go` dosyasında `your-api-key` yerine kendi API anahtarınızı yazın

## Kullanım

1. Programı çalıştırın:
```bash
go run main.go
```

2. İstemlerinizi Türkçe olarak girin. Örnekler:
   - "İstanbul'un hava durumu nasıl?"
   - "Merhaba kelimesini İspanyolca'ya çevir"

3. Programdan çıkmak için 'exit' yazın.

## Geliştirme

Yeni bir agent eklemek için:

1. `agent/registry.go` dosyasında `registerAgents` fonksiyonuna yeni agent'ı ekleyin
2. `gemini/client.go` dosyasında `buildAgentPrompt` fonksiyonuna agent için özel prompt ekleyin
3. `ExecuteAgent` fonksiyonunda yeni agent'ı işleyin

