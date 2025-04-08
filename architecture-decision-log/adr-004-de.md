# 4. MongoDB als primäre Datenbank

> [!NOTE]
> Translated from the original English version using AI.
> The translation may not be perfect, but it should convey the same meaning and intent as the original text.
> Please refer to the original English version for the most accurate information.
> You can find the original version [here](adr-004.md).

**Datum:** 08.04.2025  
**Status:** In Arbeit  
**Autor:** Oliver Schlüter

## Problemstellung

Für unser Projekt benötigen wir eine Datenbanklösung, die eine Vielzahl von Geschäftsdaten-Typen verarbeiten kann, flexibel genug ist, um sich während der Entwicklung zu verändern, und skalierbar ist, wenn wir im Laufe der Zeit mehr Funktionalitäten hinzufügen.

Wir haben folgende Datenbankoptionen zur Speicherung allgemeiner Geschäftsdaten in Betracht gezogen:
- **Relationale Datenbanken (z. B. PostgreSQL, MySQL):** Strukturierte, ausgereifte und zuverlässige Systeme. Sie sind jedoch oft unflexibel bei Schemaänderungen und erfordern möglicherweise komplexe Migrationen, um sich entwickelnde Datenmodelle zu unterstützen.
- **NoSQL-Dokumentdatenbanken (z. B. MongoDB):** Sehr flexibel, da sie es uns ermöglichen, semi-strukturierte Daten mit dynamischen Schemata zu speichern und bei Bedarf horizontal zu skalieren.
- **Key-Value Stores (z. B. Redis, DynamoDB):** Gut für Caching und schnelle Suchen, aber nicht ideal für die Verarbeitung komplexer Geschäftsdaten mit variierenden Beziehungen.

## Entscheidung

Wir haben beschlossen, **MongoDB** zur Speicherung unserer allgemeinen Geschäftsdaten zu verwenden.  
MongoDB ermöglicht es uns, Daten in einem flexiblen, schemafreien Format als JSON-ähnliche Dokumente zu speichern, was uns die Freiheit gibt, die Struktur unserer Daten schnell weiterzuentwickeln, während die Anwendung wächst.  
Diese Flexibilität ist wichtig, da unser Projekt voraussichtlich mehrere Iterationen durchlaufen wird, und MongoDB ermöglicht es uns, das Datenmodell ohne komplexe Migrationen oder Einschränkungen anzupassen.

Hauptgründe für die Wahl von MongoDB:
- **Schema-Flexibilität:** Wir können Daten in einem flexiblen, JSON-ähnlichen Format (BSON) speichern, ohne dass ein festes Schema zu Beginn erforderlich ist.
- **Horizontale Skalierbarkeit:** Die native Sharding-Funktion von MongoDB ermöglicht es, horizontal zu skalieren, wenn wir mehr Daten oder Services hinzufügen.
- **Benutzerfreundlichkeit:** MongoDB bietet eine einfache Abfragesprache und lässt sich gut in den modernen Tech-Stack integrieren, den wir verwenden.
- **Gutes Ökosystem und Community-Unterstützung:** MongoDB bietet hervorragende Dokumentation, Client-Bibliotheken und eine starke Community.
- **Datenmodell-Anpassung:** Das dokumentenbasierte Datenmodell von MongoDB passt gut zu unserem Bedarf, Entitäten darzustellen.

Jeder Microservice wird seine eigene MongoDB-Datenbank haben, um eine lose Kopplung zwischen den Services sicherzustellen, und die Datenbank wird von jedem Service unabhängig verwaltet.

## Konsequenzen

**Positiv:**
- Flexibilität bei der Speicherung dynamischer und sich entwickelnder Geschäftsdaten
- Einfache Skalierbarkeit, wenn die Anwendung wächst, insbesondere in Szenarien mit horizontaler Skalierung
- Schnell und effizient für leseintensive Anwendungen mit unterschiedlichen Datentypen
- Einfache Integration mit modernen Entwicklungsframeworks und Programmiersprachen
- Gute Entwicklererfahrung mit MongoDB-Treibern und nativer Abfragesprache

**Negativ:**
- Mögliche Herausforderungen bei der Datenkonsistenz in verteilten Umgebungen (z. B. Probleme mit der eventual consistency in einigen Szenarien)
- Nicht so stark in der Handhabung komplexer transaktionaler Anforderungen über mehrere Services hinweg (im Vergleich zu relationalen Datenbanken)
- MongoDB bietet nicht dieselben starken ACID-Garantien wie relationale Datenbanken, was uns dazu zwingen könnte, bestimmte Operationen (z. B. mehrstufige Transaktionen oder Joins) anders zu handhaben

## Referenzen

- [MongoDB.com](https://www.mongodb.com/)  