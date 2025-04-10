# 2. Einführung einer Microservices-Architektur

> [!NOTE]
> Translated from the original English version using AI.
> The translation may not be perfect, but it should convey the same meaning and intent as the original text.
> Please refer to the original English version for the most accurate information.
> You can find the original version [here](adr-002.md).

**Datum:** 08.04.2025  
**Status:** In Arbeit  
**Autor:** Oliver Schlüter

## Problemstellung

Wir entwickeln eine Webanwendung im Rahmen eines Hochschulprojekts. Obwohl wir nur ein Team aus vier Entwicklern sind, möchten wir das System so entwerfen, als wäre es eine produktionsreife Anwendung aus der realen Welt.

Unsere Ziele sind:
- Praktische Erfahrung mit modernen Softwarearchitektur-Mustern sammeln
- Den Umgang mit cloud-nativen Tools und DevOps-Praktiken erlernen
- Die betrieblichen und technischen Kompromisse von Microservices verstehen

Die Anwendung besteht aus mehreren logischen Domänen (z.fflags.B. Backend, KI, Simulation, Mail), was eine natürliche Trennung der Verantwortlichkeiten nahelegt.

Wir haben folgende Architekturansätze in Betracht gezogen:

**Monolithische Architektur**:

Einfach zu implementieren und zu deployen, ideal für kleine Teams und enge Zeitrahmen. Allerdings sammeln wir dabei keine Erfahrung mit verteilten Systemen oder Skalierungsmustern.

**Modularer Monolith**:

Bietet gewisse Struktur und Trennung der Zuständigkeiten bei gleichzeitig einfacher Bereitstellung. Ein guter Kompromiss, aber er fordert uns nicht im Hinblick auf Service-Unabhängigkeit oder infrastrukturelle Fragestellungen heraus.

**Microservices-Architektur** *(ausgewählt)*:

Komplexer in Einrichtung und Betrieb, aber ermöglicht realistische Erfahrungen mit Service-Isolation, API-Design, Kommunikation zwischen Diensten und Infrastrukturautomatisierung.

## Entscheidung

Wir werden für unsere Anwendung eine Microservices-Architektur einführen.  
Jeder Service wird einen spezifischen fachlichen Bereich abbilden und unabhängig entwickelt, bereitgestellt und getestet.  
Die Kommunikation zwischen Diensten erfolgt primär über HTTP/REST oder CloudEvents/NATS. Asynchrone Kommunikation wird dort eingesetzt, wo eine Entkopplung und eventual consistency von Vorteil sind.  
Jeder Service verwaltet seine eigene Datenbank oder sein eigenes Schema, um lose Kopplung und klare Domänenverantwortung sicherzustellen.

Zur Unterstützung dieser Architektur werden wir jeden Service mit Docker containerisieren und lokal mit Docker Compose orchestrieren. Optional evaluieren wir Kubernetes, sofern die Zeit es erlaubt.  
Zusätzlich implementieren wir grundlegende Observability-Features wie Health Checks, zentrales Logging und Monitoring. CI/CD-Pipelines werden mit GitHub Actions aufgesetzt, um automatisiertes Testen und Deployment zu ermöglichen.

## Konsequenzen

**Positiv:**
- Realitätsnahe Erfahrung mit produktionsreifer Architektur
- Bessere Trennung der Domänen und strukturierterer Code
- Gelegenheit, mit modernen Tools und Praktiken zu arbeiten (Container, Service Discovery, CI/CD usw.)
- Skalierbare und wartbare Struktur

**Negativ:**
- Erhöhte Komplexität für ein kleines Team und ein zeitlich begrenztes Projekt
- Mehr Aufwand für Infrastruktur und Betrieb
- Gefahr von Overengineering
- Höhere Einstiegshürde und langsamerer Fortschritt zu Beginn

## Referenzen

- [microservices.io](https://microservices.io/)
- [Microservices vs. monolithische Architektur](https://www.atlassian.com/microservices/microservices-architecture/microservices-vs-monolith)
- [Lektion 162 – Microservices Architecture](https://www.youtube.com/watch?v=UZQMUiVqpFs&t=55s)  