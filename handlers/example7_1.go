package handlers

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Функція для округлення чисел до двох знаків після коми
func roundFloat(value float64, precision int) string {
	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, value)
}

func Example7_1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		Sm, _ := strconv.ParseFloat(r.FormValue("Sm"), 64)
		Ik, _ := strconv.ParseFloat(r.FormValue("Ik"), 64)
		Tf, _ := strconv.ParseFloat(r.FormValue("Tf"), 64)
		Unom, _ := strconv.ParseFloat(r.FormValue("Unom"), 64)
		Tm, _ := strconv.ParseFloat(r.FormValue("Tm"), 64)

		// Константи
		Jek := 1.4 // Густина струму
		Ct := 92.0 // Константа для термічної стійкості

		// Перевіряємо, чи всі значення введено коректно
		if Sm > 0 && Ik > 0 && Tf > 0 && Unom > 0 && Tm > 0 {
			// Виконуємо розрахунки
			Im := Sm / (2 * math.Sqrt(3) * Unom)
			ImPa := 2 * Im
			Sek := Im / Jek
			Smin := (Ik * math.Sqrt(Tf)) / Ct

			// Відображаємо результати у шаблоні
			tmpl, _ := template.ParseFiles("templates/example7_1.html")
			tmpl.Execute(w, map[string]interface{}{
				"result": true,
				"Im":     roundFloat(Im, 2),
				"ImPa":   roundFloat(ImPa, 2),
				"Sek":    roundFloat(Sek, 2),
				"Smin":   roundFloat(Smin, 2),
			})
			return
		}
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/example7_1.html")
	tmpl.Execute(w, nil)
}
