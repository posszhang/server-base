package glog

import (
	"time"
)

func TestWrite(filename string) {

	Create(filename)

	for i := 0; i != 20; i++ {

		a := i
		go func() {

			j := 0
			for {
				j++
				Info("go", a, " count:", j)

				if j%1000 == 0 {
					time.Sleep(1 * time.Millisecond)
				}
			}
		}()
	}

}
