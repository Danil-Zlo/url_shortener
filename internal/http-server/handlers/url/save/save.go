package save

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/middleware"
	render "github.com/go-chi/render"
	validator "github.com/go-playground/validator"

	resp "github.com/Danil-Zlo/url_shortener/internal/lib/api/response"
	sl "github.com/Danil-Zlo/url_shortener/internal/lib/logger/slog"
	random "github.com/Danil-Zlo/url_shortener/internal/lib/random"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempry"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

// TODO: move to config
const aliasLength = 6

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.Error("invalid request"))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}
	}
}
