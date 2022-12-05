package main

import (
	"lesson_go/definition/var_parameter"
	"lesson_go/statement/panic_and_recover"
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

	panic_and_recover.PanicAneRecover()

}

func main() {
	//definition()
	statement()
}
