package fibonacci

// Структура описывающая порядковый номер числа Фибоначчи и его значение
type JsonFibElem struct {
	SerialNumber int
	Value        int
}

func GetSerialFibNum(n int) int {
	var a = 1
	var b = 1
	var c = 0
	var count = 0
	for count < n {
		count++
		b = a
		a = c
		c = a + b
	}
	return c
}
