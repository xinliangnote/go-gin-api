## Title
fix(handlergen): prevent panic when generating controller methods with incomplete interface metadata

## Related Issue
- Closes #86

## Background
When running the controller/handler generator (`cmd/handlergen`), the process could panic with:

`panic: runtime error: invalid memory address or nil pointer dereference`

This happened when the target handler interface in `internal/api/<name>/handler.go` did not satisfy generator assumptions (for example: missing method list metadata or missing required method annotations/comments).

## Root Cause
`cmd/handlergen/main.go` made unsafe assumptions:

1. `interfaceType.Methods` is always non-nil.
2. Every method always contains at least 3 leading comments:
   - summary line
   - `@Tags`
   - `@Router`
3. The first comment always contains the method name for split-based parsing.

These assumptions caused unsafe dereference/index access in edge cases.

## What Changed
Updated `cmd/handlergen/main.go` to make generation defensive and non-crashing:

1. Added a nil guard for `interfaceType.Methods`.
   - If missing, skip that interface and log a clear message.
2. Added annotation length guard for each method.
   - Require at least 3 leading comments.
   - If not enough, skip method and log a precise warning with method name and actual count.
3. Added safe fallback for summary/description text generation.
   - If method name split cannot extract suffix, fallback to method name instead of indexing blindly.

## Behavior After Fix
- Generator no longer panics on incomplete handler definitions.
- Invalid/incomplete methods are skipped with readable logs.
- Valid methods continue generating `func_*.go` as before.

## Verification
Executed locally:

1. `go test ./cmd/handlergen/...`
   - Result: pass (compile check, no test files)
2. `go test ./...`
   - Existing unrelated failure remains in `pkg/mail`:
     - `TestSend` failed due to external SMTP auth (`535 Error: authentication failed`)
   - No new failures introduced by this change in `cmd/handlergen`.

## Risk Assessment
- Low risk.
- Change is isolated to generator command path.
- Runtime API/business logic is not modified.

## Rollback Plan
If any unexpected generation behavior appears, revert this commit to restore previous generator behavior.

