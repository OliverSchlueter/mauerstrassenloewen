# 3. Einsatz von NATS für die Kommunikation zwischen Services

> [!NOTE]
> Translated from the original English version using AI.
> The translation may not be perfect, but it should convey the same meaning and intent as the original text.
> Please refer to the original English version for the most accurate information.
> You can find the original version [here](../english/adr-003.md).

**Datum:** 08.04.2025  
**Status:** In Arbeit  
**Autor:** Oliver Schlüter

## Problemstellung

Im Rahmen unserer Microservices-Architektur müssen die einzelnen Services häufig miteinander kommunizieren.  
Während sich einige Interaktionen gut für synchrone HTTP/gRPC-Aufrufe eignen (z. B. Anfragen zu Benutzerdaten), profitieren andere stark von asynchroner Nachrichtenübermittlung – etwa KI-Anfragen, Hintergrundverarbeitung, Benachrichtigungen oder lose gekoppelte, ereignisgesteuerte Workflows.

Wir möchten ein Nachrichtensystem einführen, das Folgendes ermöglicht:
- **Asynchrone, entkoppelte Kommunikation** zwischen Services
- Unterstützung für **Publish-Subscribe (Pub/Sub)** und **Message Queue**-Muster
- Einfache Einrichtung und Betrieb – geeignet für lokale Entwicklung und kleine Teams
- Gute Performance und Zuverlässigkeit für grundlegende eventbasierte Workflows

Wir haben mehrere Messaging-Lösungen in Betracht gezogen:
- **RabbitMQ**: Ausgereift und funktionsreich, aber schwergewichtiger und komplexer in der Konfiguration
- **Kafka**: Sehr leistungsstark für großvolumige, persistente Event-Streams, aber überdimensioniert für unseren Anwendungsfall und schwer lokal zu betreiben
- **Redis Streams / PubSub**: Einfach zu benutzen, bietet jedoch keine starken Messaging-Semantiken oder Garantien bezüglich der Zustellung
- **NATS**: Leichtgewichtig, extrem schnell, unterstützt verschiedene Messaging-Muster und bietet eine einfache Developer Experience

## Entscheidung

Wir haben entschieden, **NATS** als Nachrichtensystem für die Kommunikation zwischen unseren Services zu verwenden.

NATS ist eine leichtgewichtige, performante Messaging-Plattform, die Pub/Sub-, Request/Reply- und Streaming-Muster unterstützt.  
Sie lässt sich einfach mit Docker lokal betreiben und benötigt nur minimale Konfiguration.  
NATS bietet uns die nötige Flexibilität, um eine ereignisgesteuerte Kommunikation zwischen unseren Microservices umzusetzen, ohne dabei hohe betriebliche Komplexität einzuführen.  
Die Performance von NATS ist ideal für unseren kleinen bis mittleren Anwendungsbereich. Falls nötig, können wir später zu **NATS JetStream** übergehen, um Persistenz oder eine mindestens-einmal-Zustellung zu ermöglichen.

Wir werden NATS hauptsächlich einsetzen für:
- Das Publizieren von Domain-Events (z. B. „UserRegistered“, „LessonCompleted“)
- Das Auslösen asynchroner Workflows (z. B. Willkommens-E-Mails versenden)
- Die Entkopplung von Services, bei denen keine synchrone Antwort erforderlich ist

## Konsequenzen

**Positiv:**
- Ermöglicht lose gekoppelte, skalierbare Kommunikation zwischen Services
- Einfach einzurichten und zu betreiben – ideal für lokale Entwicklung und CI
- Hohe Performance bei geringem Ressourcenverbrauch
- Unterstützung mehrerer Messaging-Muster (Pub/Sub, Request/Reply)
- Leicht zu erlernen und zu integrieren

**Negativ:**
- Weniger Enterprise-Funktionen im Vergleich zu Kafka oder RabbitMQ
- Der NATS-Server muss als gemeinsame Infrastrukturkomponente verwaltet und überwacht werden

## Referenzen

- [nats.io](https://nats.io/)
- [Cloud Native Messaging System – NATS Teil 1](https://hemantjain.medium.com/cloud-native-messaging-system-nats-part-1-ea4f25171ee9)
- [Lektion 1 – Event-Driven Architecture: Request/Reply Pattern](https://www.youtube.com/watch?v=3bxAm3XIFmk)
- [Lektion 165 – Event-Driven Architecture](https://www.youtube.com/watch?v=P0aUV4ixvBQ)  