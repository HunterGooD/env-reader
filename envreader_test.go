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

func TestLoadEnvDop(t *testing.T) {
	os.Clearenv()

	var answers = map[string]string{
		"test":     "test",
		"hello":    "world",
		"host":     "localhost",
		"port":     "8080",
		"password": "",
		"dbname":   "test",
	}
	Load("test/.env", "test/db.env")

	for key, value := range answers {
		env := os.Getenv(key)
		if value != env {
			t.Errorf("Не найдена перменная %s: ожидалось \"%s\" полученные данные \"%s\", длина: %d", key, value, env, len(env))
		}
	}
}
