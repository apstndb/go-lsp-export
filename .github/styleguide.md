# Style guide for go-lsp-export

## Generated files (`protocol/ts*.go`)

Files matching `protocol/ts*.go` are **machine-generated** from the VSCode LSP `metaModel.json`.

- All doc comments are **verbatim copies** from the upstream specification.
- Do **not** flag typos, grammar issues, or terminology inconsistencies in these comments.
- Do **not** suggest rewording doc comments.
- If a comment contains an error, it should be reported to `microsoft/vscode-languageserver-node`, not patched in this repository.

## Code generator (`protocol/generate/`)

When reviewing changes in `protocol/generate/`:

- Prefer minimal diffs vs upstream `gopls/internal/protocol/generate`.
- Fork-specific changes (e.g. `golang.org/x/exp/jsonrpc2`, `interface{}` instead of `any`) are intentional and should not be "modernized" back to upstream patterns.
- The generator intentionally does **not** patch spec-derived comments to avoid ongoing maintenance burden.

## General conventions

- Keep modifications minimal vs upstream.
- When adding changes, clearly separate them from upstream code so they can be identified during future merges.
- Do not edit `ts*.go` files directly; change the generator and rerun `go generate`.
