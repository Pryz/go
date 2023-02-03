package wasi_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
	"syscall"
	"fmt"
)

func TestFilesystem(t *testing.T) {
	// /tmp
	//t.SkipNow()
	tmp, err := os.MkdirTemp("/", "wasi_test")
	if err != nil {
		log.Fatal("MkdirTemp: ", err)
	}
	defer os.RemoveAll(tmp)

	tmpfs := os.DirFS(tmp)

	if err := os.WriteFile(filepath.Join(tmp, "hello"), []byte("hello, world\n"), 0777); err != nil {
		t.Fatal(err)
	}

	//if err := os.Symlink(filepath.Join(tmp, "hello"), filepath.Join(tmp, "hello.link")); err != nil {
	//	t.Fatal(err)
	//}

	//if err := fstest.TestFS(tmpfs, "hello", "hello.link"); err != nil {
	if err := fstest.TestFS(tmpfs, "hello"); err != nil {
		t.Fatal(err)
	}
}

func TestFS2(t *testing.T) {
	var stat syscall.Fdstat_t
	errno := syscall.Fd_fdstat_get(syscall.Fd_t(3), &stat)
	if errno != 0 {
		panic("could not get fdstat of root: " + errno.Error())
	}

	fmt.Println("rootFD")
	fmt.Printf("%+v\n", stat)
	fmt.Printf("%b\n", stat.RightsBase)
	fmt.Printf("%b\n", stat.RightsInheriting)

	tmp := "/wasi_fs2"
	fmt.Println("mkdir ", tmp)
	if err := syscall.Mkdir(tmp, 0777); err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmp)

	//tmpfs := os.DirFS(tmp)

	fmt.Println("open ", tmp+"/foobar")
	fd, err := syscall.Open(filepath.Join(tmp, "foobar"), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o777)
	if err != nil {
		panic(err)
	}


	fmt.Println("fdstat ", tmp+"/foobar")
	errno = syscall.Fd_fdstat_get(syscall.Fd_t(fd), &stat)
	if errno != 0 {
		panic("could not get fdstat of file: " + errno.Error())
	}

	fmt.Println(filepath.Join(tmp, "foobar"))
	fmt.Printf("%+v\n", stat)
	fmt.Printf("%b\n", stat.RightsBase)
	fmt.Printf("%b\n", stat.RightsInheriting)

	fmt.Println("close ", tmp+"/foobar")
	if err := syscall.Close(fd); err != nil {
		panic(err)
	}

	fmt.Println("symlink ", tmp+"/foobar")
	//err = syscall.Symlink(tmp +"/foobar", tmp+"/foobar.link")
	err = syscall.Symlink("./foobar", "./foobar.link")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(filepath.Join(tmp, "hello"), []byte("hello, world\n"), 0777); err != nil {
		t.Fatal(err)
	}


	if err := os.Symlink(filepath.Join(tmp, "hello"), "/hello.link"); err != nil {
		t.Fatal(err)
	}

	//if err := fstest.TestFS(tmpfs, "hello", "hello.link"); err != nil {
	//	t.Fatal(err)
	//}

}
