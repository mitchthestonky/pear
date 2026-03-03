# 06-03: Git Hooks

## Summary
Install/uninstall a post-commit git hook that runs `pear review --commit HEAD`.

## Event Model Refs
- E10: Hook Management

## Files to Create
- `cli/hooks/hooks.go`

## Files to Edit
- `cli/cmd/hooks.go` — replace stub, wire to hooks package

## Functions
- `Install(repoRoot string) error`
- `Uninstall(repoRoot string) error`
- `IsInstalled(repoRoot string) bool`

## Install Logic
1. Path: `{repoRoot}/.git/hooks/post-commit`
2. If exists and contains `# pear-hook-start` → already installed, print message, return
3. If exists without marker → append pear block
4. If not exists → create new file with shebang
5. Write: `\n# pear-hook-start\npear review --commit HEAD\n# pear-hook-end\n`
6. `os.Chmod(path, 0755)`

## Uninstall Logic
1. Read file
2. Remove everything between `# pear-hook-start` and `# pear-hook-end` (inclusive)
3. If remaining content is only shebang or empty → delete file
4. Else → write back

## Acceptance Criteria
- `pear hooks install` creates executable hook file
- `pear hooks install` is idempotent (running twice doesn't duplicate)
- `pear hooks uninstall` cleanly removes pear lines
- Preserves other hooks in the same file
- Works when `.git/hooks/` directory doesn't exist yet

## Dependencies
- 03-03 (review command — hook calls `pear review --commit HEAD`)
