package lisp

import(
	"fmt"
)

func Eval(ast *Node) int {
	var op []byte 
	var left, right, result int
	if ast == nil {
		panic("invalid ast")
	}
	DEBUG("left: ", ast.left.root.Value)
	op = ast.root.Value
	if ast.root.Type == OP {
		if ast.left != nil {
			if ast.left.root.Type == NUMBER {
				left = toInt(ast.left.root.Value)
			} else if ast.left.root.Type == OP {
				left = Eval(ast.left)
			} else {
				panic("not supported token type")
			}
		}
		if ast.right != nil {
			if ast.right.root.Type == NUMBER {
				right = toInt(ast.right.root.Value)
			} else if ast.right.root.Type == OP {
				right = Eval(ast.right)
			} else {
				panic("not supported token type")
			}
		}

		if var_i == 0 {
			var_i = '0'
		}
		DEBUG("INFO " ,op, left, right)
		result = eval(string(op), left, right)
		return result
	}
	return toInt(ast.root.Value)
}

func eval(op string, left, right int) int{

	fmt.Println("eval:", op, left, right)

	switch{
	case op == "+":
		return left + right
	case op == "-":
		return left - right
	case op == "*":
		return left * right
	case op == "/":
		return left / right
	}

	panic("not supported op " + string(op))
}

func toInt(val []byte) int{
	result := 0
	for _, v := range val{
		factor := 10
		result = result * factor + int(v - '0')
	}
	DEBUG("before: ", val)
	DEBUG("after: ", int(result) - int('0'))
	return int(result)
}