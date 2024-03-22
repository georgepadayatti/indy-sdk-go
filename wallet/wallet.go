package wallet

/*
#cgo CFLAGS: -I./../include
#cgo LDFLAGS: -lindy

#include <stdio.h>
#include <stdlib.h>
#include "../include/indy_core.h"

extern void cb(indy_handle_t command_handle, indy_error_t err);
static void create_wallet_cb_wrapper(indy_handle_t command_handle, indy_error_t err) {
    cb(command_handle, err);
}

static void create_wallet(char *config, char *credentials) {
    indy_create_wallet(0, config, credentials, create_wallet_cb_wrapper);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Response struct {
	Err int
}

var responseChannel = make(chan Response)

//export cb
func cb(_ C.indy_handle_t, err C.indy_error_t) {
	responseChannel <- Response{Err: int(err)}
}

// CreateWallet creates indy wallet and returns response that contains err in it.
func CreateWallet(id string, key string) Response {
	config := C.CString(fmt.Sprintf(`{"id": "%v"}`, id))
	defer C.free(unsafe.Pointer(config))

	credentials := C.CString(fmt.Sprintf(`{"key": "%v"}`, key))
	defer C.free(unsafe.Pointer(credentials))

	C.create_wallet(config, credentials)

	res := <-responseChannel

	defer close(responseChannel)
	return res
}
