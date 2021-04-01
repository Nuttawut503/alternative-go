package structural_patterns

import "fmt"

/* Decouple an abstraction from its implementation so that the two can vary independently
 */
type (
	File struct {
		text string
	}

	OSUploader interface {
		Upload(File)
	}

	LinuxUploader   struct{}
	WindowsUploader struct{}
	MacOSUploader   struct{}
)

func (linux LinuxUploader) Upload(file File) {
	fmt.Printf("Uploading... %v on Linux\n", file)
}

func (windows WindowsUploader) Upload(file File) {
	fmt.Printf("Uploading... %v on Windows\n", file)
}

func (macos MacOSUploader) Upload(file File) {
	fmt.Printf("Uploading... %v on MacOS\n", file)
}

type FileManager struct {
	File
	uploader OSUploader
}

func (manager FileManager) Upload() {
	manager.uploader.Upload(manager.File)
}

func BridgeExample() {
	fileManager := FileManager{
		File:     File{text: "Hello World!"},
		uploader: WindowsUploader{},
	}
	fileManager.Upload()
}
