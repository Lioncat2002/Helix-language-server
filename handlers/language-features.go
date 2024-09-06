package handlers

import (
	"github.com/Lioncat2002/helix-lsp/mappers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	for word, mapping := range mappers.Mappings {
		mapCopy := mapping
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word,
			Detail:     &mapCopy,
			InsertText: &mapCopy,
		})
	}

	return completionItems,nil
}
