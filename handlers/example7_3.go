package handlers

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Обробник для прикладу 7.3
func Example7_3Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		Ux_max, _ := strconv.ParseFloat(r.FormValue("Ux_max"), 64)
		Snom_t, _ := strconv.ParseFloat(r.FormValue("Snom_t"), 64)
		Rc_n, _ := strconv.ParseFloat(r.FormValue("Rc_n"), 64)
		Xc_n, _ := strconv.ParseFloat(r.FormValue("Xc_n"), 64)
		Rc_min, _ := strconv.ParseFloat(r.FormValue("Rc_min"), 64)
		Xc_min, _ := strconv.ParseFloat(r.FormValue("Xc_min"), 64)
		Ub_n, _ := strconv.ParseFloat(r.FormValue("Ub_n"), 64)
		Unn, _ := strconv.ParseFloat(r.FormValue("Unn"), 64)

		if Ux_max <= 0 || Snom_t <= 0 || Rc_n <= 0 || Xc_n <= 0 || Rc_min <= 0 || Xc_min <= 0 || Ub_n <= 0 || Unn <= 0 {
			tmpl, _ := template.ParseFiles("templates/example7_3.html")
			tmpl.Execute(w, map[string]interface{}{
				"error": "Будь ласка, введіть коректні значення для всіх полів.",
			})
			return
		}

		// Розрахунки
		Xt := (Ux_max * math.Pow(Ub_n, 2)) / (100 * Snom_t)

		// Розрахунок для нормального режиму
		Xsh_n := Xc_n + Xt
		Zsh_n := math.Hypot(Rc_n, Xsh_n)
		Ip3_n := (Ub_n * 1000) / (math.Sqrt(3) * Zsh_n)
		Ip2_n := Ip3_n * math.Sqrt(3) / 2

		// Розрахунок для мінімального режиму
		Xsh_min := Xc_min + Xt
		Zsh_min := math.Hypot(Rc_min, Xsh_min)
		Ip3_min := (Ub_n * 1000) / (math.Sqrt(3) * Zsh_min)
		Ip2_min := Ip3_min * math.Sqrt(3) / 2

		// Коефіцієнт приведення та опори
		kpr := math.Pow(Unn, 2) / math.Pow(Ub_n, 2)
		Rsh_n_adj := Rc_n * kpr
		Xsh_n_adj := Xsh_n * kpr
		Zsh_n_adj := math.Hypot(Rsh_n_adj, Xsh_n_adj)

		Rsh_min_adj := Rc_min * kpr
		Xsh_min_adj := Xsh_min * kpr
		Zsh_min_adj := math.Hypot(Rsh_min_adj, Xsh_min_adj)

		// Дійсні струми
		I3_sh_n := (Unn * 1000) / (math.Sqrt(3) * Zsh_n_adj)
		I2_sh_n := I3_sh_n * math.Sqrt(3) / 2
		I3_sh_min := (Unn * 1000) / (math.Sqrt(3) * Zsh_min_adj)
		I2_sh_min := I3_sh_min * math.Sqrt(3) / 2

		// Довжина і опори
		Ln := 0.2 + 0.35 + 0.2 + 0.6 + 2 + 2.55 + 3.37 + 3.1
		R0 := 0.64
		X0 := 0.363

		Rn := Ln * R0
		Xn := Ln * X0

		// Опори в точці 10
		Rx_n := Rn + Rsh_n_adj
		Xx_n := Xn + Xsh_n_adj
		Zx_n := math.Hypot(Rx_n, Xx_n)

		Rx_min := Rn + Rsh_min_adj
		Xx_min := Xn + Xsh_min_adj
		Zx_min := math.Hypot(Rx_min, Xx_min)

		// Струми трифазного і двофазного КЗ в точці 10
		I3_t10_n := (Unn * 1000) / (math.Sqrt(3) * Zx_n)
		I2_t10_n := I3_t10_n * math.Sqrt(3) / 2
		I3_t10_min := (Unn * 1000) / (math.Sqrt(3) * Zx_min)
		I2_t10_min := I3_t10_min * math.Sqrt(3) / 2

		tmpl, _ := template.ParseFiles("templates/example7_3.html")
		tmpl.Execute(w, map[string]interface{}{
			"Xt":          roundFloat(Xt, 2),
			"Xsh_n":       roundFloat(Xsh_n, 2),
			"Zsh_n":       roundFloat(Zsh_n, 2),
			"Ip3_n":       roundFloat(Ip3_n, 2),
			"Ip2_n":       roundFloat(Ip2_n, 2),
			"Xsh_min":     roundFloat(Xsh_min, 2),
			"Zsh_min":     roundFloat(Zsh_min, 2),
			"Ip3_min":     roundFloat(Ip3_min, 2),
			"Ip2_min":     roundFloat(Ip2_min, 2),
			"kpr":         roundFloat(kpr, 3),
			"Rsh_n_adj":   roundFloat(Rsh_n_adj, 2),
			"Xsh_n_adj":   roundFloat(Xsh_n_adj, 2),
			"Zsh_n_adj":   roundFloat(Zsh_n_adj, 2),
			"Rsh_min_adj": roundFloat(Rsh_min_adj, 2),
			"Xsh_min_adj": roundFloat(Xsh_min_adj, 2),
			"Zsh_min_adj": roundFloat(Zsh_min_adj, 2),
			"I3_sh_n":     roundFloat(I3_sh_n, 2),
			"I2_sh_n":     roundFloat(I2_sh_n, 2),
			"I3_sh_min":   roundFloat(I3_sh_min, 2),
			"I2_sh_min":   roundFloat(I2_sh_min, 2),
			"Ln":          roundFloat(Ln, 2),
			"Rn":          roundFloat(Rn, 2),
			"Xn":          roundFloat(Xn, 2),
			"Rx_n":        roundFloat(Rx_n, 2),
			"Xx_n":        roundFloat(Rx_n, 2),
			"Zx_n":        roundFloat(Rx_n, 2),
			"Rx_min":      roundFloat(Rx_min, 2),
			"Xx_min":      roundFloat(Xx_min, 2),
			"Zx_min":      roundFloat(Zx_min, 2),
			"I3_t10_n":    roundFloat(I3_t10_n, 2),
			"I2_t10_n":    roundFloat(I2_t10_n, 2),
			"I3_t10_min":  roundFloat(I3_t10_min, 2),
			"I2_t10_min":  roundFloat(I2_t10_min, 2),
		})
		return
	}

	tmpl, _ := template.ParseFiles("templates/example7_3.html")
	tmpl.Execute(w, nil)
}
