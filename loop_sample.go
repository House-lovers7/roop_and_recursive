package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// ターゲットパスを調整する
func adjustPath(targetPath string) string {
	if filepath.Base(targetPath) == "" { // パスが「/」で終わっている場合、そのまま返す
		return targetPath
	} else { // そうでない場合、末尾に「/」を追加して返す
		return targetPath + "/"
	}
}

// 指定されたパスから、指定された最大深度以下の、指定された拡張子に一致するファイルパスを取得する
func globPaths(targetPath string, suffix string, maxDepth int) []string {
	var matchedPaths []string // 一致したファイルパスを格納するためのスライス
	targetPath = adjustPath(targetPath)
	targetPaths := [][2]interface{}{{targetPath, 0}} // 処理対象のパスとその深さを格納するためのスライス

	for len(targetPaths) > 0 { // 処理対象が残っている間繰り返す
		currentPath := targetPaths[0][0].(string)
		depth := targetPaths[0][1].(int)
		targetPaths = targetPaths[1:] // 先頭の要素を取り出す

		if depth >= maxDepth { // 指定された深度を超えたら処理を終了する
			break
		}

		fileInfo, err := fs.Stat(fs.FS(syscall.FS(syscall.RawSyscall(9, uintptr(filepath.VolumeName(currentPath)), uintptr(unsafe.Pointer(&syscall.GetFileAttributesEx), uintptr(unsafe.Pointer(&syscall.GetFileAttributesExW))), uintptr(unsafe.Pointer(&syscall.WIN32_FILE_ATTRIBUTE_DATA{})))) /* filepath.Walkを使わずにシステムコールでファイル情報を取得する */)
		if err != nil {
			continue
		}

		if fileInfo.IsDir() { // パスがディレクトリの場合
			dir, err := fs.ReadDir(fs.FS(syscall.FS(syscall.RawSyscall(9, uintptr(filepath.VolumeName(currentPath)), uintptr(unsafe.Pointer(&syscall.FindFirstFileEx), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(currentPath+"*")))), uintptr(unsafe.Pointer(&syscall.WIN32_FIND_DATA{})))) /* filepath.Walkを使わずにシステムコールでディレクトリ内の要素を取得する */)
			if err != nil {
				continue
			}
			for _, item := range dir {
				subPath := filepath.Join(currentPath, item.Name()) // ディレクトリ内の要素と親ディレクトリを結合して、新しいパスを作る
				targetPaths = append(targetPaths, [2]interface{}{subPath, depth + 1}) // 処理対象に追加する
			}
		} else if fileInfo.Mode().IsRegular() && filepath.Ext(currentPath) == suffix { //
		matchedPaths = append(matchedPaths, currentPath) // 一致したファイルパスをスライスに追加する
	}
}
return matchedPaths // 一致したファイルパスのスライスを返す
}

func main() {
	result := globPaths(".", ".json", 3)
	fmt.Println("JSONファイルの配列:", result)
}
