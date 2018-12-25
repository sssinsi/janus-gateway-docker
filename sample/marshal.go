package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

// https://play.golang.org/p/WHRgvgrsG4
func unmarshalUseNumber(b []byte, i interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	if err := dec.Decode(i); err != nil {
		return err
	}
	if dec.More() {
		remaining, _ := ioutil.ReadAll(dec.Buffered())
		return fmt.Errorf("unexpected values after JSON element %q", remaining)
	}
	return nil
}

func marshaltest(jsonstring1 string, uselongint bool) string {
	fmt.Println("INPUT :", jsonstring1)

	// unmarshal
	var jsonstruct interface{}
	if uselongint {
		unmarshalUseNumber([]byte(jsonstring1), &jsonstruct)
	} else {
		json.Unmarshal([]byte(jsonstring1), &jsonstruct)
	}
	if jsonstruct == nil {
		fmt.Println("Unmarshal failed !!")
		return ""
	}
	// print struct information
	numbervalue := jsonstruct.(interface{}).(map[string]interface{})["key"]
	fmt.Print("Unmarshaled:", numbervalue)
	fmt.Println("  Type:", reflect.TypeOf(numbervalue))
	// marshal
	jsonstring2, err := json.Marshal(jsonstruct)
	if err != nil {
		panic(err)
	}
	fmt.Println("OUTPUT:", string(jsonstring2))
	return string(jsonstring2)
}

func main() {
	// setup numberstring & jsonstring
	var numberstring = "null"
	if len(os.Args) >= 2 {
		numberstring = os.Args[1]
	}
	var USE_LONGINT = false
	if len(os.Args) >= 3 {
		USE_LONGINT = true
	}
	_ = marshaltest("{\"key\":"+numberstring+"}", USE_LONGINT)
}
