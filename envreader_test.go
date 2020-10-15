package envreader

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	os.Clearenv()

	Load("test/.env")
	test := os.Getenv("test")
	hello := os.Getenv("hello")
	if test != "test" {
		t.Errorf("Не найдена перменная test: ожидалось \"test\" полученные данные \"%s\", длина: %d", test, len(test))
	}

	if hello != "world" {
		t.Errorf("Не найдена перменная hello: ожидалось \"world\" полученные данные \"%s\", длина: %d", hello, len(hello))
	}
}
