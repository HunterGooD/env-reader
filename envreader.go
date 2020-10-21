// Package envreader Для чтения .env  файлов
package envreader

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

// начало конец
const (
	BEGIN = "\033["
	END   = "m"
)

// printError ...
func printError(data string) {
	// первый аргумент красит в красный потом выводит ошибку очищает покраску и переносит на новую строку
	fmt.Printf("%sError:%s%s\n", fmt.Sprintf("%s%d%s", BEGIN, 31, END), data, fmt.Sprintf("%s%d%s", BEGIN, 0, END))
}

// Load Загрузка всех env файлов
func Load(files ...string) error {
	if len(files) == 0 {
		return errors.New("Не переданы файлы для загрузки")
	}
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			read(file)
			wg.Done()
		}(file)
	}
	wg.Wait()
	return nil
}

func read(file string) {
	info := strings.Split(file, ".")
	if info[1] != "env" {
		printError("File: " + file + " not configure file .env")
		return
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		printError(err.Error())
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		var line string = scanner.Text()
		if len(line) > 0 {
			if line[0] == '#' {
				continue
			}
		}
		env := strings.Split(line, "=")
		if len(env) == 2 {
			env[1] = strings.Trim(env[1], "\"")
			err := os.Setenv(env[0], env[1])
			if err != nil {
				printError(err.Error())
			}
		}
	}
}
