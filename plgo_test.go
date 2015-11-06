package plgo_test

import (
	"math"
	"math/cmplx"
	"reflect"
	"testing"
	"git.dev.whs/talby/plgo"
)

var pl = plgo.New()

//var fd int
//
//func init() {
//	// establish a channel to communicate via something other than
//	// arguments and return values
//	fds, _ := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM,
//			syscall.AF_UNSPEC)
//	fd = fds[0]
//	pl.Bind(nil, fmt.Sprintf(`open $fh, '<&=', %d`, fds[1]))
//}

func TestNil(t *testing.T) {
	err := pl.Bind(nil, `1`)
	if err != nil {
		t.Errorf("perl error unexpected: %s", err.Error())
	}
	err = pl.Bind(nil, `1 = 2`)
	if err == nil {
		t.Errorf("perl error expected")
	}
}

func TestBool(t *testing.T) {
	var id func(bool) bool
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want bool) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v bool) => %v", want, have)
		}
	}
	ok(true)
	ok(false)

	is := func(expr string, want bool) {
		var have bool
		pl.Bind(&have, expr)
		if want != have {
			t.Errorf("is(`%s`) want %v have %v", expr, want, have)
		}
	}
	is(`undef`, false)
	is(`1 == 1`, true)
	is(`1 == 0`, false)
	is(`''`, false)
	is(`'0'`, false)
	is(`'1'`, true)
	is(`'2'`, true)
	is(`'-1'`, true)
	is(`'a string'`, true)
}

func TestInt(t *testing.T) {
	var id func(int) int
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want int) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v int) => %v", want, have)
		}
	}
	// the capacity of an "int" is system dependent
	ok(-1)
	ok(0)
	ok(1)
}

func TestInt8(t *testing.T) {
	var id func(int8) int8
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want int8) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v int8) => %v", want, have)
		}
	}
	ok(-128)
	ok(-1)
	ok(0)
	ok(1)
	ok(127)
}

func TestInt16(t *testing.T) {
	var id func(int16) int16
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want int16) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v int16) => %v", want, have)
		}
	}
	ok(-32768)
	ok(-1)
	ok(0)
	ok(1)
	ok(32767)
}

func TestInt32(t *testing.T) {
	var id func(int32) int32
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want int32) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v int32) => %v", want, have)
		}
	}
	ok(-2147483648)
	ok(-1)
	ok(0)
	ok(1)
	ok(2147483647)
}

func TestInt64(t *testing.T) {
	var id func(int64) int64
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want int64) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v int64) => %v", want, have)
		}
	}
	ok(-9223372036854775808)
	ok(-1)
	ok(0)
	ok(1)
	ok(9223372036854775807)
}

func TestUint(t *testing.T) {
	var id func(uint) uint
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want uint) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v uint) => %v", want, have)
		}
	}
	// the capacity of a "uint" is system dependent
	ok(0)
	ok(1)
}

func TestUint8(t *testing.T) {
	var id func(uint8) uint8
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want uint8) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v uint8) => %v", want, have)
		}
	}
	ok(0)
	ok(1)
	ok(255)
}

func TestUint16(t *testing.T) {
	var id func(uint16) uint16
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want uint16) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v uint16) => %v", want, have)
		}
	}
	ok(0)
	ok(1)
	ok(65535)
}

func TestUint32(t *testing.T) {
	var id func(uint32) uint32
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want uint32) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v uint32) => %v", want, have)
		}
	}
	ok(0)
	ok(1)
	ok(4294967295)
}

func TestUint64(t *testing.T) {
	var id func(uint64) uint64
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want uint64) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v uint64) => %v", want, have)
		}
	}
	ok(0)
	ok(1)
	ok(18446744073709551615)
}

func TestUintptr(t *testing.T) {
	var id func(uintptr) uintptr
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want uintptr) {
		have := id(want)
		if want != have {
			t.Errorf("id(%v uintptr) => %v", want, have)
		}
	}
	ok(uintptr(0))
	ok(uintptr(1))
	ok(^uintptr(0))
}

// matching floating point numbers exactly is always a little sketchy.
func TestFloat32(t *testing.T) {
	var id func(float32) float32
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want float32) {
		have := id(want)
		if want != have &&
			!(math.IsNaN(float64(want)) && math.IsNaN(float64(have))) {
			t.Errorf("id(%v float32) => %v", want, have)
		}
	}
	ok(float32(math.Inf(-1)))
	ok(-math.MaxFloat32)
	ok(-1.0)
	ok(-math.SmallestNonzeroFloat32)
	ok(0.0)
	ok(math.SmallestNonzeroFloat32)
	ok(1.0)
	ok(math.MaxFloat32)
	ok(float32(math.Inf(1)))
	ok(float32(math.NaN()))
}

func TestFloat64(t *testing.T) {
	var id func(float64) float64
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want float64) {
		have := id(want)
		if want != have &&
			!(math.IsNaN(want) && math.IsNaN(have)) {
			t.Errorf("id(%v float64) => %v", want, have)
		}
	}
	ok(math.Inf(-1))
	ok(-math.MaxFloat64)
	ok(-1.0)
	ok(-math.SmallestNonzeroFloat64)
	ok(0.0)
	ok(math.SmallestNonzeroFloat64)
	ok(1.0)
	ok(math.MaxFloat64)
	ok(math.Inf(1))
	ok(math.NaN())
}

func TestComplex64(t *testing.T) {
	var id func(complex64) complex64
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want complex64) {
		have := id(want)
		if want == have ||
			(cmplx.IsNaN(complex128(want)) && cmplx.IsNaN(complex128(have))) ||
			(cmplx.IsInf(complex128(want)) && cmplx.IsInf(complex128(have))) {
			return
		}
		t.Errorf("id(%v complex64) => %v", want, have)
	}
	vals := []float32{
		float32(math.Inf(-1)),
		-math.MaxFloat32,
		-1.0,
		-math.SmallestNonzeroFloat32,
		0.0,
		math.SmallestNonzeroFloat32,
		1.0,
		math.MaxFloat32,
		float32(math.Inf(1)),
		float32(math.NaN()),
	}
	for _, re := range vals {
		for _, im := range vals {
			ok(complex(re, im))
		}
	}
	ok(complex64(-cmplx.Inf()))
	ok(complex64(cmplx.Inf()))
	ok(complex64(cmplx.NaN()))
}

func TestComplex128(t *testing.T) {
	var id func(complex128) complex128
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want complex128) {
		have := id(want)
		if want == have ||
			(cmplx.IsNaN(want) && cmplx.IsNaN(have)) ||
			(cmplx.IsInf(want) && cmplx.IsInf(have)) {
			return
		}
		t.Errorf("id(%v complex128) => %v", want, have)
	}
	vals := []float64{
		math.Inf(-1),
		-math.MaxFloat64,
		-1.0,
		-math.SmallestNonzeroFloat64,
		0.0,
		math.SmallestNonzeroFloat64,
		1.0,
		math.MaxFloat64,
		math.Inf(1),
		math.NaN(),
	}
	for _, re := range vals {
		for _, im := range vals {
			ok(complex(re, im))
		}
	}
	ok(-cmplx.Inf())
	ok(cmplx.Inf())
	ok(cmplx.NaN())
}

func TestFunc(t *testing.T) {
	var id func(func(int) int) func(int) int
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want func(int) int) {
		have := id(want)
		if have(18) != want(18) {
			t.Errorf("id(%v func() int) => %v", want, have)
		}
	}
	ok(func(v int) int {
		return v + 54321
	})
}

func TestMap(t *testing.T) {
	var id func(map[string]int) map[string]int
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want map[string]int) {
		have := id(want)
		if !reflect.DeepEqual(have, want) {
			t.Errorf("id(%v map[string]int) => %v", want, have)
		}
	}
	ok(map[string]int{
		"fun": 12,
		"jam": 8,
	})
}

func TestSV(t *testing.T) {
	ok := func(mk, ck string) {
		var val *plgo.SV
		var chk func(*plgo.SV) bool
		pl.Bind(&val, mk)
		pl.Bind(&chk, "sub {"+ck+"}")
		if !chk(val) {
			t.Errorf("sub { %s }->(%s) fail", ck, mk)
		}
	}
	// IV
	ok(`321`, `$_[0] == 321`)
	// NV
	ok(`32.1`, `$_[0] == 32.1`)
	// RV
	ok(`$val = 15; \$val`, `$_[0] == \$val`)
	// PV
	ok(`'a string'`, `$_[0] eq 'a string'`)
	// PVMG
	ok(`$^T`, `$_[0] == $^T`)
	// REGEXP
	ok(`qr/fun/`, `"funny" =~ $_[0] and "jello" !~ $_[0]`)
	// PVAV
	ok(`[ 1 .. 3 ]`, `"@{$_[0]}" eq '1 2 3'`)
	// PVHV
	ok(`{ a => 'b' }`, `"@{[ %{$_[0]} ]}" eq 'a b'`)
	// PVCV
	ok(`sub { 543 }`, `$_[0]->() == 543`)
	//	ok(`undef`, `not defined $_[0]`)
	//	//ok(nil)
}

func TestSlice(t *testing.T) {
	var id func([]int) []int
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want []int) {
		have := id(want)
		if !reflect.DeepEqual(have, want) {
			t.Errorf("id(%v []int) => %v", want, have)
		}
	}
	ok([]int{})
	ok([]int{1, 2, 3})
}

func TestString(t *testing.T) {
	var id func(string) string
	pl.Bind(&id, `sub { $_[0] }`)
	ok := func(want string) {
		have := id(want)
		if !reflect.DeepEqual(have, want) {
			t.Errorf("id(%v string) => %v", want, have)
		}
	}
	ok("")
	ok("a string")
}

func TestMulti(t *testing.T) {
	var fn func(int, int) (int, int, int, int, int)
	pl.Bind(&fn, `sub {
		use strict;
		use warnings;
		# the xgcd algorithm
		my($u, $v) = @_;
		my(@t, @u, @v, $q);
		@u = (1, 0, $u);
		@v = (0, 1, $v);
		while($v[2] != 0) {
			$q = int($u[2] / $v[2]);
			$t[0] = $u[0] - ($v[0] * $q);
			$u[0] = $v[0];
			$v[0] = $t[0];
			$t[1] = $u[1] - ($v[1] * $q);
			$u[1] = $v[1];
			$v[1] = $t[1];
			$t[2] = $u[2] - ($v[2] * $q);
			$u[2] = $v[2];
			$v[2] = $t[2];
		}
		return($u[0], $u, 0 - $u[1], $v, $u[2]) if $u[0] > 0;
		return($u[1], $v, 0 - $u[0], $u, $u[2]);
	}`)
	a, b, c, d, e := fn(12345, 54321)
	if a != 3617 || b != 12345 || c != 822 || d != 54321 || e != 3 {
		t.Errorf("%d*%d-%d*%d ?= %d", a, b, c, d, e)
	}
}