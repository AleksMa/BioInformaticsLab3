package main

import "fmt"

const GapSymbol = "-"

var Infinity int = -1000

var Open = -10
var Extend = -1

type NeedlemanWunsch struct {
	TopSequence  *Sequence
	LeftSequence *Sequence
	Result       Matrix
	M, I, D      [][]int
	SF           ScoringFunc
	GapValue     int
}

func NewNeedlemanWunsch(first, second *Sequence, sf ScoringFunc, GapValue int) *NeedlemanWunsch {
	nw := &NeedlemanWunsch{
		TopSequence:  second,
		LeftSequence: first,
		Result:       make(Matrix, len(first.Value)+1),
		M: make([][]int, len(first.Value)+1),
		I: make([][]int, len(first.Value)+1),
		D: make([][]int, len(first.Value)+1),
		SF:           sf,
		GapValue:     GapValue,
	}
	// Аллоцируем первую строку
	nw.Result[0] = make(Line, len(second.Value)+1)
	nw.M[0] = make([]int, len(second.Value)+1)
	nw.I[0] = make([]int, len(second.Value)+1)
	nw.D[0] = make([]int, len(second.Value)+1)
	// Обнуляем (0, 0)
	nw.Result[0][0] = &Cell{
		Distance: 0,
		Dir:      NullDirection,
	}
	nw.M[0][0] = 0
	nw.I[0][0] = Infinity
	nw.D[0][0] = Infinity
	// Обнуляем первую строку
	for i := range second.Value {
		nw.Result[0][i+1] = &Cell{
			Distance: GapValue * (i + 1),
			Dir:      LeftDirection,
		}
		nw.M[0][i+1] = Infinity
		nw.I[0][i+1] = Open + i*Extend
		nw.D[0][i+1] = Infinity
	}
	// Аллоцируем оставшиеся строки, зануляем первый столбец
	for i := range first.Value {
		lena := len(second.Value) + 1
		nw.Result[i+1] = make(Line, lena)
		nw.Result[i+1][0] = &Cell{
			Distance: GapValue * (i + 1),
			Dir:      TopDirection,
		}
		nw.M[i+1] = make([]int, lena)
		nw.I[i+1] = make([]int, lena)
		nw.D[i+1] = make([]int, lena)

		nw.M[i+1][0] = Infinity
		nw.I[i+1][0] = Infinity
		nw.D[i+1][0] = Open + i*Extend
	}
	return nw
}

// Функция вывода таблицы для отладки
func (nw *NeedlemanWunsch) Print() {
	fmt.Println("Result:")
	for i := 0; i <= len(nw.LeftSequence.Value); i++ {
		for j := 0; j <= len(nw.TopSequence.Value); j++ {
			fmt.Print(nw.Result[i][j].Distance, ", ", nw.Result[i][j].Dir)
			fmt.Print("   | ")
		}
		fmt.Println()
	}
	fmt.Println("M:")
	for i := 0; i <= len(nw.LeftSequence.Value); i++ {
		for j := 0; j <= len(nw.TopSequence.Value); j++ {
			fmt.Print(nw.M[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("I:")
	for i := 0; i <= len(nw.LeftSequence.Value); i++ {
		for j := 0; j <= len(nw.TopSequence.Value); j++ {
			fmt.Print(nw.I[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("D:")
	for i := 0; i <= len(nw.LeftSequence.Value); i++ {
		for j := 0; j <= len(nw.TopSequence.Value); j++ {
			fmt.Print(nw.D[i][j], " ")
		}
		fmt.Println()
	}
}

func (nw *NeedlemanWunsch) Solve() (string, string, int) {
	// Рекурсивно определяем значения матрицы, одновременно определяя и направления
	nw.determine(len(nw.LeftSequence.Value), len(nw.TopSequence.Value))

	cell := nw.Result[len(nw.LeftSequence.Value)][len(nw.TopSequence.Value)]

	score := cell.Distance
	firstRes, secondRes := "", ""

	// Двигаемся от правой нижней ячейки матрицы к левой верхней, и строим с конца строки-выравнивания
	fp, sp := len(nw.LeftSequence.Value)-1, len(nw.TopSequence.Value)-1
	//gapTop, gapLeft := false, false
	for cell.Dir != NullDirection {
		if cell.Dir == DiagonalDirection {
			firstRes = string(rune(nw.LeftSequence.Value[fp])) + firstRes
			secondRes = string(rune(nw.TopSequence.Value[sp])) + secondRes
			sp--
			fp--
		} else if cell.Dir == TopDirection {
			//gapTop = true
			firstRes = string(rune(nw.LeftSequence.Value[fp])) + firstRes
			secondRes = GapSymbol + secondRes
			fp--
		} else if cell.Dir == LeftDirection {
			//gapLeft = true
			firstRes = GapSymbol + firstRes
			secondRes = string(rune(nw.TopSequence.Value[sp])) + secondRes
			sp--
		} else if cell.Dir == TopFinDirection {
			//gapTop = false
			firstRes = string(rune(nw.LeftSequence.Value[fp])) + firstRes
			secondRes = GapSymbol + secondRes
			fp--
		} else if cell.Dir == LeftFinDirection {
			//gapLeft = false
			firstRes = GapSymbol + firstRes
			secondRes = string(rune(nw.TopSequence.Value[sp])) + secondRes
			sp--
		}
		cell = nw.Result[fp+1][sp+1]
	}

	//fmt.Println(gapTop, gapLeft)

	return firstRes, secondRes, score
}

// Рекурсивное заполнение матрицы
func (nw *NeedlemanWunsch) determine(i, j int) {
	if nw.Result[i][j] != nil {
		return
	}

	if i > 0 && j > 0 {
		nw.determine(i-1,j-1)
		m1 := nw.M[i-1][j-1] + nw.SF[nw.LeftSequence.Value[i-1]][nw.TopSequence.Value[j-1]]
		m2 := nw.I[i-1][j-1] + nw.SF[nw.LeftSequence.Value[i-1]][nw.TopSequence.Value[j-1]]
		m3 := nw.D[i-1][j-1] + nw.SF[nw.LeftSequence.Value[i-1]][nw.TopSequence.Value[j-1]]

		maxVal, _ := max3(m1, m2, m3)
		nw.M[i][j] = maxVal
	}

	gap2, gap3 := false, false

	if j > 0 {
		nw.determine(i,j-1)
		m1 := nw.I[i][j-1] + Extend
		m2 := nw.M[i][j-1] + Open
		m3 := nw.D[i][j-1] + Open

		maxVal, numMax := max3(m1, m2, m3)
		if numMax != 1 {
			gap2 = true
		}
		nw.I[i][j] = maxVal
	}

	if i > 0 {
		nw.determine(i-1,j)
		m1 := nw.D[i-1][j] + Extend
		m2 := nw.M[i-1][j] + Open
		m3 := nw.I[i-1][j] + Open

		maxVal, numMax := max3(m1, m2, m3)
		if numMax != 1 {
			gap3 = true
		}
		nw.D[i][j] = maxVal
	}

	c1, c2, c3 := nw.M[i][j], nw.I[i][j], nw.D[i][j]
	maxVal, maxNum := max3(c2, c3, c1)

	nw.Result[i][j] = &Cell{
		Distance: maxVal,
	}
	curCell := nw.Result[i][j]

	switch maxNum {
	case 3:
		curCell.Dir = DiagonalDirection
	case 1:
		if gap2 {
			curCell.Dir = LeftFinDirection
		} else {
			curCell.Dir = LeftDirection
		}
	case 2:
		if gap3 {
			curCell.Dir = TopFinDirection
		} else {
			curCell.Dir = TopDirection
		}
	}
}

// Достаточно примитивная функция выбора максимума из трех с указанием номера максимального
func max3(a, b, c int) (int, int) {
	if a >= b {
		if a >= c {
			return a, 1
		}
		return c, 3
	}
	if b >= c {
		return b, 2
	}
	return c, 3
}

func max2(a, b int) (int, bool) {
	if a >= b {
		return a, true
	}
	return b, false
}
