package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// func main() {
// 	var s, sep string
// 	for i := 1; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	s, sep := "", ""
// 	for idx, arg := range os.Args[1:] {
// 		println(arg, idx)
// 		// s = sep + arg
// 		// sep = ""
// 	}
// 	fmt.Println(s)
// 	fmt.Println(sep)
// }

// func main() {
// 	// counts := make(map[string]int)
// 	// input := bufio.NewScanner(os.Stdin)
// 	// for input.Scan() {
// 	// 	counts[input.Text()]++

// 	// }

// 	// for line, n := range counts {
// 	// 	if n > 1 {
// 	// 		fmt.Println("%d\t%s\n", n, line)
// 	// 	}
// 	// }
// 	print(os.Args[1])
// }

// func main() {
// 	for _, url := range os.Args[1:] {
// 		hasPrefix := strings.HasPrefix(url, "http://")
// 		if !hasPrefix {
// 			url = "http://" + url
// 		}
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Printf("%v", err)
// 			os.Exit(1)
// 		}
// 		dst := os.Stdout
// 		b, err := io.Copy(dst, resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Printf("%v", err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 		fmt.Println(resp.StatusCode)
// 	}
// }

// func main() {
// 	fmt.Println("HELLO")
// 	greeter()
// }

// func greeter() {
// 	fmt.Println("Hey there mode users")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", serveHome).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":3001", r))
// }

//	func serveHome(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("<h1>This is server side rendered page</h1>"))
//	}
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
