package main

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

// VarGuide имитирует "класс" с методами, которые показывают основные приемы работы с переменными в Go.
type VarGuide struct{}

func (VarGuide) section(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}

func (g VarGuide) zeroValues() {
	g.section("Нулевые значения")
	var (
		b bool   // false
		i int    // 0
		s string // ""
		p *int   // nil
	)
	fmt.Printf("bool: %v, int: %v, string: %q, pointer: %v\n", b, i, s, p)
}

func (g VarGuide) typedAndShortDecl() {
	g.section("Объявление: явные типы и короткие литералы")
	var age int = 30       // явный тип
	var name = "Gopher"    // тип выводится из значения
	pi := 3.1415           // короткая форма (только внутри функций)
	answer, ok := 42, true // множественное присваивание
	fmt.Printf("age=%d, name=%s, pi=%.4f, answer=%d, ok=%v\n", age, name, pi, answer, ok)
}

func (g VarGuide) signedUnsigned() {
	g.section("Знаковые и беззнаковые числа")
	var small int8 = -8
	var large uint32 = 4000000000
	fmt.Printf("int8=%d, uint32=%d\n", small, large)
}

func (g VarGuide) floatsAndComplex() {
	g.section("float и complex")
	f32 := float32(1.5)
	f64 := 2.718281828
	z := complex(2, 3) // 2+3i
	fmt.Printf("float32=%.2f, float64=%.6f, complex=%v, |complex|=%.2f\n", f32, f64, z, cmplx.Abs(z))
}

func (g VarGuide) runeAndByte() {
	g.section("rune и byte")
	var letter rune = 'A' // rune = int32, символ Unicode
	var raw byte = 65     // byte = uint8
	fmt.Printf("rune=%c (%U), byte=%d\n", letter, letter, raw)
}

func (g VarGuide) stringsAreImmutable() {
	g.section("Строки неизменяемы")
	s := "go"
	fmt.Printf("строка: %s, длина байтов: %d\n", s, len(s))
	// s[0] = 'G' // так нельзя; нужно создавать новую строку
	s = "G" + s[1:]
	fmt.Printf("новая строка после \"изменения\": %s\n", s)
}

func (g VarGuide) arraysAndSlices() {
	g.section("Массивы и срезы")
	arr := [3]int{1, 2, 3}     // массив фиксированной длины
	slice := []int{1, 2, 3, 4} // срез гибкой длины
	slice = append(slice, 5)   // безопасное расширение
	fmt.Printf("array=%v, len=%d\n", arr, len(arr))
	fmt.Printf("slice=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
}

func (g VarGuide) maps() {
	g.section("Отображения (map)")
	ages := map[string]int{"Ann": 23}
	ages["Bob"] = 30
	fmt.Printf("map=%v, Bob существует? %v\n", ages, true)
	val, exists := ages["Kate"]
	fmt.Printf("чтение несуществующего ключа: value=%d, exists=%v\n", val, exists)
}

func (g VarGuide) structs() {
	g.section("Структуры")
	type Player struct {
		Name  string
		Score int
	}
	p := Player{Name: "Alex", Score: 10}
	p.Score += 5
	fmt.Printf("Player: %+v\n", p)
}

func (g VarGuide) pointers() {
	g.section("Указатели")
	var p *int // nil, пока ни на что не указывает
	fmt.Printf("p is nil? %v\n", p == nil)

	value := 10
	p = &value // берем адрес переменной
	*p = 20    // меняем оригинал через указатель
	fmt.Printf("value=%d через указатель=%d\n", value, *p)
}

func (g VarGuide) constants() {
	g.section("Константы и iota")
	const pi = 3.14 // тип выводится, пока не нужен конкретный
	const typedPi float64 = 3.14
	const (
		Red = iota
		Green
		Blue
	)
	fmt.Printf("pi=%v (type: %T), typedPi=%v, iota: %d %d %d\n", pi, pi, typedPi, Red, Green, Blue)
}

func (g VarGuide) typeConversion() {
	g.section("Явное приведение типов")
	var small int8 = 42
	big := int(small) // явное приведение; автоматического нет
	f := float64(big)
	fmt.Printf("small(int8)=%d -> big(int)=%d -> float64=%.1f\n", small, big, f)
}

func (g VarGuide) shadowing() {
	g.section("Теневание (shadowing)")
	x := 1
	if true {
		x := 2 // новая переменная скрывает внешнюю
		fmt.Printf("внутри блока x=%d\n", x)
	}
	fmt.Printf("снаружи блока x=%d\n", x)
}

func (g VarGuide) sizeInfo() {
	g.section("Размеры типов")
	fmt.Printf("size of bool: %d, int: %d, int64: %d, pointer: %d bytes\n",
		unsafe.Sizeof(false),
		unsafe.Sizeof(int(0)),
		unsafe.Sizeof(int64(0)),
		unsafe.Sizeof(new(int)),
	)
}

func main() {
	g := VarGuide{}
	g.zeroValues()
	g.typedAndShortDecl()
	g.signedUnsigned()
	g.floatsAndComplex()
	g.runeAndByte()
	g.stringsAreImmutable()
	g.arraysAndSlices()
	g.maps()
	g.structs()
	g.pointers()
	g.constants()
	g.typeConversion()
	g.shadowing()
	g.sizeInfo()
}
