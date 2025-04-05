# 1. Record architecture decisions

---

Date: 2025-04-05

Status: draft

Author: Oliver Schlüter

## Problem

The architecture of a software system is a critical aspect that influences its design, development, and maintenance. 
However, architecture decisions are often made in an ad-hoc manner, leading to inconsistencies and difficulties in understanding the rationale behind certain choices. 
This ADR aims to establish a structured approach to record architecture decisions, ensuring that they are documented, communicated, and accessible to all stakeholders.

**Alternatives:**

* Use the architecture decision record framework for documenting architecture decisions
* Do not record architecture decisions: This would lead to a lack of documentation and understanding of the system's architecture
* Use a different format for recording decisions: This could lead to inconsistencies and difficulties in finding relevant information
* Use a different location for the ADRL: This could lead to confusion and make it harder for stakeholders to find the information they need

## Decision

Whenever an architecture decision is made, we will store it in an architecture decision record log in the `architecture-decision-log` directory of the project repository. 
The content will be in Markdown format and follows this naming convention: `XXXX-new-decision.md`. Example: `0001-record-architecture-decisions.md`.

Each ADR must include the following sections:

- **Header**: Metadata about the decision, including the date, status, and author
- **Problem**: A brief description of the problem being addressed and the context in which the decision is made
- **Decision**: A clear and concise statement of the decision made
- **Consequences**: The implications of the decision, including any trade-offs or potential risks
- **References**: Any relevant documents, links, or resources that provide additional context or information about the decision

If the decision includes resources (e.g. images), they will be stored in the `architecture-decision-log/resources` directory. 
All resource file names must be prefixed with the ADR number, e.g. `0001-architecture-diagram.png`. 
This ensures that resources are easily identifiable and associated with the correct decision.

## Consequences

* Every time an architecture decision is made, there is a need to create a new ADR entry in the log
* Every architecture decision must be reviewed and accepted by all team members
* We do not modify existing ADR entries, but we can add new ones to clarify or update previous decisions

## References

* [The Architecture Decision Record Homepage](https://adr.github.io/)
* [Basics of Architecture Decision Records (ADR)](https://medium.com/@nolomokgosi/basics-of-architecture-decision-records-adr-e09e00c636c6)

