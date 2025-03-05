package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	kernel32         = windows.NewLazySystemDLL("kernel32.dll")
	messageBox       = user32.NewProc("MessageBoxW")
	getConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	showWindow       = user32.NewProc("ShowWindow")
)

const (
	MB_OK   = 0x00000000
	SW_HIDE = 0
)

func main() {
	// Скрываем консольное окно, если оно есть
	hwnd, _, _ := getConsoleWindow.Call()
	if hwnd != 0 {
		showWindow.Call(hwnd, SW_HIDE)
	}

	// Бесконечный цикл сообщений
	for {
		showErrorDialog()
	}
}

func showErrorDialog() {
	title, _ := windows.UTF16PtrFromString("Ошибка")
	text, _ := windows.UTF16PtrFromString("Теперь на твоём компьютере 21 вирус, поздравляю!")
	messageBox.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), MB_OK)
}
