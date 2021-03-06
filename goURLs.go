// goURLs
package main

import "fmt"
import "net/http"

func main() {
	urls := []string{"http://qchat.cn", "http://baidu.com"}
	done := make(chan string)
	for _, u := range urls {
		go func(u string) {
			resp, err := http.Get(u)

			if err != nil {
				done <- u + err.Error() // + http.Cookie.Name

			} else {
				done <- u + resp.Status
			}
		}(u)
	}
	for _ = range urls {
		fmt.Println(<-done)
	}
}
