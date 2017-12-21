package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var t = template.Must(template.New("list").Parse(`
		<html>
			<body>
				<table>
					<tr style="text-align: left">
						<th>item</th>
						<th>price</th>
					</tr>
					{{range $item, $price := .Database}}
					<tr>
						<td>{{$item}}</td>
						<td>{{$price}}</td>
					</tr>
					{{end}}
				</table>
			</body>
		</html>
	`))
	err := t.Execute(w, struct {
		Database map[string]dollars
	}{
		Database: db,
	})
	if err != nil {
		panic(err)
	}
}
