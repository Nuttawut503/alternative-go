package creational_patterns

import "fmt"

/* Define an interface for creating an object,
but let subclasses decide which class to instantiate.
The Factory method lets a class defer instantiation it uses to subclasses.
*/

type File struct {
	content, extension string
}

func (file File) String() string {
	return fmt.Sprintf("extension: %s\ncontent: %s", file.extension, file.content)
}

type FileMaker interface {
	MakeFile(string) File
}

type FileFactory struct {
	text  string
	maker FileMaker
}

func (fileFactory *FileFactory) Write(text string) {
	fileFactory.text += text
}

func (fileFactory *FileFactory) MakeFile() File {
	return fileFactory.maker.MakeFile(fileFactory.text)
}

type TxtMaker struct{}

type PDFMaker struct{}

func (txtMaker TxtMaker) MakeFile(text string) File {
	file := File{
		text,
		"txt",
	}
	return file
}

func (pdfMaker PDFMaker) MakeFile(text string) File {
	file := File{
		text,
		"pdf",
	}
	return file
}

type FileType int

const (
	TXT FileType = iota
	PDF
)

func (fileType FileType) String() string {
	return []string{"TXT", "PDF"}[fileType]
}

func NewFileFactory(fileType FileType) FileFactory {
	var fileFactory FileFactory
	switch fileType {
	case TXT:
		fileFactory = FileFactory{
			maker: TxtMaker{},
		}
	case PDF:
		fileFactory = FileFactory{
			maker: PDFMaker{},
		}
	}
	return fileFactory
}

func FactoryMethodExample() {
	fileFactory := NewFileFactory(PDF)
	fileFactory.Write("Hello")
	fileFactory.Write(" World!")
	file := fileFactory.MakeFile()
	fmt.Println(file)
}
