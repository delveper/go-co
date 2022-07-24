package main

//
// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )
//
// func main() {
// 	var urls []string = []string{
// 		"https://google.com",
// 		"https://twitter.com",
// 		"https://tutorialedge.net",
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(len(urls))
//
// 	for _, url := range urls {
// 		go fetch(url, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("process finished")
// }
//
// func fetch(url string, wg *sync.WaitGroup) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(res.Body, "\n\n\n\n")
//
// 	wg.Done()
// }
