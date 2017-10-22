package stack

import (
	"bytes"
	"math"
	"strconv"
	"testing"
)

func TestStart(t *testing.T) {
	var num, y Elem
	num = 11
	var stack StackLink
	stack.InitStack()
	for {
		if num == 0 {
			break
		}
		y = Elem(num % 2)
		num = Elem(math.Floor((float64(num) / 2)))
		t.Logf("y: %d", y)
		// 入栈
		stack.Push(y)
	}
	p := new(Elem)
	var buffer bytes.Buffer
	for {
		if stack.Empty() {
			break
		}
		// 出栈
		err := stack.Pop(p)
		if err != nil {
			t.Logf(err.Error())
		}
		t.Log(byte(*p))
		buffer.WriteString(strconv.Itoa(int(*p)))
		// buffer.WriteByte(byte(*p))
	}

	t.Logf("二进制为: %s", buffer.String())
	// t.Log(buffer.Bytes())

	t.Logf("end")
}
