# 1. Einführung eines ADR-Prozesses

> [!NOTE]
> Translated from the original English version using AI.
> The translation may not be perfect, but it should convey the same meaning and intent as the original text.
> Please refer to the original English version for the most accurate information.
> You can find the original version [here](../english/adr-001.md).

**Datum:** 05.04.2025  
**Status:** In Arbeit  
**Autor:** Oliver Schlüter

## Problemstellung

Die Architektur eines Softwaresystems ist ein entscheidender Aspekt, der dessen Entwurf, Entwicklung und Wartung beeinflusst.  
Allerdings werden Architekturentscheidungen häufig spontan getroffen, was zu Inkonsistenzen und Schwierigkeiten beim Nachvollziehen der Beweggründe führen kann.  
Dieses ADR (Architecture Decision Record) soll einen strukturierten Ansatz zur Dokumentation von Architekturentscheidungen etablieren, damit diese nachvollziehbar, kommuniziert und für alle Beteiligten zugänglich sind.

**Alternativen:**

- Architekturentscheidungen nicht dokumentieren
- Ein anderes Format verwenden (z.B. Klartext, PDF)
- ADRs an einem anderen Ort speichern
- Ein anderes Framework oder Template für ADRs verwenden

## Entscheidung

Alle Architekturentscheidungen müssen mithilfe von Markdown-basierten ADRs dokumentiert werden, die dem in diesem Dokument definierten Template folgen.  
ADRs werden im Verzeichnis `docs/src/adr/` des GitHub-Repositories gespeichert, um die Vorteile von Git-Historie und -Versionierung zu nutzen.  
Dateinamen folgen dem Format: `adr-XXXX.md`, wobei `XXXX` ein vierstellig mit Nullen aufgefüllter Index ist (z.B. `adr-0001.md`).  

**Jedes ADR muss folgende Abschnitte enthalten:**

- **Header**: Metadaten zur Entscheidung, inklusive Datum, Status und Autor
- **Problem**: Eine kurze Beschreibung des zu lösenden Problems und des Kontexts der Entscheidung
- **Entscheidung**: Eine klare und prägnante Darstellung der getroffenen Entscheidung
- **Konsequenzen**: Die Auswirkungen der Entscheidung, einschließlich möglicher Kompromisse oder Risiken
- **Referenzen**: Relevante Dokumente, Links oder Ressourcen, die zusätzlichen Kontext zur Entscheidung liefern

**Mögliche Zustände eines ADRs:**

- **Vorgeschlagen**: Die Entscheidung ist vorgeschlagen, aber noch nicht angenommen
- **Angenommen**: Die Entscheidung wurde akzeptiert und umgesetzt
- **Abgelöst durch XXXX**: Die Entscheidung ist nicht mehr relevant und wurde durch eine neue ersetzt
- **Ersetzt XXXX**: Die Entscheidung ersetzt eine vorherige

## Konsequenzen

- Bei jeder Architekturentscheidung muss ein neuer ADR-Eintrag im Log erstellt werden
- Jede Entscheidung muss vom gesamten Team überprüft und genehmigt werden, um als angenommen zu gelten
- Bestehende ADR-Einträge werden nicht verändert, jedoch können neue hinzugefügt werden, um frühere Entscheidungen zu klären oder zu aktualisieren

## Referenzen

- [Die Homepage der Architecture Decision Records (ADR)](https://adr.github.io/)
- [Grundlagen zu Architecture Decision Records (ADR)](https://medium.com/@nolomokgosi/basics-of-architecture-decision-records-adr-e09e00c636c6)
- [Lektion 55 – Architecture Decision Records](https://www.youtube.com/watch?v=LMBqGPLvonU)
- [Lektion 141 – Umgang mit Architekturentscheidungen](https://www.youtube.com/watch?v=PoarX66AO5s)  