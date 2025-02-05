# AI Destekli Siber Güvenlik Port Analiz Aracı - 30 Günde 30 Hack Aracı Yarışması 1. Gün

[![1. Gün](https://img.shields.io/badge/Gün-1-brightgreen)](https://github.com/rasperon/portscanner)

Bu proje, 30 günde 30 hack aracı geliştirme yarışmamın bir parçası olarak geliştirilen ilk araçtır. Açık ağ portlarını analiz etmek ve detaylı bir güvenlik değerlendirme raporu sunmak için yapay zekanın gücünden yararlanır.  Potansiyel güvenlik açıklarını belirlemek ve öneriler sunmak için Gemini AI modelini kullanır.

## Özellikler

*   **AI Destekli Analiz:** Açık portları analiz etmek için Gemini AI modelini kullanır.
*   **Detaylı Güvenlik Değerlendirmesi:** Güvenlik etkileri ve açıklıkları hakkında kapsamlı bir rapor sunar.
*   **Risk Seviyesi Değerlendirmesi:** Her açık port için risk seviyesini (Düşük/Orta/Yüksek) belirler.
*   **Spesifik Öneriler:** Her portu güvenceye almak için özel öneriler sunar.
*   **Markdown Çıktısı:** Analiz raporunu bir Markdown dosyasına (`ai_output.md`) kaydeder.

## Kurulum

1.  **Depoyu klonlayın:**

    ```bash
    git clone https://github.com/rasperon/portscanner.git
    cd portscanner
    ```

2.  **Bağımlılıkları yükleyin:**

    ```bash
    go mod download
    ```

3.  **Gemini API Anahtarınızı ayarlayın:**

    *   [Google AI Studio](https://makersuite.google.com/)'dan bir Gemini API anahtarı edinin.
    *   **Önemli:** API anahtarınızı doğrudan koda yazMAYIN! Bunun yerine, bir ortam değişkeni olarak ayarlayın.

        ```bash
        export GEMINI_API_KEY="API_ANAHTARINIZ"
        ```

4.  **Uygulamayı çalıştırın:**

    ```bash
    go run main.go
    ```

## Kullanım

1.  **Açık port verilerinizi hazırlayın:** Aşağıdaki formatta açık portların bir listesini sağlayın:

    ```
    22/tcp açık  ssh
    80/tcp açık  http
    443/tcp açık  https
    ```

2.  **`api/api.go` dosyasındaki `AnalyzePorts` fonksiyonunu** açık port verileri ve istenen dille birlikte çalıştırın.

## Yapılandırma

`api/api.go` dosyasında aşağıdaki parametreleri yapılandırabilirsiniz:

*   **Model:** Bir Gemini modeli seçin (örneğin, `gemini-pro`).
*   **Sıcaklık (Temperature):** Daha yaratıcı veya daha az yaratıcı yanıtlar için sıcaklığı ayarlayın.
*   **Maksimum Çıktı Token Sayısı (Max Output Tokens):** Oluşturulan raporun maksimum uzunluğunu kontrol edin.

## Sorumluluk Reddi

Bu araç yalnızca eğitim ve etik hackleme amaçları için tasarlanmıştır. Sorumlu bir şekilde kullanın ve yalnızca test etme izniniz olan sistemlerde kullanın. Yazar, bu aracın herhangi bir kötüye kullanımından sorumlu değildir.

## Katkıda Bulunma

Katkılar memnuniyetle karşılanır! Lütfen hata düzeltmeleri, iyileştirmeler veya yeni özelliklerle ilgili çekme istekleri gönderin.

## Lisans

[MIT Lisansı](LICENSE)