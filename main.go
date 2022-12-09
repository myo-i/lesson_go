package main

import (
	"lesson_go/definition/var_parameter"
	"lesson_go/goroutine"
	"lesson_go/pointer"
	"lesson_go/statement/exercises"
	"lesson_go/structs"
)

func definition() {
	//hello.Hello()
	//imports.Imports()
	//variable.Variable()
	//_const.Const()
	//numeric.Numeric()
	//boolean.Boolean()
	//cast.Cast()
	//array.Array()
	//slice.Slice()
	//slice_make_cap.MakeCap()
	//_map.Map()
	//_byte.Byte()
	//_func.Func()
	//closure.Closure()
	var_parameter.Parameter()
}

func statement() {
	//_if.If()
	//_for.For()
	//_range.Range()
	//_switch.Switch()
	//_defer.Defer()
	//log.Logging()
	//error.Error()
	//panic_and_recover.PanicAneRecover()
	exercises.Exercise()
}

func pointer() {
	//_pointer.Pointer()
	//_pointer.NewAndMake()
	_pointer.Struct()
}

func structs() {
	//_structs.PointerValue()
	//_structs.NonStruct()
	//_structs.Interface()
	//_structs.TypeAssertion()
	//_structs.Stringer()
	//_structs.CustomError()
	_structs.Exercise()
}

func goroutine() {
	_goroutine.Goroutine()
}

func main() {
	//definition()
	//statement()
	//pointer()
	//structs()
	goroutine()
}
