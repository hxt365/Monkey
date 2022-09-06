package evaluator

import (
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GetMyBuiltin("len"),
	"puts":  object.GetMyBuiltin("puts"),
	"first": object.GetMyBuiltin("first"),
	"last":  object.GetMyBuiltin("last"),
	"rest":  object.GetMyBuiltin("rest"),
	"push":  object.GetMyBuiltin("push"),
}
