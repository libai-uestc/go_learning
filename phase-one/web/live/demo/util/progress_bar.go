package util

import "fmt"

type Bar struct {
	persent int64
	cur     int64
	total   int64
	rate    string
	graph   string
}

func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.persent = bar.getPercent()
	for i := 0; i < int(bar.persent); i += 2 {
		bar.rate += bar.graph
	}
}

func (bar *Bar) NewOptionWithGraph(start, total int64, graph string) {
	bar.graph = graph
	bar.NewOption(start, total)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.persent
	bar.persent = bar.getPercent()
	if bar.persent != last && bar.persent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.persent, bar.cur, bar.total) // \r将光标移动到当前行的开头
}

func (bar *Bar) Finish() {
	fmt.Println()
}

// go get github.com/gin
// go get lib2
// go get lib3

// go mod tidy
