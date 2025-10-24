package expressions

import "fmt"

func (tree *ParserT) Dump() any {
	var (
		dump  = make(map[string]any)
		nodes = make([]any, len(tree.ast))
	)

	for i := range tree.ast {
		node := make(map[string]any)
		node["key"] = tree.ast[i].key.String()
		node["value"] = tree.ast[i].Value()
		node["pos"] = tree.ast[i].pos
		if tree.ast[i].dt != nil {
			dt, err := tree.ast[i].dt.GetValue()
			if err == nil && dt != nil {
				node["dt.prim"] = dt.Primitive.String()
				node["dt.murex"] = dt.DataType
				node["dt.value"] = dt.Value
			} else {
				node["dt"] = fmt.Sprintf("%v", err)
			}
		} else {
			node["dt"] = "unset"
		}

		nodes[i] = node
	}

	dump["ast"] = nodes
	dump["charPos"] = tree.charPos
	dump["charOffset"] = tree.charOffset
	dump["astPos"] = tree.astPos
	dump["expression"] = string(tree.expression)
	dump["statement"] = map[string]any{
		"command":    tree.statement.String(),
		"parameters": tree.statement.Parameters(),
	}

	return dump
}
