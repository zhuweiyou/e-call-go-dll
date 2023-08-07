package main

/*
#include <stdlib.h>

typedef void (__stdcall *Callback)(char *response);
void callCallback(Callback callback, char *response);
*/
import "C"
import (
	"strings"
	"unsafe"
)

//export Call
func Call(req *C.char, cb C.Callback) {
	reqBytes, _ := GbkToUtf8([]byte(C.GoString(req)))
	resBytes, _ := Utf8ToGbk(Handler(reqBytes))
	resString := C.CString(string(resBytes))
	defer C.free(unsafe.Pointer(resString))
	C.callCallback(cb, resString)
}

func Handler(reqBytes []byte) []byte {
	reqString := string(reqBytes)
	if strings.ToUpper(reqString) == "PING" {
		return []byte("PONG")
	}
	return []byte("ERROR")
}

func main() {
}
