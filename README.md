# docconv

[![Go reference](https://pkg.go.dev/badge/code.sajari.com/docconv/v2.svg)](https://pkg.go.dev/code.sajari.com/docconv/v2)
[![Build status](https://github.com/sajari/docconv/workflows/Go/badge.svg?branch=master)](https://github.com/sajari/docconv/actions)
[![Report card](https://goreportcard.com/badge/code.sajari.com/docconv/v2)](https://goreportcard.com/report/code.sajari.com/docconv/v2)
[![Sourcegraph](https://sourcegraph.com/github.com/sajari/docconv/v2/-/badge.svg)](https://sourcegraph.com/github.com/sajari/docconv/v2)

Go 語言的封裝函式庫，可將 PDF、DOC、DOCX、XML、HTML、RTF、ODT、Pages 文件以及影像（可選擇安裝的相依套件）轉為純文字。

## 安裝

若尚未安裝 Go，請先依照 [官方指南](https://golang.org/doc/install) 完成安裝。

接著抓取並建置程式碼：

```console
$ go install code.sajari.com/docconv/v2/docd@latest
```

更多關於安裝後 `docd` 可執行檔路徑的說明，請參閱 `go help install`。安裝完成後請確定其路徑已加入 `PATH` 環境變數。

## 依賴套件

- tidy
- wv
- popplerutils
- unrtf
- libreoffice（或 unoconv）
- https://github.com/JalfResi/justext

### Debian 系 Linux

```console
$ sudo apt-get install poppler-utils wv unrtf tidy libreoffice
$ go get github.com/JalfResi/justext
```

### macOS

```console
$ brew install poppler-qt5 wv unrtf tidy-html5 libreoffice
$ go get github.com/JalfResi/justext
```

### 可選依賴

若要讓 `docconv` 支援影像轉換，必須先 [安裝並建置 gosseract](https://github.com/otiai10/gosseract/tree/v2.2.4)。

完成後，於建置、下載或測試 `docconv` 時加上 `-tags ocr` 參數，即可啟用影像處理：

```console
$ go get -tags ocr code.sajari.com/docconv/v2/...
```

在 macOS 上若遇到錯誤，可透過 brew 安裝 [tesseract](https://tesseract-ocr.github.io) 解決：

```console
$ brew install tesseract
```

## docd 工具

`docd` 可以以下列幾種方式運行：

1.  預設在 8888 埠口上運行的服務

    可透過 multipart POST 傳送文件，服務會回傳純文字內容與後設資訊的 JSON。

2.  於 Docker container 中執行的服務

    同樣以服務方式執行，但封裝在 Docker container 中。
    官方映像位於 https://hub.docker.com/r/sajari/docd。

    你也可以自行建置：

    ```console
    $ cd docd
    $ docker build -t docd .
    ```

3.  直接使用指令列。

    將檔案路徑作為參數，例如：

    ```console
    $ docd -input document.pdf
    ```

### 其他旗標

- `addr` - HTTP 伺服器綁定地址，預設為 ":8888"
- `readability-length-low` - 配合 ?readability=1 參數設定 readability 的 length low
- `readability-length-high` - 配合 ?readability=1 參數設定 readability 的 length high
- `readability-stopwords-low` - 配合 ?readability=1 參數設定 stopwords low
- `readability-stopwords-high` - 配合 ?readability=1 參數設定 stopwords high
- `readability-max-link-density` - 配合 ?readability=1 參數設定最大連結密度
- `readability-max-heading-distance` - 配合 ?readability=1 參數設定最大標題距離
- `readability-use-classes` - 配合 ?readability=1 參數指定 readability 類別，逗號分隔

### 啟動服務

```console
$ # 於 8000 埠口啟動
$ docd -addr :8000
```

## 範例程式

以下範例僅供參考，實際情況通常是透過 HTTP 收取檔案或從檔案系統讀取。

這應該足以讓你開始使用。

### 使用情境一：在本機執行

> 注意：假設你已安裝好[依賴套件](#依賴套件)。

```go
package main

import (
	"fmt"

	"code.sajari.com/docconv/v2"
)

func main() {
	res, err := docconv.ConvertPath("your-file.pdf")
	if err != nil {
		// TODO: handle
	}
	fmt.Println(res)
}
```

### 使用情境二：透過網路請求

```go
package main

import (
	"fmt"

	"code.sajari.com/docconv/v2/client"
)

func main() {
        // 建立新客戶端，預設端點為 localhost:8888
        c := client.New()

	res, err := client.ConvertPath(c, "your-file.pdf")
	if err != nil {
		// TODO: handle
	}
	fmt.Println(res)
}
```

也可以透過 `curl` 直接請求：

```console
$ curl -s -F input=@your-file.pdf http://localhost:8888/convert
```

### 將 DOCX 轉換為 PDF

若已安裝 libreoffice，可利用下列輔助函式將 DOCX 檔轉成 PDF：

```go
package main

import (
        "fmt"

        "code.sajari.com/docconv/v2"
)

func main() {
        pdfPath, err := docconv.ConvertDocxToPDF("document.docx", "/tmp")
        if err != nil {
                // TODO: handle
        }
        fmt.Println("PDF 輸出於", pdfPath)
}
```
