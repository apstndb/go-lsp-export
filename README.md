My personal fork of [gopls/internal/protocol](https://github.com/golang/tools/tree/master/gopls/internal/protocol) for generic LSP server development.

## Modifications

- Migrate JSON-RPC 2 implementation to [golang.org/x/exp/jsonrpc2](https://pkg.go.dev/golang.org/x/exp/jsonrpc2) from internal packages ([jsonrpc2](https://pkg.go.dev/golang.org/x/tools/internal/jsonrpc2) and [jsonrpc2_v2](https://pkg.go.dev/golang.org/x/tools/internal/jsonrpc2_v2)).
- Remove some internal dependencies.
- Vendor rest internal dependencies.
  - [diff](https://pkg.go.dev/golang.org/x/tools/internal/diff)
  - [pathutil](https://pkg.go.dev/golang.org/x/tools/gopls/internal/util/pathutil)
  - [safetoken](https://pkg.go.dev/golang.org/x/tools/gopls/internal/util/safetoken)
  - [xcontext](https://pkg.go.dev/golang.org/x/tools/internal/xcontext)

## Rationale

- [x/tools/gopls: export LSP types so they can be imported by external consumers `#67658`](https://github.com/golang/go/issues/67658)