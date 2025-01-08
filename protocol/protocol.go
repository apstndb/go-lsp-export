// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/exp/jsonrpc2"
	"golang.org/x/telemetry/crashmonitor"

	"github.com/apstndb/go-lsp-export/internal/util/bug"
	"github.com/apstndb/go-lsp-export/internal/xcontext"
)

var (
	// RequestCancelledError should be used when a request is cancelled early.
	RequestCancelledError = jsonrpc2.NewError(-32800, "JSON RPC cancelled")
)

type ClientCloser interface {
	Client
	io.Closer
}

type connSender interface {
	io.Closer

	Notify(ctx context.Context, method string, params interface{}) error
	Call(ctx context.Context, method string, params interface{}) *jsonrpc2.AsyncCall
}

type clientDispatcher struct {
	sender connSender
}

func (c *clientDispatcher) Close() error {
	return c.sender.Close()
}

// ClientDispatcher returns a Client that dispatches LSP requests across the
// given jsonrpc2 connection.
func ClientDispatcher(conn *jsonrpc2.Connection) ClientCloser {
	return &clientDispatcher{sender: conn}
}

type clientConn = jsonrpc2.Connection

/*
func ClientDispatcherV2(conn *jsonrpc2_v2.Connection) ClientCloser {
	return &clientDispatcher{clientConnV2{conn}}
}

type clientConnV2 struct {
	conn *jsonrpc2_v2.Connection
}

func (c clientConnV2) Close() error {
	return c.conn.Close()
}

func (c clientConnV2) Notify(ctx context.Context, method string, params interface{}) error {
	return c.conn.Notify(ctx, method, params)
}

func (c clientConnV2) Call(ctx context.Context, method string, params interface{}, result interface{}) error {
	call := c.conn.Call(ctx, method, params)
	err := call.Await(ctx, result)
	if ctx.Err() != nil {
		detached := xcontext.Detach(ctx)
		c.conn.Notify(detached, "$/cancelRequest", &CancelParams{ID: call.ID().Raw()})
	}
	return err
}
*/

// ServerDispatcher returns a Server that dispatches LSP requests across the
// given jsonrpc2 connection.
func ServerDispatcher(conn *jsonrpc2.Connection) Server {
	return &serverDispatcher{sender: conn}
}

/*
func ServerDispatcherV2(conn *jsonrpc2_v2.Connection) Server {
	return &serverDispatcher{sender: clientConnV2{conn}}
}
*/

type serverDispatcher struct {
	sender connSender
}

func ClientHandler(client Client, handler jsonrpc2.Handler) jsonrpc2.Handler {
	return jsonrpc2.HandlerFunc(func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		if ctx.Err() != nil {
			return nil, RequestCancelledError
		}
		handled, resp, err := clientDispatch(ctx, client, req)
		if handled || err != nil {
			return resp, err
		}
		return handler.Handle(ctx, req)
	})
}

func ServerHandler(server Server, handler jsonrpc2.Handler) jsonrpc2.Handler {
	return jsonrpc2.HandlerFunc(func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		if ctx.Err() != nil {
			return nil, RequestCancelledError
		}
		handled, resp, err := serverDispatch(ctx, server, req)
		if handled || err != nil {
			return resp, err
		}
		return handler.Handle(ctx, req)
	})
}

/*
func Handlers(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return CancelHandler(
		jsonrpc2.AsyncHandler(
			jsonrpc2.MustReplyHandler(handler)))
}
*/

/*
func CancelHandler(handler jsonrpc2.Handler) jsonrpc2.Handler {
	handler, canceller := jsonrpc2.CancelHandler(handler)
	return jsonrpc2.HandlerFunc(func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		if req.Method != "$/cancelRequest" {
			// TODO(iancottrell): See if we can generate a reply for the request to be cancelled
			// at the point of cancellation rather than waiting for gopls to naturally reply.
			// To do that, we need to keep track of whether a reply has been sent already and
			// be careful about racing between the two paths.
			// TODO(iancottrell): Add a test that watches the stream and verifies the response
			// for the cancelled request flows.
			replyWithDetachedContext := func(ctx context.Context, resp interface{}, err error) error {
				// https://microsoft.github.io/language-server-protocol/specifications/specification-current/#cancelRequest
				if ctx.Err() != nil && err == nil {
					err = RequestCancelledError
				}
				ctx = xcontext.Detach(ctx)
				return reply(ctx, resp, err)
			}
			return handler.Handle(ctx, req)
		}
		var params CancelParams
		if err := UnmarshalJSON(req.Params, &params); err != nil {
			return nil, sendParseError(ctx, err)
		}
		if n, ok := params.ID.(float64); ok {
			canceller(jsonrpc2.NewIntID(int64(n)))
		} else if s, ok := params.ID.(string); ok {
			canceller(jsonrpc2.NewStringID(s))
		} else {
			return nil, sendParseError(ctx, fmt.Errorf("request ID %v malformed", params.ID))
		}
		return nil, nil
	})
}
*/

func Call(ctx context.Context, conn jsonrpc2.Connection, method string, params interface{}, result interface{}) error {
	call := conn.Call(ctx, method, params)
	if ctx.Err() != nil {
		conn.Cancel(call.ID())
	}
	return call.Await(ctx, result)
}

func cancelCall(ctx context.Context, sender connSender, id jsonrpc2.ID) {
	ctx = xcontext.Detach(ctx)
	// Note that only *jsonrpc2.ID implements json.Marshaler.
	sender.Notify(ctx, "$/cancelRequest", &CancelParams{ID: &id})
}

// UnmarshalJSON unmarshals msg into the variable pointed to by
// params. In JSONRPC, optional messages may be
// "null", in which case it is a no-op.
func UnmarshalJSON(msg json.RawMessage, v any) error {
	if len(msg) == 0 || bytes.Equal(msg, []byte("null")) {
		return nil
	}
	return json.Unmarshal(msg, v)
}

func sendParseError(ctx context.Context, err error) error {
	return fmt.Errorf("%w: %s", jsonrpc2.ErrParse, err)
}

// NonNilSlice returns x, or an empty slice if x was nil.
//
// (Many slice fields of protocol structs must be non-nil
// to avoid being encoded as JSON "null".)
func NonNilSlice[T comparable](x []T) []T {
	if x == nil {
		return []T{}
	}
	return x
}

func recoverHandlerPanic(method string) {
	// Report panics in the handler goroutine,
	// unless we have enabled the monitor,
	// which reports all crashes.
	if !crashmonitor.Supported() {
		defer func() {
			if x := recover(); x != nil {
				bug.Reportf("panic in %s request", method)
				panic(x)
			}
		}()
	}
}
