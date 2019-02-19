// +build !freebsd

package mint

import (
	"os"

	// On "freebsd/FreeBSD-10.4-STABLE" OS image,
	// Go installed by `pkg install` might NOT have `syscall.Mprotect`
	// causing such error: "bou.ke/monkey/replace_unix.go:13:10: undefined: syscall.Mprotect".
	// See https://www.freebsd.org/cgi/man.cgi?sektion=2&query=mprotect
	// TODO: Fix the image for https://github.com/otiai10/gosseract/blob/master/test/runtimes/freebsd.Vagrantfile#L4
	"bou.ke/monkey"
)

// Exit ...
func (testee *Testee) Exit(expectedCode int) Result {

	fun, ok := testee.actual.(func())
	if !ok {
		panic("mint error: Exit only can be called for func type value")
	}

	var actualCode int
	patch := monkey.Patch(os.Exit, func(code int) {
		actualCode = code
	})
	fun()
	patch.Unpatch()

	testee.actual = actualCode
	if judge(actualCode, expectedCode, testee.not, testee.deeply) {
		return testee.result
	}
	testee.expected = expectedCode
	return testee.failed(failExitCode)
}
