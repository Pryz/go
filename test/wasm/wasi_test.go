package wasi_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
    //"io/fs"
)

func TestFsys(t*testing.T) {
    // `tmp` will be created at the root of the FS.
	tmp := "wasi_testfs"

    fmt.Println("Create filesystem")
	err := os.Mkdir(tmp, 0777)
	if err != nil {
		log.Fatal("MkdirTemp: ", err)
	}
	defer func() {
        fmt.Println("Cleaning things up")
        //os.RemoveAll(tmp)
    }()

    //tmp = "./" + tmp
	//tmpfs := os.DirFS(tmp)
	tmpfs := os.DirFS("/")

	if err := os.WriteFile(filepath.Join(tmp, "hello"), []byte("hello, world\n"), 0777); err != nil {
		t.Fatal(err)
	}

    //fmt.Println("Open filesystem")
    //f, err := tmpfs.Open(tmp)
    //if err != nil {
    //    fmt.Println("failed to open dir ", err)
    //    t.Fatal(err)
    //}

    //d, ok := f.(fs.ReadDirFile)
    //if !ok {
    //    t.Fatalf("not a directory")
    //}

    //entries, err := d.ReadDir(-1)
    //if err != nil {
    //    t.Fatal(err)
    //}

    hello_path := "wasi_testfs/hello"
    link_path := "wasi_testfs/hello.link"

	if err := os.Symlink(hello_path, link_path); err != nil {
	    t.Fatal(err)
	}

	if err := fstest.TestFS(tmpfs, "hello", "hello.link"); err != nil {
		t.Fatal(err)
	}

}

//FIXME
//func preparePath(path string, followTrailingSymlink bool) (*byte, size_t) {
//	if path == "" || path[0] != '/' {
//		path = wd + "/" + path
//	}
//
//	parts := strings.Split(path[1:], "/")
//	resolvedPath := ""
//	for i, part := range parts {
//		resolvedPath += "/" + part
//		if i == len(parts)-1 && !followTrailingSymlink {
//			break
//		}
//		for {
//			dest, err := readlink("." + resolvedPath)
//			if err != nil {
//				break
//			}
//			if dest[0] != '/' {
//				i := strings.LastIndexByte(resolvedPath, '/')
//				dest = resolvedPath[:i] + "/" + dest
//			}
//			resolvedPath = dest
//		}
//	}
//
//    // Oh my
//	return &[]byte("." + resolvedPath)[0], size_t(1 + len(resolvedPath))
//}

//func TestDirent(t*testing.T) {
//    t.SkipNow()
//    fh, err := os.Open("/")
//    if err != nil {
//        t.Fatal(err)
//    }
//
//    buf := make([]byte, 8192)
//
//    _, err = fh.ReadDirent(buf)
//    if err != nil {
//        t.Fatal(err)
//    }
//
//    names := []string{}
//    consumed, count, newnames := syscall.ParseDirent(buf, 50, names)
//    fmt.Println(consumed, count, newnames)
//    fmt.Println(names)
//}

func TestFilesystem(t *testing.T) {
	// /tmp
	t.SkipNow()
	tmp := "/wasi_testfs"
	err := os.Mkdir(tmp, 0777)
	if err != nil {
		log.Fatal("MkdirTemp: ", err)
	}
	defer os.RemoveAll(tmp)

	tmpfs := os.DirFS(tmp)

	if err := os.WriteFile(filepath.Join(tmp, "hello"), []byte("hello, world\n"), 0777); err != nil {
		t.Fatal(err)
	}

	// WASI doesn't support symlink with absolute path ?
	if err := os.Chdir(tmp); err != nil {
	    t.Fatal(err)
	}

	//if err := os.Symlink("hello", "hello.link"); err != nil {
	//t.Fatal(err)
	//}

	if err := fstest.TestFS(tmpfs, "hello", "hello.link"); err != nil {
		t.Fatal(err)
	}
}
