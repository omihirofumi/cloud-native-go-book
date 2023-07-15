package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const instanceCount = 1000

const trialDuration = 4 * time.Minute

const bucketWidth = time.Second

var requestBuckets []int

var currentBucketIndex int

var requestEvents chan bool = make(chan bool)

var backoffFunction func() string = withExponentialBackoffAndJitter

var maxStartupDelay = bucketWidth

var startTime = time.Now()

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	bucketCount := int(trialDuration / bucketWidth)
	requestBuckets = make([]int, bucketCount)
	log("Bucket count:  %d\n", bucketCount)

	go catchEvents()

	log("Starting %d backoff processes\n", instanceCount)
	for i := 0; i < instanceCount; i++ {
		go func() {
			delay := time.Duration(rand.Int63n(int64(maxStartupDelay)))
			time.Sleep(delay)
			backoffFunction()
		}()
	}
	log("%d backoff processes started\n", instanceCount)

	for currentBucketIndex = 0; currentBucketIndex < bucketCount; currentBucketIndex++ {
		time.Sleep(bucketWidth)

		i := currentBucketIndex
		if i >= bucketCount {
			i = bucketCount - 1
		}

		log("Bucket %d: %d data points\n", currentBucketIndex, requestBuckets[i])
	}

	sum := 0
	for i := 0; i < bucketCount; i++ {
		sum += requestBuckets[i]
		fmt.Println(requestBuckets[i])
	}

	log("Avg: %d\n", sum/bucketCount)

}

func withNoBackoff() string {
	res, err := sendRequest()
	for err != nil {
		res, err = sendRequest()
	}

	return res
}

func withDelayBackoff() string {
	res, err := sendRequest()
	for err != nil {
		time.Sleep(2000 * time.Millisecond)
		res, err = sendRequest()
	}

	return res
}

func withExponentialBackoff() string {
	res, err := sendRequest()
	base, cap := time.Second, time.Minute

	for backoff := base; err != nil; backoff <<= 1 {
		if backoff > cap {
			backoff = cap
		}
		time.Sleep(backoff)
		res, err = sendRequest()
	}

	return res
}

func withExponentialBackoffAndJitter() string {
	res, err := sendRequest()
	base, cap := time.Second, time.Minute

	for backoff := base; err != nil; backoff <<= 1 {
		if backoff > cap {
			backoff = cap
		}

		jitter := rand.Int63n(int64(backoff * 3))
		sleep := base * time.Duration(jitter)
		time.Sleep(sleep)
		res, err = sendRequest()
	}

	return res
}

func sendRequest() (string, error) {
	delay := time.Duration(rand.Int63n(100)+rand.Int63n(100)) * time.Millisecond

	time.Sleep(delay / 2)
	requestEvents <- true
	time.Sleep(delay / 2)

	return "", errors.New("")
}

func catchEvents() {
	for range requestEvents {
		requestBuckets[currentBucketIndex]++
	}
}

func log(f string, i ...interface{}) {
	since := time.Now().Sub(startTime)
	t := time.Time{}.Add(since)
	tf := t.Format("15:04:05")

	fmt.Printf(tf+" "+f, i...)
}
