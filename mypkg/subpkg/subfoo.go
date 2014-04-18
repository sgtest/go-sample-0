package subpkg

import "github.com/sgtest/go-sample-0/mypkg"

type Subqux mypkg.Qux

func subfunc() {
	mypkg.Foo()
}
