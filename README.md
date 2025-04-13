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
   - Örnek kullanım: "İstanbul'un hava durumu nasıl?"

2. **Çeviri Agent'ı (translate-agent)**
   - Metin çevirisi yapar
   - Gerekli parametreler: `text`, `to`
   - Örnek kullanım: "Merhaba kelimesini İspanyolca'ya çevir"

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

1. `agent/registry.go` dosyasında `registerAgents` fonksiyonuna yeni agent'ı ekleyin:
```go
r.agents["yeni-agent"] = Agent{
    Name: "yeni-agent",
    Execute: func(params map[string]interface{}) (string, error) {
        // Agent'ın çalışma mantığı
        return "Sonuç", nil
    },
}
```

2. `gemini/client.go` dosyasında `buildPrompt` fonksiyonunda agent'ı tanımlayın:
```go
- yeni-agent: Açıklama, "parametre1" ve "parametre2" parametreleri gerekir.
```

## Proje Yapısı

```
.
├── agent/
│   └── registry.go      # Agent kayıt ve yönetimi
├── gemini/
│   └── client.go        # Gemini API entegrasyonu
├── orchestrator/
│   └── engine.go        # Agent orkestrasyonu
├── main.go              # Ana uygulama
└── README.md           # Dokümantasyon
```

