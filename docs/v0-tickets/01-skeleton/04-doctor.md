# 01-04: Doctor Command

## Summary
Health check command. Validates git, config, API key presence and validity.

## Event Model Refs
- E3: Doctor (4 checks)

## Files to Edit
- `cli/cmd/doctor.go` — replace stub with checks

## Checks (in order)
1. `exec.LookPath("git")` → "✓ git: installed" / "✗ git: not found"
2. `config.Load()` succeeds → "✓ config: valid" / "✗ config: invalid"
3. `config.ActiveProvider().APIKey != ""` → "✓ API key: set" / "✗ API key: missing"
4. LLM test request (minimal prompt) → "✓ API key: valid" / "✗ API key: invalid. HTTP {status}"

## Acceptance Criteria
- `pear doctor` runs all 4 checks, prints results with ✓/✗ prefixes
- Exits 0 if all pass, exits 1 if any fail
- Check 4 requires a working LLM client — if LLM package isn't ready yet, stub with a TODO and test the first 3 checks

## Dependencies
- 01-02 (config package)
- 02-01 (LLM interface — for check 4, can stub initially)

## Notes
- Check 4 can use a minimal prompt like "respond with OK" with max_tokens=5
- Doctor is also called from init wizard (01-03) — expose as a callable function, not just a Cobra command
