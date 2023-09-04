package expressions

func expLogicalAnd(tree *ParserT) error {
	/*leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}*/ /*
		convertScalarToBareword(left)

		if left.key != symbols.Bareword {
			return raiseError(tree.expression, left, 0, fmt.Sprintf(
				"left side of %s should be a bareword, instead got %s",
				tree.currentSymbol().key, left.key))
		}

		v, dt, err := tree.getVar(left.value, varAsValue)
		if err != nil {
			if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
				// var doesn't exist and we have strict types disabled so lets create var
				v, dt, err = float64(0), types.Number, nil
			} else {
				return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
		}

		var result interface{}

		switch dt {
		case types.Number, types.Float:
			if right.dt.Primitive != primitives.Number {
				return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
					"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
			}
			result = v.(float64) + right.dt.Value.(float64)

		case types.Integer:
			if right.dt.Primitive != primitives.Number {
				return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
					"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
			}
			result = float64(v.(int)) + right.dt.Value.(float64)

		case types.Boolean:
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s", tree.currentSymbol().key, dt))

		case types.Null:
			switch right.dt.Primitive {
			case primitives.String:
				result = right.dt.Value.(string)
			case primitives.Number:
				result = right.dt.Value.(float64)
			default:
				return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
					"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
			}

		default:
			if right.dt.Primitive != primitives.String {
				return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
					"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
			}
			result = v.(string) + right.dt.Value.(string)
		}

		err = tree.setVar(left.value, result, right.dt.DataType())
		if err != nil {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}

		return tree.foldAst(&astNodeT{
			key: symbols.Calculated,
			pos: tree.ast[tree.astPos].pos,
			dt: &primitives.DataType{
				Primitive: primitives.Null,
				Value:     nil,
			},
		})*/
	return nil
}
