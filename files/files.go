package files

import "io/ioutil"

func WriteFile(path string, data []byte) {
	ioutil.WriteFile(path, data, 0644)
}
