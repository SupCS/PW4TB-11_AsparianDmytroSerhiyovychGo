package handlers

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func Example7_2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		Usn, _ := strconv.ParseFloat(r.FormValue("Usn"), 64)
		Sk, _ := strconv.ParseFloat(r.FormValue("Sk"), 64)

		// Константи
		UkPercentage := 10.5 // Відсоток напруги короткого замикання
		Sn := 6.3            // Номінальна потужність трансформатора

		// Перевіряємо, чи всі значення введено коректно
		if Usn > 0 && Sk > 0 {
			// Виконуємо розрахунки
			Xc := (Usn * Usn) / Sk
			Xt := (UkPercentage / 100) * ((Usn * Usn) / Sn)
			Xsum := Xc + Xt
			Ip0 := (Usn * 1000) / (math.Sqrt(3) * Xsum)

			// Відображаємо результати у шаблоні
			tmpl, _ := template.ParseFiles("templates/example7_2.html")
			tmpl.Execute(w, map[string]interface{}{
				"result": true,
				"Xc":     roundFloat(Xc, 2),
				"Xt":     roundFloat(Xt, 2),
				"Xsum":   roundFloat(Xsum, 2),
				"Ip0":    roundFloat(Ip0, 2),
			})
			return
		}
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/example7_2.html")
	tmpl.Execute(w, nil)
}
