package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var defaultUser string = "Karmarchari32"

func ErrorHappened(msg ...interface{}) {
	fmt.Print("Error")
	fmt.Print(msg...)
	os.Exit(1)
}

func printl(msg ...interface{}) {
	fmt.Println(msg...)

}

func generateResp(w http.ResponseWriter, data interface{}, err error) {

	if data != nil {
		m := make(map[string]interface{})
		m["status"] = "ok"
		m["data"] = data
		m["err"] = err // this could be used to send deprecated api Warning
		response, err2 := enJson(m)
		if err2 != nil {
			errorResp(w, err2)
			return
		}
		w.Write(response)
		return
	}
	errorResp(w, err)
}

func errorResp(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.Write([]byte(`{"msg":"` + err.Error() + `","status":"failed"}`))
}

func enJson(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func deJson(raw []byte, out interface{}) error {
	return json.Unmarshal(raw, out)
}
