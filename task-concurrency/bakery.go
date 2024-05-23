package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

type cake struct {
	bakedBy  int
	bakeTime int
	packedBy int
	packTime int
}

type counter struct {
	k int
	sync.Mutex
}

func (c *counter) dec() int {
	c.Lock()
	defer c.Unlock()
	if c.k == 0 {
		return -1
	}
	c.k--
	return c.k
}

var cakeCounter counter

func RunBakery(n, m, k, t1, t2 int) {
	packCh := make(chan cake, k)
	checkCh := make(chan cake, k)
	doneCh := make(chan struct{})
	cakeCounter.k = k
	var bakeWg sync.WaitGroup
	var packWg sync.WaitGroup
	var checkWg sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fmt.Printf("n: %d, m: %d, k: %d, t1: %d, t2: %d\n", n, m, k, t1, t2)

	bakeWg.Add(n)
	go func() {
		bakeWg.Wait()
		close(packCh)
	}()

	packWg.Add(m)
	go func() {
		packWg.Wait()
		close(checkCh)
		close(doneCh)
	}()

	checkWg.Add(k)

	for i := 0; i < n; i++ {
		go bake(ctx, i+1, t1, packCh, &bakeWg)
	}

	for i := 0; i < m; i++ {
		go pack(i+1, t2, packCh, checkCh, &packWg)
	}

	check(checkCh, doneCh)
}

func bake(ctx context.Context, id, t1 int, packCh chan<- cake, bakeWg *sync.WaitGroup) {
	//fmt.Println("(bake) G", id, "start")
	done := false

	for {
		select {
		case <-ctx.Done():
			done = true
		default:
		}

		currK := cakeCounter.dec()
		//fmt.Println("(bake) G", id, "currK:", currK)
		if currK == -1 || done {
			break
		}
		T1 := id + t1

		time.Sleep(time.Duration(T1) * time.Millisecond)

		c := cake{bakedBy: id, bakeTime: T1}

		packCh <- c
	}
	bakeWg.Done()
	//fmt.Println("(bake) G", id, "end")
}

func pack(id, t2 int, packCh <-chan cake, checkCh chan<- cake, packWg *sync.WaitGroup) {
	//fmt.Println("(pack) G", id, "start")
	for c := range packCh {
		T2 := id + t2

		time.Sleep(time.Duration(T2) * time.Millisecond)

		c.packedBy = id
		c.packTime = T2

		checkCh <- c
	}
	packWg.Done()
	//fmt.Println("(pack) G", id, "end")
}

func check(checkCh <-chan cake, doneCh <-chan struct{}) {
	//fmt.Println("(check) start")
	<-doneCh

	count := 0

	for c := range checkCh {
		fmt.Println(c)
		count++
	}

	fmt.Println("count:", count)
	//fmt.Println("(check) end")
}
