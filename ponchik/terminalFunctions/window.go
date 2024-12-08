package windows

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func SetWindow(width, height int) error {
	// Define _COORD structure
	type _COORD struct {
		X int16
		Y int16
	}

	// Define _SMALL_RECT structure
	type _SMALL_RECT struct {
		Left   int16
		Top    int16
		Right  int16
		Bottom int16
	}

	// Get handle to the console output
	handle, err := syscall.GetStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return fmt.Errorf("failed to get console handle: %v", err)
	}

	// Set console screen buffer size
	coord := _COORD{
		X: int16(width),
		Y: int16(height),
	}
	ret, _, err := windows.NewLazySystemDLL("kernel32.dll").
		NewProc("SetConsoleScreenBufferSize").
		Call(uintptr(handle), uintptr(*(*int32)(unsafe.Pointer(&coord))))
	if ret == 0 {
		return fmt.Errorf("failed to set screen buffer size: %v", err)
	}

	// Define console window rectangle
	rect := _SMALL_RECT{
		Top:    0,
		Left:   0,
		Bottom: int16(height - 1),
		Right:  int16(width - 1),
	}
	ret, _, err = windows.NewLazySystemDLL("kernel32.dll").
		NewProc("SetConsoleWindowInfo").
		Call(uintptr(handle), uintptr(1), uintptr(unsafe.Pointer(&rect)))
	if ret == 0 {
		return fmt.Errorf("failed to set console window size: %v", err)
	}

	return nil
}
