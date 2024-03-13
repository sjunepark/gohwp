# hwp

> 한글과컴퓨터 문서(`.hwp`)와 상호작용하기 위한 Library 입니다.

## 설치

```bash
go get github.com/sjunepark/hwp
```

## 사용법

```go
package main

import (
	"fmt"
	"github.com/sjunepark/hwp"
	"github.com/sjunepark/hwp/internal/parser"
	"github.com/sjunepark/hwp/internal/reader"
)

func main() {
	raw, encrypted, err := reader.Read("data/example.hwp")
	if err != nil {
		fmt.Println(err)
	}
	if encrypted {
		fmt.Println("Document is encrypted")
	}
}


```