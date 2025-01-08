// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated for LSP. DO NOT EDIT.

package protocol

// Code generated from protocol/metaModel.json at ref release/protocol/3.17.6-next.9 (hash c94395b5da53729e6dff931293b051009ccaaaa4).
// https://github.com/microsoft/vscode-languageserver-node/blob/release/protocol/3.17.6-next.9/protocol/metaModel.json
// LSP metaData.version = 3.17.0.

import (
	"context"

	"golang.org/x/exp/jsonrpc2"
)

type Client interface {
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#logTrace
	LogTrace(ctx context.Context, params *LogTraceParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#progress
	Progress(ctx context.Context, params *ProgressParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#client_registerCapability
	RegisterCapability(ctx context.Context, params *RegistrationParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#client_unregisterCapability
	UnregisterCapability(ctx context.Context, params *UnregistrationParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#telemetry_event
	Event(ctx context.Context, params *interface{}) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#textDocument_publishDiagnostics
	PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#window_logMessage
	LogMessage(ctx context.Context, params *LogMessageParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#window_showDocument
	ShowDocument(ctx context.Context, params *ShowDocumentParams) (*ShowDocumentResult, error)
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#window_showMessage
	ShowMessage(ctx context.Context, params *ShowMessageParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#window_showMessageRequest
	ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error)
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#window_workDoneProgress_create
	WorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_applyEdit
	ApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error)
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_codeLens_refresh
	CodeLensRefresh(ctx context.Context) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_configuration
	Configuration(ctx context.Context, params *ParamConfiguration) ([]LSPAny, error)
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_diagnostic_refresh
	DiagnosticRefresh(ctx context.Context) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_foldingRange_refresh
	FoldingRangeRefresh(ctx context.Context) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_inlayHint_refresh
	InlayHintRefresh(ctx context.Context) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_inlineValue_refresh
	InlineValueRefresh(ctx context.Context) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_semanticTokens_refresh
	SemanticTokensRefresh(ctx context.Context) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_textDocumentContent_refresh
	TextDocumentContentRefresh(ctx context.Context, params *TextDocumentContentRefreshParams) error
	// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification#workspace_workspaceFolders
	WorkspaceFolders(ctx context.Context) ([]WorkspaceFolder, error)
}

func clientDispatch(ctx context.Context, client Client, r *jsonrpc2.Request) (bool, any, error) {
	defer recoverHandlerPanic(r.Method)
	switch r.Method {
	case MethodLogTrace:
		var params LogTraceParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.LogTrace(ctx, &params)
		return true, nil, err

	case MethodProgress:
		var params ProgressParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.Progress(ctx, &params)
		return true, nil, err

	case MethodRegisterCapability:
		var params RegistrationParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.RegisterCapability(ctx, &params)
		return true, nil, err

	case MethodUnregisterCapability:
		var params UnregistrationParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.UnregisterCapability(ctx, &params)
		return true, nil, err

	case MethodEvent:
		var params interface{}
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.Event(ctx, &params)
		return true, nil, err

	case MethodPublishDiagnostics:
		var params PublishDiagnosticsParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.PublishDiagnostics(ctx, &params)
		return true, nil, err

	case MethodLogMessage:
		var params LogMessageParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.LogMessage(ctx, &params)
		return true, nil, err

	case MethodShowDocument:
		var params ShowDocumentParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		resp, err := client.ShowDocument(ctx, &params)
		if err != nil {
			return true, nil, err
		}
		return true, resp, nil

	case MethodShowMessage:
		var params ShowMessageParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.ShowMessage(ctx, &params)
		return true, nil, err

	case MethodShowMessageRequest:
		var params ShowMessageRequestParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		resp, err := client.ShowMessageRequest(ctx, &params)
		if err != nil {
			return true, nil, err
		}
		return true, resp, nil

	case MethodWorkDoneProgressCreate:
		var params WorkDoneProgressCreateParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.WorkDoneProgressCreate(ctx, &params)
		return true, nil, err

	case MethodApplyEdit:
		var params ApplyWorkspaceEditParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		resp, err := client.ApplyEdit(ctx, &params)
		if err != nil {
			return true, nil, err
		}
		return true, resp, nil

	case MethodCodeLensRefresh:
		err := client.CodeLensRefresh(ctx)
		return true, nil, err

	case MethodConfiguration:
		var params ParamConfiguration
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		resp, err := client.Configuration(ctx, &params)
		if err != nil {
			return true, nil, err
		}
		return true, resp, nil

	case MethodDiagnosticRefresh:
		err := client.DiagnosticRefresh(ctx)
		return true, nil, err

	case MethodFoldingRangeRefresh:
		err := client.FoldingRangeRefresh(ctx)
		return true, nil, err

	case MethodInlayHintRefresh:
		err := client.InlayHintRefresh(ctx)
		return true, nil, err

	case MethodInlineValueRefresh:
		err := client.InlineValueRefresh(ctx)
		return true, nil, err

	case MethodSemanticTokensRefresh:
		err := client.SemanticTokensRefresh(ctx)
		return true, nil, err

	case MethodTextDocumentContentRefresh:
		var params TextDocumentContentRefreshParams
		if err := UnmarshalJSON(r.Params, &params); err != nil {
			return true, nil, sendParseError(ctx, err)
		}
		err := client.TextDocumentContentRefresh(ctx, &params)
		return true, nil, err

	case MethodWorkspaceFolders:
		resp, err := client.WorkspaceFolders(ctx)
		if err != nil {
			return true, nil, err
		}
		return true, resp, nil

	default:
		return false, nil, nil
	}
}

func (s *clientDispatcher) LogTrace(ctx context.Context, params *LogTraceParams) error {
	return s.sender.Notify(ctx, "$/logTrace", params)
}
func (s *clientDispatcher) Progress(ctx context.Context, params *ProgressParams) error {
	return s.sender.Notify(ctx, "$/progress", params)
}
func (s *clientDispatcher) RegisterCapability(ctx context.Context, params *RegistrationParams) error {
	return s.sender.Call(ctx, "client/registerCapability", params).Await(ctx, nil)
}
func (s *clientDispatcher) UnregisterCapability(ctx context.Context, params *UnregistrationParams) error {
	return s.sender.Call(ctx, "client/unregisterCapability", params).Await(ctx, nil)
}
func (s *clientDispatcher) Event(ctx context.Context, params *interface{}) error {
	return s.sender.Notify(ctx, "telemetry/event", params)
}
func (s *clientDispatcher) PublishDiagnostics(ctx context.Context, params *PublishDiagnosticsParams) error {
	return s.sender.Notify(ctx, "textDocument/publishDiagnostics", params)
}
func (s *clientDispatcher) LogMessage(ctx context.Context, params *LogMessageParams) error {
	return s.sender.Notify(ctx, "window/logMessage", params)
}
func (s *clientDispatcher) ShowDocument(ctx context.Context, params *ShowDocumentParams) (*ShowDocumentResult, error) {
	var result *ShowDocumentResult
	if err := s.sender.Call(ctx, "window/showDocument", params).Await(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) ShowMessage(ctx context.Context, params *ShowMessageParams) error {
	return s.sender.Notify(ctx, "window/showMessage", params)
}
func (s *clientDispatcher) ShowMessageRequest(ctx context.Context, params *ShowMessageRequestParams) (*MessageActionItem, error) {
	var result *MessageActionItem
	if err := s.sender.Call(ctx, "window/showMessageRequest", params).Await(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) WorkDoneProgressCreate(ctx context.Context, params *WorkDoneProgressCreateParams) error {
	return s.sender.Call(ctx, "window/workDoneProgress/create", params).Await(ctx, nil)
}
func (s *clientDispatcher) ApplyEdit(ctx context.Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	var result *ApplyWorkspaceEditResult
	if err := s.sender.Call(ctx, "workspace/applyEdit", params).Await(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) CodeLensRefresh(ctx context.Context) error {
	return s.sender.Call(ctx, "workspace/codeLens/refresh", nil).Await(ctx, nil)
}
func (s *clientDispatcher) Configuration(ctx context.Context, params *ParamConfiguration) ([]LSPAny, error) {
	var result []LSPAny
	if err := s.sender.Call(ctx, "workspace/configuration", params).Await(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) DiagnosticRefresh(ctx context.Context) error {
	return s.sender.Call(ctx, "workspace/diagnostic/refresh", nil).Await(ctx, nil)
}
func (s *clientDispatcher) FoldingRangeRefresh(ctx context.Context) error {
	return s.sender.Call(ctx, "workspace/foldingRange/refresh", nil).Await(ctx, nil)
}
func (s *clientDispatcher) InlayHintRefresh(ctx context.Context) error {
	return s.sender.Call(ctx, "workspace/inlayHint/refresh", nil).Await(ctx, nil)
}
func (s *clientDispatcher) InlineValueRefresh(ctx context.Context) error {
	return s.sender.Call(ctx, "workspace/inlineValue/refresh", nil).Await(ctx, nil)
}
func (s *clientDispatcher) SemanticTokensRefresh(ctx context.Context) error {
	return s.sender.Call(ctx, "workspace/semanticTokens/refresh", nil).Await(ctx, nil)
}
func (s *clientDispatcher) TextDocumentContentRefresh(ctx context.Context, params *TextDocumentContentRefreshParams) error {
	return s.sender.Call(ctx, "workspace/textDocumentContent/refresh", params).Await(ctx, nil)
}
func (s *clientDispatcher) WorkspaceFolders(ctx context.Context) ([]WorkspaceFolder, error) {
	var result []WorkspaceFolder
	if err := s.sender.Call(ctx, "workspace/workspaceFolders", nil).Await(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
