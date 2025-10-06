# 🏨 Sentiric Vertical Hospitality Service - Mantık ve Akış Mimarisi

**Stratejik Rol:** Konaklama ve otelcilik sektörüne özel iş mantığını (otel arama, rezervasyon yapma, oda durumu kontrolü) sunar. Agent'lar bu servisi, sektör dışı karmaşık mantıkla uğraşmadan kullanır.

---

## 1. Temel Akış: Otel Arama (FindHotels)

```mermaid
sequenceDiagram
    participant Agent as Agent Service
    participant VHS as Hospitality Service
    participant Postgres as Kurumsal DB
    
    Agent->>VHS: FindHotels(location="Antalya", check_in="2025-11-01")
    
    Note over VHS: 1. DB Sorgusu veya Harici API Çağrısı
    VHS->>Postgres: SELECT * FROM hotels WHERE...
    Postgres-->>VHS: Hotel Data
    
    Note over VHS: 2. Sonuçları Mesaj Formatına Çevir
    VHS-->>Agent: FindHotelsResponse(results: [HotelResult, ...])
```

## 2. İş Mantığı Alanları

* Hotel Envanteri: Müsaitlik ve fiyat sorgulamaları.
* Rezervasyon Yönetimi: BookRoom RPC'si, harici sistemlere API çağrısı yaparak kesin rezervasyonu yönetir.
* Müşteri İlişkileri: Sadakat puanları, VIP statüsü gibi kullanıcı verilerini işlemek için user-service'e bağımlı olacaktır.