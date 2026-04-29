# go-lsp-export

Personal fork of `gopls/internal/protocol` from `golang.org/x/tools`, extracted so external LSP servers can import the types.

## Repo structure

- `protocol/` — Main package with generated LSP types (`tsclient.go`, `tsserver.go`, `tsprotocol.go`, `tsjson.go`, `tsmethod.go`) and hand-written helpers (`protocol.go`, `mapper.go`, `uri.go`, etc.)
- `protocol/generate/` — Tool that generates `protocol/` from the VSCode LSP JSON spec
- `internal/` — Vendored upstream helpers kept to minimize dependencies: `diff`, `util/pathutil`, `util/safetoken`, `xcontext`
- `tools/` — Git submodule of `golang.org/x/tools` (upstream reference)
- `vscode-languageserver-node/` — Git submodule of `microsoft/vscode-languageserver-node` (used by generator)

## Key architectural facts

- Go 1.25.0+ required.
- JSON-RPC implementation uses `golang.org/x/exp/jsonrpc2` instead of the upstream internal `jsonrpc2` / `jsonrpc2_v2` packages.
- The `protocol` package exposes `Client` and `Server` interfaces plus dispatch helpers (`ClientHandler`, `ServerHandler`, `ClientDispatcher`, `ServerDispatcher`) wired to `golang.org/x/exp/jsonrpc2`.
- `go generate` in `protocol/` runs `go run ./generate`, which downloads the VSCode LSP repo, parses `metaModel.json`, and regenerates the `ts*.go` files.
  - Default LSP spec ref is hard-coded in `protocol/generate/main.go` (`lspGitRef = "release/protocol/3.17.6-next.9"`).
  - To use the local submodule instead of a fresh clone: `go run ./generate -d ../vscode-languageserver-node`
- Generated files carry a `// Code generated for LSP. DO NOT EDIT.` header and include the spec git hash.

## Testing

```bash
# Run all tests (generate tests are skipped by default)
go test ./...

# Run a single package
go test ./protocol

# Run generator tests (requires vscode-languageserver-node clone in $HOME/vscode-languageserver-node)
go test ./protocol/generate
```

## Dependencies

The remaining upstream inter-dependencies that were *not* removed:

- `github.com/apstndb/gotoolsdiff` (fork of `golang.org/x/tools/internal/diff`)
- `golang.org/x/tools` (for `internal/event`, etc.)
- `golang.org/x/exp/jsonrpc2`
- `golang.org/x/telemetry`

## Conventions

- Keep modifications minimal vs. upstream. Prefer preserving upstream code style and structure so future syncs are easy.
- When adding changes, clearly separate them from upstream code so they can be identified during future merges.
- Do not edit `ts*.go` files directly; change the generator in `protocol/generate/` and rerun `go generate`.
