# 5. Ollama für das Bereitstellen von KI-Modellen

> [!NOTE]
> Translated from the original English version using AI.
> The translation may not be perfect, but it should convey the same meaning and intent as the original text.
> Please refer to the original English version for the most accurate information.
> You can find the original version [here](../english/adr-005.md).

**Datum:** 04.07.2025  
**Status:** Entwurf  
**Autor:** Oliver Schlüter

## Problem

Wir benötigen eine leichtgewichtige, entwicklerfreundliche Lösung zur Bereitstellung von großen Sprachmodellen (LLMs). Das System sollte:

- Einfach zu deployen und zu integrieren sein.
- In Offline- und Self-Hosting-Setups gut funktionieren, um Datenschutz zu gewährleisten.
- Schnelle Iteration und Entwicklung ermöglichen, ohne den operativen Aufwand traditioneller ML-Server.
- Eine stabile API für den programmgesteuerten Zugriff auf Modelle bereitstellen.
- Keine unnötigen Betriebs- oder Lizenzkosten verursachen.

Bestehende Optionen wie die API von OpenAI oder vollständige ML-Serverlösungen sind entweder cloudabhängig, teuer oder für unsere aktuellen Anforderungen zu komplex.

## Entscheidung

Wir werden **Ollama** für das Bereitstellen von KI-Modellen in Entwicklungs- und leichtgewichtigen Produktionsumgebungen verwenden.

Ollama bietet ein einfaches CLI und eine HTTP-API zum Ausführen und Interagieren mit Modellen wie Llama, Mistral, Gemma usw. Es ermöglicht uns:

- Die Containerisierung von Modellen mit Docker für konsistente CI/CD- und Deployment-Prozesse.
- Einfaches Wechseln zwischen verschiedenen Modellen oder Versionen mittels `ollama pull` und `ollama run`.
- Effizientes Ausführen quantisierter Modelle – auch auf ressourcenbeschränkten Maschinen.
- Vollständig kostenfreien Betrieb – ohne Lizenzkosten oder API-Gebühren.

Ollama passt sehr gut zu unseren Anforderungen hinsichtlich Benutzerfreundlichkeit, lokalem Fokus, schneller Iteration – und null Laufzeitkosten.

## Konsequenzen

**Positiv**

- **Benutzerfreundlichkeit**: Schnelle Einrichtung von Modellen mit minimaler Konfiguration.
- **Entwicklereffizienz**: Entwickler können lokal LLM-Funktionen testen und iterieren – ohne Cloud-Abhängigkeiten.
- **Datenschutz**: Sensible Daten bleiben lokal; keine externe API-Nutzung erforderlich.
- **Portabilität**: Docker-Support ermöglicht konsistente Umgebungen über Teams und Deployments hinweg.
- **Kostenfreiheit**: Keine API-Gebühren oder kommerziellen Lizenzen nötig.
- **Schnelles Prototyping**: Kurze Feedbackzyklen und Experimente möglich.

**Negativ**

- **Begrenzte Flexibilität**: Weniger anpassbar als vollwertige Servicelösungen.
- **Skalierbarkeitsgrenzen**: Möglicherweise ungeeignet für stark frequentierten Produktivbetrieb.
- **Ökosystem-Abhängigkeit**: Gebunden an das Ollama-Format und dessen Roadmap; begrenzter Support für benutzerdefinierte Modelle außerhalb des Ökosystems.

## Referenzen

- [Ollama](https://ollama.com/)