# Package Envreader

Пакет для языка Golang для чтения .env файлов. Работает с файлами по типу.

```text
test=test

hello=world

test1="test"
```

То значение переменной `test1` будет `test` значение `hello` будет `world`

## Установка

```bash
go get github.com/HunterGooD/env-reader
```

## Использование

Подключить как обычный Go пакет и вызвать функцию Load

```go
package main

import (
    "fmt"
    "os"
    renv "github.com/HunterGooD/env-reader"
)

func init() {
    renv.Load(".env", "db.env", "apikeys.env")
}

func main() {
    // использовать переменные  через os.Getenv()
    fmt.Println(os.Getenv("ваш ключ для переменой"))
}
```
