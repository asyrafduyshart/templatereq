package templatereq

import (
	"testing"
)

var template = "https://www.testweb.com/$TEST"
var template_func = `http://hello.com/$func("hash:$TEST")`
var template_json = `'{"name":"John", "age":$func("hash:$TEST"), "car":null}'`

func TestReplaceMap(t *testing.T) {
	init := replaceByMap(template, map[string]string{
		"TEST": "HELLO",
	})
	expect := "https://www.testweb.com/HELLO"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncName(t *testing.T) {
	init := funcName(`func("hash:test")`)
	expect := "hash"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncErr(t *testing.T) {
	init := funcName(`("hash:test")`)
	expect := `("hash:test")`
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFunctionHash(t *testing.T) {
	init := funcHash("123")
	expect := "1916298011"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestReplaceFuncWithValue(t *testing.T) {
	init := replaceFuncWithValue(`http://hello.com/$func("hash:test")`)
	expect := `http://hello.com/2949673445`
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestReplaceFuncWithValueErr(t *testing.T) {
	init := replaceFuncWithValue(`http://hello.com/$fun("hash:test")`)
	expect := `http://hello.com/$fun("hash:test")`
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncReplaceMap(t *testing.T) {
	change := replaceByMap(template_func, map[string]string{
		"TEST": "test",
	})

	init := replaceFuncWithValue(change)
	expect := "http://hello.com/2949673445"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncReplaceMapJSON(t *testing.T) {
	change := replaceByMap(template_json, map[string]string{
		"TEST": "test",
	})

	init := replaceFuncWithValue(change)
	expect := `{"name":"John", "age":2949673445, "car":null}`

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncSortKey(t *testing.T) {
	funcSortKey(map[string]string{
		"orange":     "1",
		"apple":      "2",
		"mango":      "3",
		"strawberry": "4",
	})
}

func TestFuncUrlEncode(t *testing.T) {
	funcUrlEncode(
		map[string]string{
			"orange":     "1",
			"apple":      "2",
			"mango":      "3",
			"strawberry": "4",
		},
	)
}

func TestFuncMD5(t *testing.T) {
	r := "apple=2&mango=3&orange=1&strawberry=4"
	funcMd5(r)
}

func TestFuncSha256(t *testing.T) {
	r := "NKG2022-11-18 09:01:472022-11-18 09:01:47SecretKey"
	funcSha256(r)
}

func TestFuncSha256ToLowerCase(t *testing.T) {
	r := "NKG2022-11-18 09:01:472022-11-18 09:01:47SecretKey"
	funcSha256ToLowerCase(r)
}

func TestFuncMD5Base64(t *testing.T) {
	r := "apple=2&mango=3&orange=1&strawberry=4"
	funcBase64(r)
}

func TestFuncNormalizeDateWithAdjustment(t *testing.T) {
	// normal
	init := replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z")`)
	expect := "2022-11-07 04:40:39"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// add 5 min
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:add:60*5")`)
	expect = "2022-11-07 04:45:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// add 5 hour
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:add:60*60*5")`)
	expect = "2022-11-07 09:40:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// add 5 day
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:add:60*60*24*5")`)
	expect = "2022-11-12 04:40:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// subtract 5 min
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:subtract:60*5")`)
	expect = "2022-11-07 04:35:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// subtract 5 hour
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:subtract:60*60*5")`)
	expect = "2022-11-06 23:40:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// subtract 5 day
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:subtract:60*60*24*5")`)
	expect = "2022-11-02 04:40:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}
