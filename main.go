package main

import (
	"fmt"
	"reflect"
	"runtime"

	//"github.com/d4l3k/go-pry/pry"

	//"github.com/itisnotdone/gostudy/fundamentals"

	//"github.com/itisnotdone/gostudy/tutorials/welcome"
	//"github.com/itisnotdone/gostudy/tutorials/basics"
	//"github.com/itisnotdone/gostudy/tutorials/flowcontrol"
	//"github.com/itisnotdone/gostudy/tutorials/moretypes"
	//"github.com/itisnotdone/gostudy/tutorials/methods"
	"github.com/itisnotdone/gostudy/others"
)

//import "github.com/d4l3k/go-pry/pry"

func main() {
	println()
	println()
	println()

	//runFunction(others.Func01)
	//runFunction(others.FuncSender)

	//runFunction(others.DoMarshalJSON)
	//runFunction(others.DoUnmarshalJSON)
	runFunction(others.DoMarshalYAML)

	//runFunction(methods.Methods01)
	//runFunction(methods.Methods02)
	//runFunction(methods.Methods03)
	//runFunction(methods.Methods04)
	//runFunction(methods.Methods05)
	//runFunction(methods.Methods06)
	//runFunction(methods.Methods07)
	//runFunction(methods.Methods08)
	//runFunction(methods.Methods09)

	//runFunction(moretypes.Moretypes01)
	//runFunction(moretypes.Moretypes02)
	//runFunction(moretypes.Moretypes03)
	//runFunction(moretypes.Moretypes04)
	//runFunction(moretypes.Moretypes05)
	//runFunction(moretypes.Moretypes06)
	//runFunction(moretypes.Moretypes07)
	//runFunction(moretypes.Moretypes08)
	//runFunction(moretypes.Moretypes09)
	//runFunction(moretypes.Moretypes10)
	//runFunction(moretypes.Moretypes11)
	//runFunction(moretypes.Moretypes12)
	//runFunction(moretypes.Moretypes13)
	//runFunction(moretypes.Moretypes14)
	//runFunction(moretypes.Moretypes15)
	//runFunction(moretypes.Moretypes16)
	//runFunction(moretypes.Moretypes17)
	//runFunction(moretypes.Moretypes19)
	//runFunction(moretypes.Moretypes20)
	//runFunction(moretypes.Moretypes22)
	//runFunction(moretypes.Moretypes24)
	//runFunction(moretypes.Moretypes25)

	//runFunction(flowcontrol.Flowcontrol01)
	//runFunction(flowcontrol.Flowcontrol02)
	//runFunction(flowcontrol.Flowcontrol05)
	//runFunction(flowcontrol.Flowcontrol06)
	//runFunction(flowcontrol.Flowcontrol07)
	//runFunction(flowcontrol.Flowcontrol08)
	//runFunction(flowcontrol.Flowcontrol09)
	//runFunction(flowcontrol.Flowcontrol10)
	//runFunction(flowcontrol.Flowcontrol11)
	//runFunction(flowcontrol.Flowcontrol12)
	//runFunction(flowcontrol.Flowcontrol13)

	//runFunction(basics.Basics1)
	//runFunction(basics.Basics2)
	//runFunction(basics.Basics3)
	//runFunction(basics.Basics4)
	//runFunction(basics.Basics6)
	//runFunction(basics.Basics7)
	//runFunction(basics.Basics8)
	//runFunction(basics.Basics9)
	//runFunction(basics.Basics10)
	//runFunction(basics.Basics9)
	//runFunction(basics.Basics11)
	//runFunction(basics.Basics12)
	//runFunction(basics.Basics13)
	//runFunction(basics.Basics14)
	//runFunction(basics.Basics15)
	//runFunction(basics.Basics16)

	//runFunction(welcome.Welcome1)
	//runFunction(welcome.Welcome4)

	//runFunction(fundamentals.Hello)
	//runFunction(fundamentals.Func111)
	//runFunction(fundamentals.Func222)

	println()
	println()
	println()
}

func runFunction(fn func()) {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
	fn()
	fmt.Println()
}
