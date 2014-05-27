package mint_test

import "testing"
import "github.com/otiai10/mint"

func TestMint_ToBe(t *testing.T) {
	mint.Expect(t, 1).ToBe(1)
}
func TestMint_ToBe_Fail(t *testing.T) {
	r := mint.Expect(t, 2).Dry().ToBe(1).Result
	// assert mint by using mint
	mint.Expect(t, r.OK).ToBe(false)
}

func TestMint_TypeOf(t *testing.T) {
	mint.Expect(t, "foo").TypeOf("string")
}
func TestMint_TypeOf_Fail(t *testing.T) {
	r := mint.Expect(t, "foo").Dry().TypeOf("int").Result
	// assert mint by using mint
	mint.Expect(t, r.OK).ToBe(false)
}

func TestMint_Not(t *testing.T) {
	mint.Expect(t, 100).Not().ToBe(200)
	mint.Expect(t, "foo").Not().TypeOf("int")
}
func TestMint_Not_Fail(t *testing.T) {
	r := mint.Expect(t, "foo").Dry().Not().TypeOf("string").Result
	// assert mint by using mint
	mint.Expect(t, r.OK).Not().ToBe(true)
}
