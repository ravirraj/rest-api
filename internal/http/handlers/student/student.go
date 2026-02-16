package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"

	// "log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ravirraj/rest-api/internal/storage"
	types "github.com/ravirraj/rest-api/internal/type"
	"github.com/ravirraj/rest-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		//request validation

		if err := validator.New().Struct(student); err != nil {
			validatesErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validatesErrs))
			return

		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("User created successfully", slog.String("userid", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
		}
		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}
