package DashboardController

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("resource/views/dashboard.html")
	if err != nil {
		panic(err)
	}

	_ = temp.Execute(w, nil)
}
