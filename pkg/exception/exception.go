package exception

import (
	"fmt"
	"runtime/debug"
)

func Must(e any) {
	if e != nil {
		panic(e)
	}
}

func TryCatch(try func(), catch func(e any)) {
	defer func() {
		if e := recover(); e != nil {
			catch(e)
		}
	}()
	try()
}

func Try(try func()) any {
	var ret any = nil
	TryCatch(try, func(e any) {
		fmt.Printf("%s\n", debug.Stack())
		ret = e
	})
	return ret
}

func toError(e any) error {
	if e == nil {
		return nil
	}
	if err, ok := e.(error); ok {
		return err
	}
	return fmt.Errorf("%v", e)
}

func TryWithError(try func()) error {
	return toError(Try(try))
}
