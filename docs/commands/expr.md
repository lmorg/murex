# Expressions (`expr`)

> Expressions: mathematical, string comparisons, logical operators

## Description

`expr` is the underlying builtin which handles all expression parsing and
evaluation in Murex.

Idiomatic Murex would be to write expressions without explicitly calling the
underlying builtin:

```
# idiomatic expressions
1 + 2

# non-idiomatic expressions
expr 1 + 2
```

Though you can invoke them via `expr` if needed, please bare in mind that
expressions have special parsing rules to make them more ergonomic. So if you
write an expression as a command (ie prefixed with `expr`) then it will be
parsed as a statement. This means more complex expressions might parse in
unexpected ways and thus fail. You can still raise a bug if that does happens.

A full list of operators is available in the [Operators And Tokens](/docs/user-guide/operators-and-tokens.md)
document.

## Usage

```
expression -> <stdout>

statement (expression)

expr expression -> <stdout>
```

## Examples

### Basic Expressions

```
» 3 * (3 + 1)
12
```

### Statements with inlined expressions

Any parameter surrounded by parenthesis is first evaluated as an [expression](/docs/parser/expr-inlined.md),
then as a [string](/docs/parser/brace-quote-func.md)".

```
» out (3 * 2)
6
```

### Functions

Expressions also support running commands as [C-style functions](/docs/parser/c-style-fun.md), for example:

```
» 5 * out(5)
25

» datetime(--in {now} --out {unix}) / 60
28339115.783333335

» $file_contents = open(example_file.txt)
```

Please note that currently the only functions supported are ones who's names
are comprised entirely of alpha, numeric, underscore and/or exclamation marks.

### Arrays

```
» %[apples oranges grapes]
[
    "apples",
    "oranges",
    "grapes"
]
```

([read more](/docs/parser/create-array.md))

### Objects

Sometimes known as dictionaries or maps:

```
» %{ Age: { Tom: 20, Dick: 30, Sally: 40 } }
{
    "Age": {
        "Dick": 30,
        "Sally": 40,
        "Tom": 20
    }
}
```

([read more](/docs/parser/create-object.md))

## Detail

### Order of Operations

The order of operations follows the same rules as the C programming language,
which itself is an extension of the order of operations in mathematics, often
referred to as PEMDAS or MODMAS ([read more](https://en.wikipedia.org/wiki/Order_of_operations)).

The [Wikipedia article](https://en.wikipedia.org/wiki/Order_of_operations#Programming_languages)
summarises that order succinctly however the detailed specification is defined
by its implementation, as seen in the code below:

```go
package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/consts"
)

func (tree *ParserT) executeExpr() (*primitives.DataType, error) {
	err := tree.validateExpression()
	if err != nil {
		return nil, err
	}

	for i := range orderOfOperations {
		err = executeExpression(tree, orderOfOperations[i])
		if err != nil {
			return nil, err
		}
	}

	if len(tree.ast) > 1 {
		return nil, fmt.Errorf(
			"expression failed to execute correctly (AST results > 1).\n%s",
			consts.IssueTrackerURL)
	}

	return tree.ast[0].dt, nil
}

// To allow for extendability and developer expectations, the order of operations
// will follow what is defined by (for example) C, as outlined in the following:
// https://en.wikipedia.org/wiki/Order_of_operations#Programming_languages
// Not all operations will be available in murex and some are likely to be added
// in future versions of this package.
//
// Please also note that the slice below is just defining the groupings. Each
// operator within the _same_ group will then be processed from left to right.
// Read the `executeExpression` function further down this source file to view
// every supported operator
var orderOfOperations = []symbols.Exp{
	// 01. Function call, scope, array/member access
	// 02. (most) unary operators, sizeof and type casts (right to left)
	// 03. Multiplication, division, modulo
	symbols.Multiply,

	// 04. Addition and subtraction
	symbols.Add,

	// 04.1 Merge
	symbols.Merge,

	// 05. Bitwise shift left and right
	// 06. Comparisons: less-than and greater-than
	symbols.GreaterThan,

	// 07. Comparisons: equal and not equal
	symbols.EqualTo,

	// 08. Bitwise AND
	// 09. Bitwise exclusive OR (XOR)
	// 10. Bitwise inclusive (normal) OR
	// 11. Logical AND
	symbols.LogicalAnd,

	// 12. Logical OR
	symbols.LogicalOr,

	// 13. Conditional expression (ternary)
	symbols.Elvis,

	// 14. Assignment operators (right to left)
	symbols.Assign,

	// 15. Comma operator
}

func executeExpression(tree *ParserT, order symbols.Exp) (err error) {
	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		node := tree.ast[tree.astPos]

		if node.key < order {
			continue
		}

		switch node.key {

		// 15. Comma operator
		// 14. Assignment operators (right to left)
		case symbols.Assign:
			err = expAssign(tree, true)
		case symbols.AssignUpdate:
			err = expAssign(tree, false)
		case symbols.AssignAndAdd:
			//err = expAssignAdd(tree)
			err = expAssignAndOperate(tree, _assAdd)
		case symbols.AssignAndSubtract:
			err = expAssignAndOperate(tree, _assSub)
		case symbols.AssignAndDivide:
			err = expAssignAndOperate(tree, _assDiv)
		case symbols.AssignAndMultiply:
			err = expAssignAndOperate(tree, _assMulti)
		case symbols.AssignOrMerge:
			err = expAssignMerge(tree)

		// 13. Conditional expression (ternary)
		case symbols.NullCoalescing:
			err = expNullCoalescing(tree)
		case symbols.Elvis:
			err = expElvis(tree)

		// 12. Logical OR
		case symbols.LogicalOr:
			err = expLogicalOr(tree)

		// 11. Logical AND
		case symbols.LogicalAnd:
			err = expLogicalAnd(tree)

		// 10. Bitwise inclusive (normal) OR
		// 09. Bitwise exclusive OR (XOR)
		// 08. Bitwise AND
		// 07. Comparisons: equal and not equal
		case symbols.EqualTo:
			err = expEqualTo(tree)
		case symbols.NotEqualTo:
			err = expNotEqualTo(tree)
		case symbols.Like:
			err = expLike(tree, true)
		case symbols.NotLike:
			err = expLike(tree, false)
		case symbols.Regexp:
			err = expRegexp(tree, true)
		case symbols.NotRegexp:
			err = expRegexp(tree, false)

		// 06. Comparisons: less-than and greater-than
		case symbols.GreaterThan:
			err = expGtLt(tree, _gtF, _gtS)
		case symbols.GreaterThanOrEqual:
			err = expGtLt(tree, _gtEqF, _gtEqS)
		case symbols.LessThan:
			err = expGtLt(tree, _ltF, _ltS)
		case symbols.LessThanOrEqual:
			err = expGtLt(tree, _ltEqF, _ltEqS)

		// 05. Bitwise shift left and right

		// 04.1 Merge
		case symbols.Merge:
			err = expMerge(tree)

			// 04. Addition and subtraction
		case symbols.PlusPlus:
			err = expPlusPlus(tree, 1)
		case symbols.MinusMinus:
			err = expPlusPlus(tree, -1)
		case symbols.Add:
			err = expAdd(tree)
		case symbols.Subtract:
			err = expSubtract(tree)

		// 03. Multiplication, division, modulo
		case symbols.Multiply:
			err = expMultiply(tree)
		case symbols.Divide:
			err = expDivide(tree)

		// 02. (most) unary operators, sizeof and type casts (right to left)
		// 01. Function call, scope, array/member access

		default:
			err = raiseError(tree.expression, node, 0, fmt.Sprintf(
				"no code written to handle symbol (%s)",
				consts.IssueTrackerURL))
		}

		if err != nil {
			return err
		}

		tree.astPos = 0
	}

	return nil
}
```

## See Also

* [( expression )](../parser/expr-inlined.md):
  Inline expressions
* [Open File (`open`)](../commands/open.md):
  Open a file with a preferred handler
* [Output String (`out`)](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [Strict Types In Expressions](../user-guide/strict-types.md):
  Expressions can auto-convert types or strictly honour data types
* [`%[]` Array Builder](../parser/create-array.md):
  Quickly generate arrays
* [`%{}` Object Builder](../parser/create-object.md):
  Quickly generate objects (dictionaries / maps)
* [`*=` Multiply By Operator](../parser/multiply-by.md):
  Multiplies a variable by the right hand value (expression)
* [`*` Multiplication Operator](../parser/multiplication.md):
  Multiplies one numeric value with another (expression)
* [`+=` Add With Operator](../parser/add-with.md):
  Adds the right hand value to a variable (expression)
* [`+` Addition Operator](../parser/addition.md):
  Adds two numeric values together (expression)
* [`-=` Subtract By Operator](../parser/subtract-by.md):
  Subtracts a variable by the right hand value (expression)
* [`-` Subtraction Operator](../parser/subtraction.md):
  Subtracts one numeric value from another (expression)
* [`/=` Divide By Operator](../parser/divide-by.md):
  Divides a variable by the right hand value (expression)
* [`/` Division Operator](../parser/division.md):
  Divides one numeric value from another (expression)
* [`<~` Assign Or Merge](../parser/assign-or-merge.md):
  Merges the right hand value to a variable on the left hand side (expression)
* [`?:` Elvis Operator](../parser/elvis.md):
  Returns the right operand if the left operand is falsy (expression)
* [`??` Null Coalescing Operator](../parser/null-coalescing.md):
  Returns the right operand if the left operand is empty / undefined (expression)

<hr/>

This document was generated from [builtins/core/expressions/expressions_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/expressions/expressions_doc.yaml).