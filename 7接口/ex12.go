package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mux sync.RWMutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.del)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.New("list").Parse(`
<!DOCTYPE html>
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
	<table>
		{{range $k, $v := .}}
		<tr>
			<td>{{ $k}}</td>
			<td>{{ $v}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
`)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%q already exists!\n", item)
		return
	}

	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %s\n", priceStr)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s created.", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %s\n", priceStr)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s updated.", item, price)
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s deleted", item)
}
