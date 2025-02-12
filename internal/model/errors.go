package model

import "errors"

var (
	ErrInvalidRequest      = errors.New("неверный запрос")
	ErrUnauthorized        = errors.New("неавторизован")
	ErrInternalServerError = errors.New("внутренняя ошибка сервера")

	ErrInsufficientCoins  = errors.New("недостаточно монет")
	ErrUserNotFound       = errors.New("пользователь не найден")
	ErrItemNotFound       = errors.New("товар не найден")
	ErrInvalidTransaction = errors.New("недопустимая транзакция")
	ErrSelfTransaction    = errors.New("нельзя отправить монеты самому себе")
	ErrNegativeAmount     = errors.New("количество монет не может быть отрицательным")
	ErrInvalidCredentials = errors.New("неверные учетные данные")
)
