package main

import (
	"html/template"
	"log"
	"net/http"

	"gopl/7接口/eval"
)

var calc = template.Must(template.New("calc").Parse(`
<h1>calculater</h1>
<form method="get" action="/">
  <input type="text" name="expr" value="{{.Expr}}"/>
  <input type="submit" value="calc"/>
  <p>{{.Result}}</p>
</form>
`))

type data struct {
	Expr   string
	Result float64
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		d := &data{"", 0.0}

		if len(q["expr"]) > 0 {
			d.Expr = q["expr"][0]
			expr, err := eval.Parse(q["expr"][0])
			if err == nil {
				d.Result = expr.Eval(eval.Env{})
			}
		}
		calc.Execute(w, d)
	}

	http.HandleFunc("/", handler)
	log.Println("Running server in localhost:8000 ...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
