# 7. Grafana Loki für Logging

> [!NOTE]
> Translated from the original English version using AI.
> The translation may not be perfect, but it should convey the same meaning and intent as the original text.
> Please refer to the original English version for the most accurate information.
> You can find the original version [here](../english/adr-007.md).

**Datum:** 04.07.2025  
**Status:** Entwurf  
**Autor:** Oliver Schlüter

## Problem

Wir benötigen eine zentrale, effiziente und skalierbare Logging-Lösung, die es uns ermöglicht:

- Logs aus mehreren Services und Umgebungen zu aggregieren.
- Logs mit strukturierten Metadaten (z. B. Service, Umgebung, Level) zu durchsuchen und zu visualisieren.
- Logs mit Metriken und Traces zu korrelieren, um vollständige Observability zu erreichen.
- Die Kosten gering zu halten, auch bei hohem Logvolumen.
- Komplexe Infrastrukturen mit hohem Ressourcenbedarf (z. B. ELK-Stack) zu vermeiden.

Traditionelle Lösungen wie der ELK-Stack (Elasticsearch, Logstash, Kibana) oder Cloud-basierte Logging-Dienste (z. B. Datadog, Logz.io) bringen entweder hohe Kosten, operative Komplexität oder eingeschränkte Flexibilität für On-Premise-Deployments mit sich.

## Entscheidung

Wir werden **Grafana Loki** als unser zentrales System zur Logaggregation und -abfrage einsetzen.

Loki ist ein horizontal skalierbares, hochverfügbares Logaggregationssystem, das von Prometheus inspiriert wurde.  
Es wurde auf Kosteneffizienz ausgelegt und integriert sich nahtlos in das Grafana-Ökosystem.

**Hauptgründe für die Wahl von Loki:**

- Label-basiertes Indexing (kein Volltext), was Speicherbedarf und Betriebskosten reduziert.
- Enge Integration mit Grafana, wodurch einheitliche Dashboards für Metriken, Logs und Traces möglich sind.
- Einfache Architektur, leicht deploybar via Docker, Kubernetes oder als Standalone-Binary.
- Unterstützt Promtail, Fluent Bit und andere Log Shipper zur Log-Erfassung.
- Open Source und kostenlos, mit optionalen Enterprise-Funktionen bei Bedarf.

## Konsequenzen

**Positiv:**

- **Kosteneffizient**: Minimaler Indexierungsaufwand und effiziente Speicherung senken Infrastrukturkosten.
- **Skalierbar**: Funktioniert gut mit Kubernetes, Docker und verteilten Umgebungen.
- **Vereinheitlichte Observability**: Logs lassen sich einfach mit Metriken und Traces in Grafana korrelieren.
- **Flexibel**: Logs können aus Dateien, journald, Docker oder systemd mit Promtail oder Fluent Bit gesammelt werden.
- **Open Source**: Kein Vendor Lock-in; vollständig selbst gehostet und kostenlos nutzbar.
- **Einfache Wartung**: Deutlich einfacher zu betreiben als ELK-basierte Setups.

**Negativ:**

- **Eingeschränkte Volltextsuche**: Loki ist für labelbasierte Suche optimiert, nicht für Volltextindexierung (unterstützt aber Regex und Filter).
- **Lernkurve bei Labels**: Effiziente Abfragen hängen vom Label-Design und der Vertrautheit mit Lokis Abfragesprache (LogQL) ab.
- **Kein natives Log-Alerting**: Für Benachrichtigungen aus Logs ist eine Kombination mit Grafana Alerting oder externen Tools erforderlich.

## Referenzen

- [Grafana Loki](https://grafana.com/oss/loki/)