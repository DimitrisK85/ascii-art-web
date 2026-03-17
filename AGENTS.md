# AGENTS.md

## Project  ascii-art-web

### Project Overview
 - This repository contains ascii-art-web 

## AGENT usage

1. We will analyze each project together
2. Discuss about projects requirements
3. Agree on a plan to follow for completion
4. Write down simple tasks to follow step by step until project completion
5. When we start implementing the code keep a log in task_decomposition_index.txt
6. Prompt me to think to the right direction
7. Be a guide and tutor so i can learn from you and solidify - amplify  my knowledge



## The project is audit-driven:
1. Correctness against the subject examples is mandatory
2. The final output must exactly match expected results

## How the Author Works
### When contributing to this repository, follow these principles:
1. Structure before code
 - Always agree on folder structure, responsibilities, and data flow before writing logic
 - Do not invent new layers or abstractions unless explicitly requested
 - Specification first

2. The project subject and PRD are the source of truth
3. Tests reflect verbatim examples from the subject
4. No extra behavior beyond what is specified
5. Simple > Clever
6. Prefer readable, explicit Go code
7. Avoid premature optimization or “smart” tricks
8. No unnecessary interfaces, generics, or reflection


## TDD Workflow (Mandatory)
### This project is developed using TDD (Test-Driven Development):
 - Red — Write a failing test that describes the intended behavior
 - Keep tests aligned with the subject examples and project rules
 - Green — Implement the smallest possible change to make the test pass
 - Do not add extra behavior beyond what the test/spec requires
 - Keep changes focused to the rule being addressed
 - Clean — Refactor for clarity with all tests still passing
 - Refactors must not change behavior
 - Preserve the canonical structure unless explicitly discussed

## AI Usage Tracking
### AI NOTE
 - Starting now, contributors must keep an AI usage index specifically for how AI is used during task decomposition. This is an audit requirement and may be checked later.

#### Required Setup:
1. Create folder: ai/
2. Maintain a running log file:
   ai/task-decomposition-index.txt
3. Each Log Entry Must Include:
 - Tool / model used
 - Date
 - What was asked
 - What was copied vs. edited
 - Which task card it affected
4. Rules
 - Log entries must be factual and transparent
 - Do not retroactively fabricate entries
 - Keep entries concise but complete enough for audit review

## Testing Philosophy
Audit Tests

Golden Tests

### Tests are:
1. Unit tests for TDD until we have a working code
2. End-to-end to test like the auditors would test the code
3. Using Go’s standard testing package only
4. TDD Reminder:
 - All changes should follow the Red → Green → Clean loop. If a change cannot be expressed as a failing test first, stop and reconsider the requirement.

## Constraints
Standard library only (no third-party packages)


## What Agents Should NOT Do
1. Do not invent new rules
2. Do not silently “fix” input beyond the spec
3. Do not refactor structure without approval
4. Do not optimize for performance over clarity
5. Do not assume behavior not explicitly stated

## Preferred Contribution Style
1. When making changes:
 - Explain what rule or test is being addressed
 - Explain where the change belongs
 - Keep commits small and focused
 - Prefer correctness and clarity over elegance

## Final Note
1. If there is ambiguity:
 - Follow the project subject
 - Then follow existing tests
 - If still unclear, stop and ask before implementing