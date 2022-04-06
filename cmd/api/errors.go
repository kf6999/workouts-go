package main

import "net/http"

// Generic helper for logging errors
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// errorResponse() is a generic helper function that returns a JSON response to the client
// with the status code with an interface{} payload.

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse() used for unexpected problem at runtime. Logs detailed error message
// and uses errorResponse() helper to send 500 internal server error response to the client.

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "The server encountered an error and could not complete your request."
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse() used for 404 not found response, returns JSON
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found."
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse() used for 405 method not allowed response, returns JSON

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "The method is not allowed for the requested URL."
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
