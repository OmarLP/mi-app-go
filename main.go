package main

import (
	"fmt"
	"net/http"
	"os"
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

	// buscamos si el servidor nos dio un puerto especifico
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // si es local usamos 8080 por defecto
	}

	fmt.Println("Servidor iniciado en http://localhost:8080")
	fmt.Println("Servidor iniciado en el puerto " + port)

	// Ahora usamos la variable PORT
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
