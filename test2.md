```
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

func main() {
	p := mpb.New()
	max := 100 * time.Millisecond
	name1 := fmt.Sprintf("[Fct][CSP] Parse ... ")
	bar1 := p.AddBar(int64(50),
		mpb.BarWidth(50),
		mpb.PrependDecorators(
			decor.Name(name1),
			decor.CountersNoUnit("[%d/%d] ", decor.WCSyncWidth),
			decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(
				decor.Percentage(), "done!",
			),
		),
	)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 50; i++ {
		start := time.Now()
		time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
		bar1.SetCurrent(int64(i+1))
		// we need to call DecoratorEwmaUpdate to fulfill ewma decorator's contract
		bar1.DecoratorEwmaUpdate(time.Since(start))
	}

	name2 := fmt.Sprintf("[Fct][CSP] Check ... ")
	bar2 := p.AddBar(int64(77),
		mpb.BarWidth(50),
		mpb.PrependDecorators(
			decor.Name(name2),
			decor.CountersNoUnit("[%d/%d] ", decor.WCSyncWidth),
			decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(
				decor.Percentage(), "done!",
			),
		),
	)

	for i := 0; i < 77; i++ {
		start := time.Now()
		time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
		bar2.SetCurrent(int64(i+1))
		// we need to call DecoratorEwmaUpdate to fulfill ewma decorator's contract
		bar2.DecoratorEwmaUpdate(time.Since(start))
	}
	time.Sleep(1 * time.Minute)
}
```
