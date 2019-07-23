package image_picker

import (
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"

	"github.com/pkg/errors"
)

var (
	comdlg32         = syscall.NewLazyDLL("comdlg32.dll")
	getOpenFileNameW = comdlg32.NewProc("GetOpenFileNameW")
)

const (
	maxPath          = 260
	ofnExplorer      = 0x00080000
	ofnFileMustExist = 0x00001000
	ofnHideReadOnly  = 0x00000004
)

type openfilenameW struct {
	lStructSize       uint32
	hwndOwner         syscall.Handle
	hInstance         syscall.Handle
	lpstrFilter       *uint16
	lpstrCustomFilter *uint16
	nMaxCustFilter    uint32
	nFilterIndex      uint32
	lpstrFile         *uint16
	nMaxFile          uint32
	lpstrFileTitle    *uint16
	nMaxFileTitle     uint32
	lpstrInitialDir   *uint16
	lpstrTitle        *uint16
	flags             uint32
	nFileOffset       uint16
	nFileExtension    uint16
	lpstrDefExt       *uint16
	lCustData         uintptr
	lpfnHook          syscall.Handle
	lpTemplateName    *uint16
	pvReserved        unsafe.Pointer
	dwReserved        uint32
	flagsEx           uint32
}

func utf16PtrFromString(s string) *uint16 {
	b := utf16.Encode([]rune(s))
	return &b[0]
}

func stringFromUtf16Ptr(p *uint16) string {
	b := *(*[maxPath]uint16)(unsafe.Pointer(p))
	r := utf16.Decode(b[:])
	return strings.Trim(string(r), "\x00")
}

func getOpenFileName(lpofn *openfilenameW) bool {
	ret, _, _ := getOpenFileNameW.Call(uintptr(unsafe.Pointer(lpofn)), 0, 0)
	return ret != 0
}

func fileDialog(title string, fileType string) (string, error) {
	var ofn openfilenameW
	buf := make([]uint16, maxPath)

	t, _ := syscall.UTF16PtrFromString(title)

	ofn.lStructSize = uint32(unsafe.Sizeof(ofn))
	ofn.lpstrTitle = t
	ofn.lpstrFile = &buf[0]
	ofn.nMaxFile = uint32(len(buf))

	var filters string
	switch fileType {
	case "image":
		filters = "Images (*.jpeg,*.png,*.gif)\x00*.jpg;*.jpeg;*.png;*.gif\x00All Files (*.*)\x00*.*\x00\x00"
	case "video":
		filters = "Videos (*.webm,*.wmv,*.mpeg,*.mkv,*.mp4,*.avi,*.mov,*.flv)\x00*.webm;*.wmv;*.mpeg;*.mkv;*mp4;*.avi;*.mov;*.flv\x00All Files (*.*)\x00*.*\x00\x00"
	default:
		return "", errors.New("unsupported fileType")
	}

	if filters != "" {
		ofn.lpstrFilter = utf16PtrFromString(filters)
	}

	flags := ofnExplorer | ofnFileMustExist | ofnHideReadOnly

	ofn.flags = uint32(flags)

	if getOpenFileName(&ofn) {
		return stringFromUtf16Ptr(ofn.lpstrFile), nil
	}

	return "", nil

}
