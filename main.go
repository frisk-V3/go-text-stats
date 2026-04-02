package main

import (
    "bufio"
    "fmt"
    "os"
    "os/signal"
    "strings"
    "unicode/utf8"
)

// 短い解説（日本語）
// Go で「文字数」を正しく数えるには utf8.RuneCountInString を使います。
// len(s) はバイト数を返します（UTF-8 のバイト長）。

func main() {
    // Ctrl+C を受け取ってプログラムを終了するためのシグナルハンドラ
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt)
    go func() {
        <-sigCh
        fmt.Fprintln(os.Stderr, "\n受信したため終了します。")
        os.Exit(0)
    }()

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("テキストを入力してください。空行で終了します。Ctrl+C でも終了できます。")

    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            // EOF やエラーで終了
            fmt.Fprintln(os.Stderr, "\n入力が終了したため終了します。")
            return
        }
        line := scanner.Text()

        // 空行で終了
        if strings.TrimSpace(line) == "" {
            fmt.Println("空行が入力されたため終了します。")
            return
        }

        // 文字数（ルーン数）
        runeCount := utf8.RuneCountInString(line)

        // 単語数（strings.Fields を使用）
        words := strings.Fields(line)
        wordCount := len(words)

        // バイト数（len を使用）
        byteCount := len(line)

        fmt.Printf("文字数: %d\n", runeCount)
        fmt.Printf("単語数: %d\n", wordCount)
        fmt.Printf("バイト数: %d\n", byteCount)
    }
}
