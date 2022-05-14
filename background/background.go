package background

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParametersInfoW")
)

func ChangeBackground(path string) {
	imgPath, _ := windows.UTF16PtrFromString(path)
	fmt.Println("[+] Changing background now...")
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imgPath)), 0x001A)
}
