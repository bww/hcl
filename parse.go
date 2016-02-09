package hcl

import (
	"fmt"

<<<<<<< HEAD
	"github.com/bww/hcl/hcl"
	"github.com/bww/hcl/json"
=======
	"github.com/hashicorp/hcl/hcl/ast"
	hclParser "github.com/hashicorp/hcl/hcl/parser"
	jsonParser "github.com/hashicorp/hcl/json/parser"
>>>>>>> upstream/master
)

// Parse parses the given input and returns the root object.
//
// The input format can be either HCL or JSON.
func Parse(input string) (*ast.File, error) {
	switch lexMode(input) {
	case lexModeHcl:
		return hclParser.Parse([]byte(input))
	case lexModeJson:
		return jsonParser.Parse([]byte(input))
	}

	return nil, fmt.Errorf("unknown config format")
}
