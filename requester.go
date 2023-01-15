package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	reqCnt, _ := strconv.Atoi(os.Args[1])
	testCnt, _ := strconv.Atoi(os.Args[2])
	slpTime, _ := strconv.Atoi(os.Args[3])
	cnnCnt, _ := strconv.Atoi(os.Args[4])

	fmt.Printf("%v: started with request %v and connection %v\n", time.Now().String(), reqCnt, cnnCnt)
	for i := 0; i < testCnt; i++ {
		time.Sleep(time.Duration(slpTime) * time.Second)
		testCase(reqCnt, cnnCnt)
		runtime.GC()
	}
}
func testCase(reqCnt, cnnCnt int) {
	chId := make(chan int64, reqCnt)
	chErr := make(chan int64, reqCnt)
	var reqs []*reqStruct
	for i := 0; i < reqCnt; i++ {
		r := &reqStruct{
			C:           getHttpClient(cnnCnt),
			Description: RandStringBytes(10),
		}
		reqs = append(reqs, r)
	}

	var wg sync.WaitGroup
	wg.Add(reqCnt)
	// fmt.Printf("%v: started\n", time.Now().String())
	start := time.Now()
	for _, r := range reqs {
		go post(chId, chErr, r, cnnCnt, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	close(chId)
	close(chErr)
	var ids []int64
	for v := range chId {
		ids = append(ids, v)
	}
	var errs []int64
	for iv := range chErr {
		errs = append(errs, iv)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	sort.Slice(errs, func(i, j int) bool {
		return errs[i] < errs[j]
	})
	if len(ids) <= 0 {
		fmt.Printf("nothing\n")
		return
	}
	// if len(ids) != reqCnt {
	//	fmt.Printf("some connection lost(%v, %v)\n", len(ids), reqCnt)
	//	return
	// }
	fmt.Printf("%v: execution time : %v, min: %.3f ms, max: %.3f ms, median: %.3f, average: %.3f, missing: %d, %.2f%%",
		time.Now().String(), elapsed.String(),
		float32(ids[0]/1000), float32(ids[len(ids)-1]/1000), float32(ids[(len(ids)-1)/2]/1000), float32(average(ids)/1000),
		reqCnt-len(ids), float32(reqCnt-len(ids))/float32(reqCnt)*100)
	if len(errs) > 0 {
		fmt.Printf(", retries: %d, max: %d, median: %d, average: %d",
			len(errs), errs[len(errs)-1], errs[len(errs)/2], average(errs))
	}
	fmt.Println()
}

type reqStruct struct {
	C           *http.Client
	Description string
}

type errStruct struct {
	Cnt int
}

func post(chId, chErr chan int64, req *reqStruct, cnnCnt int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	myJson := bytes.NewBuffer([]byte(fmt.Sprintf(`{"description":"%s"}`, req.Description)))
	// fmt.Printf("json body: %+v\n", myJson)
	cnt := 0
	for {
		start := time.Now()
		resp, err := req.C.Post("http://localhost:8080/todo", "application/json", myJson)
		if err != nil {
			if errors.Is(err, syscall.ECONNRESET) && cnt < 10 {
				cnt++
				// req.C = getHttpClient(cnnCnt)
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
				// fmt.Println("reconnecting")
				// time.Sleep(100 * time.Millisecond)
				continue
			}
			fmt.Printf("error: %v\n", err)
			return
		}
		resp.Body.Close()
		elapsed := time.Since(start)
		chId <- elapsed.Microseconds()
		if cnt > 0 {
			chErr <- int64(cnt)
		}
		break
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getHttpClient(cnnCnt int) *http.Client {
	return &http.Client{Timeout: time.Duration(30) * time.Second, Transport: &http.Transport{MaxConnsPerHost: cnnCnt}}
}
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func average(ids []int64) int64 {
	a := int64(0)
	for _, v := range ids {
		a = a + v
	}
	return a / int64(len(ids))
}
func executionTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("execution time: %v", elapsed.String())
}
