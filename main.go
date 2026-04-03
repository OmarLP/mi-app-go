package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Usuario struct {
	Nombre string
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// ruta principal para mostrar el formulario
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.ExecuteTemplate(w, "index.html", nil)
			return
		}

		// Si es POST, capturamos el nombre
		nombre := r.FormValue("nombre")
		http.Redirect(w, r, "/bienvenida?nombre="+nombre, http.StatusSeeOther)
	})

	// ruta de bienvenida
	http.HandleFunc("/bienvenida", func(w http.ResponseWriter, r *http.Request) {
		nombre := r.URL.Query().Get("nombre")
		data := Usuario{Nombre: nombre}
		tmpl.ExecuteTemplate(w, "bienvenida.html", data)
	})

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
