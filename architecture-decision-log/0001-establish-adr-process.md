# 1. Establish ADR Process

---

Date: 2025-04-05

Status: proposed

Author: Oliver Schlüter

## Problem

The architecture of a software system is a critical aspect that influences its design, development, and maintenance. 
However, architecture decisions are often made in an ad-hoc manner, leading to inconsistencies and difficulties in understanding the rationale behind certain choices. 
This ADR aims to establish a structured approach to record architecture decisions, ensuring that they are documented, communicated, and accessible to all stakeholders.

**Alternatives:**

- Do not record architecture decisions
- Use a different format (e.g., plain text, PDF)
- Store ADRs in a different location
- Use a different framework or template for ADRs

## Decision

All architectural decisions must be recorded using Markdown-based ADRs following the template defined in this document.
ADRs will be stored in the `architecture-decision-log/` directory.
Filenames will use the format: `XXXX-title.md`, where `XXXX` is a zero-padded index (e.g., `0001-record-architecture-decisions.md`).
Related resources will go into `architecture-decision-log/resources/`, prefixed with the ADR number.

**Each ADR must include the following sections:**

- **Header**: Metadata about the decision, including the date, status, and author
- **Problem**: A brief description of the problem being addressed and the context in which the decision is made
- **Decision**: A clear and concise statement of the decision made
- **Consequences**: The implications of the decision, including any trade-offs or potential risks
- **References**: Any relevant documents, links, or resources that provide additional context or information about the decision

**The states of the ADR can be:**

- **Proposed**: The ADR is proposed and not yet accepted
- **Accepted**: The ADR is accepted and implemented
- **Superseded by XXXX**: The ADR is no longer relevant and has been replaced by a new decision
- **Supersedes XXXX**: The ADR replaces a previous decision

If the decision includes resources (e.g. images), they will be stored in the `architecture-decision-log/resources` directory.
All resource file names must be prefixed with the ADR number, e.g. `0001-architecture-diagram.png`.
This ensures that resources are easily identifiable and associated with the correct decision.

## Consequences

- Every time an architecture decision is made, there is a need to create a new ADR entry in the log
- Every architecture decision must be reviewed and accepted by all team members
- We do not modify existing ADR entries, but we can add new ones to clarify or update previous decisions

## References

- [The Architecture Decision Record Homepage](https://adr.github.io/)
- [Basics of Architecture Decision Records (ADR)](https://medium.com/@nolomokgosi/basics-of-architecture-decision-records-adr-e09e00c636c6)
- [Lesson 55 - Architecture Decision Records](https://www.youtube.com/watch?v=LMBqGPLvonU) 
- [Lesson 141 - Managing Architecture Decisions](https://www.youtube.com/watch?v=PoarX66AO5s) 
