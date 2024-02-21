package script

import (
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "time"

    "iTools/conf"
)

func ListFiles(rootPath string) []conf.File {
    var files []conf.File
    fileModTimeMap := make(map[string]time.Time)

    err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            name := filepath.Base(path)
            modTime := info.ModTime()

            // 加入base64
            base64 := EnBase64(name)

            file := conf.File{name, base64}
            files = append(files, file)

            fileModTimeMap[file.Name] = modTime
        }
        return nil
    })
    if err != nil {
        fmt.Println(err)
        return nil
    }

    // 對文件按修改時間排序
    sort.Slice(files, func(i, j int) bool {
        return fileModTimeMap[files[i].Name].After(fileModTimeMap[files[j].Name])
    })

    return files
}

func FileNameWithRandNum(fileName string) string {
    ext := filepath.Ext(fileName)
    baseName := strings.Replace(fileName, ext, "", -1)
    rand := RandNumInt(0, 999)

    return fmt.Sprintf(`%s-%03v%s`, baseName, rand, ext)
}
