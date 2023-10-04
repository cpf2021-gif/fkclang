package evaluator

import "fkclang/object"

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Map:
				return &object.Integer{Value: int64(len(arg.Pairs))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"front": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `front` must be Array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"back": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `front` must be Array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `rest` must be Array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `push` must be Array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if args[1].Type() != object.ArrayObj {
				newElements := make([]object.Object, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &object.Array{Elements: newElements}
			} else {
				arr1 := args[1].(*object.Array)
				newElements := make([]object.Object, length+len(arr1.Elements))
				copy(newElements, arr.Elements)
				copy(newElements[length:], arr1.Elements)

				return &object.Array{Elements: newElements}
			}
		},
	},
	"unshift": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `unshift` must be Array, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if args[1].Type() != object.ArrayObj {
				newElements := make([]object.Object, length+1)
				copy(newElements[1:], arr.Elements)
				newElements[0] = args[1]

				return &object.Array{Elements: newElements}
			} else {
				arr1 := args[1].(*object.Array)
				newElements := make([]object.Object, length+len(arr1.Elements))
				copy(newElements[len(arr1.Elements):], arr.Elements)
				copy(newElements, arr1.Elements)

				return &object.Array{Elements: newElements}
			}
		},
	},
	"isInteger": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() == object.IntegerObj {
				return TRUE
			}
			if args[0].Type() == object.ArrayObj {
				arr := args[0].(*object.Array)
				for _, v := range arr.Elements {
					if v.Type() != object.IntegerObj {
						return FALSE
					}
				}
				return TRUE
			}
			return FALSE
		},
	},
	"subArray": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3",
					len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("argument to `subArray` must be Array, got %s",
					args[0].Type())
			}
			if args[1].Type() != object.IntegerObj {
				return newError("argument to `subArray` must be Integer, got %s",
					args[1].Type())
			}
			if args[2].Type() != object.IntegerObj {
				return newError("argument to `subArray` must be Integer, got %s",
					args[2].Type())
			}
			arr := args[0].(*object.Array)
			start := args[1].(*object.Integer).Value
			end := args[2].(*object.Integer).Value
			if start < 0 || end < 0 || start > end || start > int64(len(arr.Elements)) {
				return newError("argument to `subArray` is invalid")
			}
			newElements := make([]object.Object, end-start)
			copy(newElements, arr.Elements[start:end])
			return &object.Array{Elements: newElements}
		},
	},
	"newError": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.StringObj {
				return newError("argument to `newError` must be String, got %s",
					args[0].Type())
			}
			return newError(args[0].Inspect())
		},
	},
}
