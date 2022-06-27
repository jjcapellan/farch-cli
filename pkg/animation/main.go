package animation

import (
	"fmt"
	"strings"
	"time"
)

const frameTime = 200 * time.Millisecond
const maxInfoLength int = 30

var _info string
var playing bool = false

func AnimatedInfo(info string) {
	frames := [4]string{"-", "\\", "|", "/"}
	_info = validateInfo(info)
	playing = true

	fmt.Print("\n")

	i := 0
	for playing {
		if i > 3 {
			i = 0
		}
		fmt.Print("\r[" + frames[i] + "] " + _info)
		i++
		time.Sleep(frameTime)
	}

	fmt.Print("\r[X] " + _info + "\n")
}

func validateInfo(info string) string {
	str := info
	length := len(str)

	if length >= maxInfoLength {
		str = str[:20]
		return str
	}

	n := maxInfoLength - length
	str = str + strings.Repeat(" ", n)

	return str
}

func SetInfo(info string) {
	_info = validateInfo(info)
}

func StopAnim(info string) {
	_info = validateInfo(info)
	playing = false
}
