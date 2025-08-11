package slogdiscard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHadler())
}

type DiscardHadler struct{}

func NewDiscardHadler() *DiscardHadler {
	return &DiscardHadler{}
}

func (h *DiscardHadler) Handle(_ context.Context, _ slog.Record) error {
	// Игнорируем запись журнала
	return nil
}

func (h *DiscardHadler) WithAttrs(_ []slog.Attr) slog.Handler {
	// Возвращает тот же обработчик, так как нет атрибутов для сохранения
	return h
}

func (h *DiscardHadler) WithGroup(_ string) slog.Handler {
	// Возвращает тот же обработчик, так как нет группы ддля сохранения
	return h
}

func (h *DiscardHadler) Enabled(_ context.Context, _ slog.Level) bool {
	// Всегда возвращает false, так как запись в журнал игнорируется
	return false
}
