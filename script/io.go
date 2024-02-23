package script

import (
    "io/ioutil"
)

func ReadFile(path string) string {
    bytes, err := ioutil.ReadFile(path)
    printError(err)
    return string(bytes)
}

func WriteFile(content, path string) {
    err := ioutil.WriteFile(path, []byte(content), 0777)
    printError(err)
}
