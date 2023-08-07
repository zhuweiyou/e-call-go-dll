package main

/*
typedef void (__stdcall *Callback)(char *response);
void callCallback(Callback callback, char *response) {
    (callback)(response);
}
*/
import "C"
