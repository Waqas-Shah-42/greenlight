package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"system_info": data}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered while encoding json", http.StatusInternalServerError)
		return
	}

	// js := `{"status": available, environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(js))

	// fmt.Fprintln(w,"states: available")
	// fmt.Fprintf(w,"environment: %s\n", app.config.env)
	// fmt.Fprintf(w,"version: %s\n", version)
}
