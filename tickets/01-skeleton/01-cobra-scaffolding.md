# 01-01: Cobra CLI Scaffolding

## Summary
Set up the Go module and Cobra command tree with all subcommands stubbed.

## Event Model Refs
- E1: Application Bootstrap → CLI.RouteCommand

## Files to Create
- `cli/main.go` — entry point, executes root command
- `cli/cmd/root.go` — root command (no args → interactive mode)
- `cli/cmd/watch.go` — `pear watch` stub
- `cli/cmd/ask.go` — `pear ask "question"` stub
- `cli/cmd/review.go` — `pear review [--commit HEAD]` stub
- `cli/cmd/teach.go` — `pear teach [topic]` stub
- `cli/cmd/progress.go` — `pear progress` stub
- `cli/cmd/init_cmd.go` — `pear init` stub
- `cli/cmd/doctor.go` — `pear doctor` stub
- `cli/cmd/hooks.go` — `pear hooks install/uninstall` stub
- `cli/go.mod` — module `github.com/pearcode/pear`

## Acceptance Criteria
- `go build` produces a `pear` binary
- `pear --help` shows all subcommands with descriptions
- `pear watch`, `pear ask "test"`, `pear review`, `pear teach`, `pear progress`, `pear init`, `pear doctor`, `pear hooks install`, `pear hooks uninstall` all print "not implemented yet" and exit 0
- `pear review --commit HEAD` parses the flag correctly
- `pear teach goroutines` captures the topic arg

## Dependencies
- None (first ticket)

## Notes
- Use `cobra-cli` or hand-write — either is fine
- Root command's `Run` will eventually launch TUI interactive mode (ticket 04-05)
- Keep stubs minimal: `fmt.Println("not implemented yet")`
