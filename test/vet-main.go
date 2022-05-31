package test

// import "sync"

// func AssignmentLock() {
// 	var lock sync.Mutex

// 	l := lock
// 	l.Lock()
// 	defer l.Unlock()
// }

// func LoopClosure() {
// 	var wg sync.WaitGroup
// 	for _, v := range []int{0, 1, 2, 3} {
// 		wg.Add(1)
// 		go func() {
// 			fmt.Println(v)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }
