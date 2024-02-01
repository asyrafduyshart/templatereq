package templatereq

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
)

var (
	template      = "https://www.testweb.com/$TEST"
	template_func = `http://hello.com/$func("hash:$TEST")`
	template_json = `'{"name":"John", "age":$func("hash:$TEST"), "car":null}'`
)

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
	init := funcName(`$func("hash:test")`)
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

func TestFuncDESCBC(t *testing.T) {
	r := "stringToEncrypt:keystring:ivstring"
	funcDESCBC(r)
}

func TestReplaceFuncDESCBC(t *testing.T) {
	a := replaceFuncWithValue(`$func("des-cbc:method=GetAllBetDetailsForTransactionID&Key=6E1556ABD55F44ECA802C512DBFAA0AE&Time=20231201160640&TransactionID=1765712:param:g9G16nTs:param:g9G16nTs")`)
	b := replaceFuncWithValue(`$func("chain:prepend::method=GetAllBetDetailsForTransactionID&Key=6E1556ABD55F44ECA802C512DBFAA0AE&Time=20231201160640&TransactionID=1765712>>append::GgaIMaiNNtg>>append::20231201160640>>append::g9G16nTs>>encrypt::md5")`)

	fmt.Printf("\na>>>%v\n", a)
	fmt.Printf("\nb>>>%v\n", b)
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

	// normal with tz
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:format:YYYY-MM-DDThh:mm:ssZ")`)
	expect = "2022-11-07T04:40:39Z"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// nano
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39.99999324Z:format:YYYY-MM-DD hh:mm:ss.nnnn")`)
	expect = "2022-11-07 04:40:39.9999"

	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// nano with tz
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39.99999324Z:format:YYYY-MM-DDThh:mm:ss.nnnnZ")`)
	expect = "2022-11-07T04:40:39.9999Z"

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

func TestFuncDateTimeFormat(t *testing.T) {
	replaceFuncWithValue(`$func("chain:dateNow::Standard>>append:::format:RFC3339>>encrypt::dateTimeFormat")`)
}

func TestFuncDateTimeZoneFormat(t *testing.T) {
	init := replaceFuncWithValue(`$func("chain:dateNow::Standard>>append:::format:America/Lima>>encrypt::dateTimeZoneFormat>>append:::format:2006-01-02T15:04:05>>encrypt::dateTimeFormat")`)
	fmt.Printf("\nTestFuncDateTimeZoneFormat>>>%v\n", init)
}

func TestAESECB(t *testing.T) {
	init := replaceFuncWithValue(`$func("aes-ecb:sometext:param:keywith16length!")`)
	fmt.Printf("\nTestAESECB>>>%s\n", init)
}

func TestFuncNormalizeDateWithAdjustmentAndFormat(t *testing.T) {
	// add 5 min
	init := replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:add:60*5")`)
	expect := "2022-11-07 04:45:39"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// add 5 hour with tz
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-06T23:40:39Z:add:60*60*5:format:YYYY-MM-DDThh:mm:ssZ")`)
	expect = "2022-11-07T04:40:39Z"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// subtract 5 hour with tz
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39Z:subtract:60*60*5:format:YYYY-MM-DDThh:mm:ssZ")`)
	expect = "2022-11-06T23:40:39Z"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// test nano format
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39.99999Z:subtract:60*60*5:format:YYYY-MM-DD hh:mm:ss.nnnn")`)
	expect = "2022-11-06 23:40:39.9999"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	// test nano format with tz
	init = replaceFuncWithValue(`$func("dateFormat:2022-11-07T04:40:39.99999Z:subtract:60*60*5:format:YYYY-MM-DDThh:mm:ss.nnnnZ")`)
	expect = "2022-11-06T23:40:39.9999Z"
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncToLowerCaseEncryption(t *testing.T) {
	normal := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F"
	expect := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953e9-9830-4c41-9e7d-173fe93a784f"
	init := funcEncrToLowerCase(normal)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	hash := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:hash"
	expect = "3113325545"
	init = funcEncrToLowerCase(hash)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	md5 := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:md5"
	expect = "dcf64ec74c4d892c6d541b2288e82ac4"
	init = funcEncrToLowerCase(md5)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	base64 := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:base64"
	expect = "3PZOx0xNiSxtVBsiiOgqxA=="
	init = funcEncrToLowerCase(base64)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	sha256 := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:sha256"
	expect = "35355aa4d79c39110dd97563c3d4de0b991847a85432ecab3712bf14177ebb12"
	init = funcEncrToLowerCase(sha256)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	sha256 = "NKG2022-11-18 09:01:472022-11-18 09:01:47SecretKey:encr:sha256"
	expect = "2f6057738e3de44cec9af0376d4bf41a3982e0aef7eb0a0b6833cb8744d2c5fc"
	init = funcEncrToLowerCase(sha256)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncToUpperCaseEncryption(t *testing.T) {
	normal := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F"
	expect := "40813778-6A91-ED11-9D7A-00224819278B3C784854-2A22-EA11-A601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F"
	init := funcEncrToUpperCase(normal)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	hash := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:hash"
	expect = "598253769"
	init = funcEncrToUpperCase(hash)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	md5 := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:md5"
	expect = "c3439f20866862eb7a4d1e24542ef33b"
	init = funcEncrToUpperCase(md5)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	base64 := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:base64"
	expect = "w0OfIIZoYut6TR4kVC7zOw=="
	init = funcEncrToUpperCase(base64)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	sha256 := "40813778-6a91-ed11-9d7a-00224819278b3c784854-2a22-ea11-a601-281878584619644953E9-9830-4C41-9E7D-173FE93A784F:encr:sha256"
	expect = "fd2a6d020847dd7973d2e386b2068473c56114f4ae2ecaad385ab6f5d57c1dd0"
	init = funcEncrToUpperCase(sha256)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}

	sha256 = "NKG2022-11-18 09:01:472022-11-18 09:01:47SecretKey:encr:sha256"
	expect = "c0df99660f43a0b56bdda4df6142736f697e5cbb27328893944bdb2c0d4e5dce"
	init = funcEncrToUpperCase(sha256)
	if init != expect {
		t.Errorf("got %v, want %v", expect, init)
	}
}

func TestFuncReplaceToUUID(t *testing.T) {
	template_uuid := `http://hello.com/$func("uuid:$UUID")`
	init := replaceFuncWithValue(template_uuid)

	guid := strings.Split(init, "/")[3]
	_, err := uuid.Parse(guid)
	if err != nil {
		t.Errorf("got %v", err)
	}
	fmt.Println(init)
}

func TestFuncDecodeBase64ToStr(t *testing.T) {
	str := "b74bf098-b7ec-4593-80b9-e4194fbc12bf"
	base64 := "Yjc0YmYwOTgtYjdlYy00NTkzLTgwYjktZTQxOTRmYmMxMmJm"

	base64ToStr := funcDecodeBase64ToStr(base64)
	if base64ToStr != str {
		t.Errorf("result from decode base64 not matched!")
	}
}

func TestChainingFunction(t *testing.T) {
	chain := "chain:arrayPos::secondTick::http://dev1,http://dev2,http://dev3,http://dev4"
	encrypt := "@dateOffset:::subtract:60*60*4:format:YYMMDD>>append::AAAAA>>encrypt::md5>>prepend::ID_1=123456&ID_2=123456>>encrypt::md5"
	prepend := ">>prepend::/api/public?Name=abcdefg&ID=abc&Key=abcdef"
	append := ">>append::abcdef&Lang=id-ID"
	join := `$func("` + chain + encrypt + append + prepend + `")`

	offset := replaceFuncWithValue(`$func("dateOffset:::subtract:60*60*4:format:YYMMDD")`)
	result := replaceFuncWithValue(join)

	fmt.Println("offset: ", offset)
	fmt.Println("result: ", result)
}

func TestGetArrayPositionBySeconds(t *testing.T) {
	server := replaceFuncWithValue(`$func("arrayPos:secondTick::http://dev1,http://dev2,http://dev3,http://dev4")`)
	fmt.Println("SERVER_RESULT: ", server)
}

func TestGetDateNow(t *testing.T) {
	server := replaceFuncWithValue(`$func("dateNow:Unix")`)
	fmt.Println("SERVER_RESULT: ", server)
}

func TestHttpReqChain(t *testing.T) {
	// # normal function
	txt := `$func("httpReq:GET::https://pokeapi.co/api/v2/pokemon?offset=1&limit=1::{"Connection":"keep-alive"}::""::results.0.name")`
	result := replaceFuncWithValue(txt)
	fmt.Println(result)

	// # with chain function
	txt = `$func("chain:httpReq::GET::https://pokeapi.co/api/v2/pokemon?offset=1&limit=1::{"Connection":"keep-alive"}::""::results.0.name")`
	result = replaceFuncWithValue(txt)
	fmt.Println(result)

}
