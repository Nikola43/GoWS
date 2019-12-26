package test

import (
	"github.com/nikola43/testsocket/utils"
	"log"
	"runtime"
	"sync"
	"testing"
	"time"
)

var curMem uint64

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	//GiB // 1073741824
	//TiB // 1099511627776             (超过了int32的范围)
	//PiB // 1125899906842624
	//EiB // 1152921504606846976
	//ZiB // 1180591620717411303424    (超过了int64的范围)
	//YiB // 1208925819614629174706176
)

func TestGetAll(t *testing.T) {

	url := "http://localhost:8080/api/user"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InNzQGdtYWlsLmNvbSIsImV4cCI6MTU3MDc1MjcyNX0.5SiPPXjU1iSnW_BWFtvzm1yyB6pCPUSBhpQV29Fs8Rw"

	start := time.Now()
	utils.SecureGetRequest(url, token, nil)
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func TestInsert(t *testing.T) {

	//url := "http://localhost:8080/api/user/new"
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InNzQGdtYWlsLmNvbSIsImV4cCI6MTU3MDc1MjcyNX0.5SiPPXjU1iSnW_BWFtvzm1yyB6pCPUSBhpQV29Fs8Rw"

	i := 0
	start := time.Now()
	for i = 0; i < 1000; i++ {
		//user := &models.User{Name: utils.GenerateRandomString(10), Username: utils.GenerateRandomString(10), Password: utils.GenerateRandomString(10)}
		//utils.PostRequest(url, token, user)
		time.Sleep(50 * time.Microsecond)
	}
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

func demoFunc() {
	time.Sleep(time.Duration(10) * time.Millisecond)
}


func TestAntsPool(t *testing.T) {
	pool, _ := ants.NewPool(100000)
	start := time.Now()

		//defer Release()
	var wg sync.WaitGroup
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		_ = pool.Submit(func() {
			_ = 1 * 3 + 4 * 75* 24 / 203 * 5

			//time.Sleep(time.Duration(10) * time.Millisecond)
			wg.Done()
		})
	}
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)

	t.Logf("pool, capacity:%d", pool.Cap())
	t.Logf("pool, running workers number:%d", pool.Running())
	t.Logf("pool, free workers number:%d", pool.Free())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}

func TestGetUser(t *testing.T) {

	url := "http://localhost:8080/api/user/login"
	pool, _ := ants.NewPool(100000)
	start := time.Now()

	//defer Release()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		_ = pool.Submit(func() {
			utils.GetRequest(url, nil)
			time.Sleep(time.Duration(10) * time.Millisecond)
			wg.Done()
		})
	}
	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)

	t.Logf("pool, capacity:%d", pool.Cap())
	t.Logf("pool, running workers number:%d", pool.Running())
	t.Logf("pool, free workers number:%d", pool.Free())

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}
