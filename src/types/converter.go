package types

import "github.com/syucream/spar/src/types"

type Converter interface {
	Convert(statements *types.DDStatements) (string, error)
}
