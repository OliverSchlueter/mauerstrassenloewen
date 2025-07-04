# 6. Verwendung von Qdrant für Retrieval-Augmented Generation (RAG) und Vektorsuche

!!!info Translated using AI
Translated from the original English version using AI.
The translation may not be perfect, but it should convey the same meaning and intent as the original text.
Please refer to the original English version for the most accurate information.
You can find the original version [here](../english/adr-006.md).
!!!

**Datum:** 04.07.2025  
**Status:** Entwurf  
**Autor:** Oliver Schlüter

## Problem

Wir benötigen eine Vektor-Datenbank zur Unterstützung von Retrieval-Augmented Generation (RAG) und semantischer Suche. Das System muss in der Lage sein:

- Hochdimensionale Vektor-Embeddings zu speichern und effizient abzufragen.
- Hybride Suche (semantisch + Schlüsselwortfilterung) zu ermöglichen.
- Mit wachsenden Dokumentensammlungen zu skalieren.
- Sich leicht in unsere KI-Serving- und Embedding-Pipelines integrieren zu lassen.
- Kostenlos nutzbar und sowohl lokal als auch in der Cloud einsetzbar zu sein.

Alternativen wie Pinecone, Weaviate und Vespa bieten ähnliche Funktionen, haben jedoch Nachteile hinsichtlich Kosten, Komplexität oder Cloud-Abhängigkeit.  
Wir möchten eine Lösung, die schnell, produktionsreif, quelloffen ist und uns vollständige Kontrolle über Deployment und Daten gibt.

## Entscheidung

Wir werden **Qdrant** als unsere Vektor-Datenbank für RAG- und semantische Suchfunktionen verwenden.

Qdrant bietet eine leistungsstarke, produktionsreife Vektorsuchmaschine mit folgenden Vorteilen:

- Native Unterstützung für Filter, Metadaten und hybride Abfragen (Vektor + Payload).
- Docker-basierte Bereitstellung und Cloud-unabhängige Skalierbarkeit.
- Hervorragende Performance bei ANN-Suchen (Approximate Nearest Neighbor).
- REST-API mit guter Dokumentation.
- Kompatibel mit gängigen Embedding-Modellen.
- Vollständig Open Source und kostenlos nutzbar, mit aktiver Community und optionalem kommerziellen Support.

## Konsequenzen

**Positiv:**

- **Effiziente RAG-Implementierung**: Schnelle semantische Abfrage relevanter Dokumente zur Verbesserung von LLM-Antworten.
- **Benutzerfreundlichkeit**: Klare APIs und SDKs erleichtern die Integration in unsere Pipeline.
- **Skalierbar und schnell**: Hochperformantes Vektorindexing mit Unterstützung für Filter und Echtzeit-Updates.
- **Kostenlos und Open Source**: Keine Lizenz- oder Nutzungskosten; selbst gehostet möglich.
- **Docker- und Kubernetes-freundlich**: Einfach in Entwicklungs-, Staging- und Produktionsumgebungen einsetzbar.

**Negativ:**

- **Index-Feinabstimmung**: Performance-Tuning (z. B. HNSW-Parameter) kann etwas Experimentieren erfordern.
- **Lernkurve**: Etwas Erfahrung mit vektorbasierter Suche ist notwendig.
- **Kein integriertes UI**: Im Gegensatz zu manchen Wettbewerbern (z. B. Weaviate) fehlt Qdrant ein integriertes UI zur Datenexploration (externe Tools sind jedoch verfügbar).

## Referenzen

- [Qdrant](https://qdrant.tech/)