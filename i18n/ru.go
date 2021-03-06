package i18n

var RuLocal = map[string]string{
	"start":              "[BOT] Добро пожаловать в чат! Твой ник - '%s'. Для справки нажми /help.",
	"start_new":          "[BOT] Привет! Твой ник - '%s'.\nДобро пожаловать в уютный чатик :з\nЗдесь только тёплое общение.\nЧтобы выделяться, меняй ник, для этого надо написать /nick и через пробел свой никнейм\nПример /nick Петушок\n\nСпасибо, что присоединился к нам :з",
	"new_user":           "[BOT] '%s' вошел в чат. Он новенький!",
	"you_are_banned":     "[BOT] Ты забанен.",
	"already_in_chat":    "[BOT] Ты уже в чате.",
	"stop":               "[BOT] Ты вышел из чата. Нажми /start, чтобы вернуться.",
	"leave_chat":         "[BOT] '%s' вышел из чата.",
	"entry_user":         "[BOT] '%s' вошел в чат.",
	"not_in_chat":        "[BOT] Ты не в чате. Чтобы войти, нажми /start",
	"audio_from_user":    "%s отправил(а) аудио.",
	"photo_from_user":    "%s отправил(а) фото.",
	"document_from_user": "%s отправил(а) файл.",
	"video_from_user":    "%s отправил(а) видео.",
	"voice_from_user":    "%s отправил(а) войс.",
	"list_item":          "#%s '%s' был(а) %s назад\n",
	"list_item_online":   "#%s '%s' онлайн\n",
	"list":               "[BOT] Пользователи в чате:\n%s",
	"short_nickname":     "[BOT] Слишком короткий ник.",
	"nickname_exists":    "[BOT] Извини, но этот ник уже занят.",
	"new_nick":           "[BOT] Твой новый ник - '%s'.",
	"new_user_nick":      "[BOT] '%s' сменил ник на '%s'.",
	"help":               "[BOT] Действующие команды:\n/start - войти в чат\n/stop - выйти из чата\n/nick {новый ник} - сменить ник\n/list - список пользователей\n/info {id пользователя} - информация о пользователе (без id пользователя - информация о себе)\n/me {текст} - написать от 3-го лица\n% {текст} - написать от 3-го лица",
	"info":               "[BOT] Информация о '%s':\nВ чате уже %s\nПоследний онлайн: %s",
	"info_online":        "[BOT] Информация о '%s':\nВ чате уже %s\nПоследний онлайн: онлайн",
	"user_not_found":     "[BOT] Пользователь не найден.",
	"me":                 "'%s' <i>%s</i>",
	"leave_chat_with_ban": "[BOT] '%s' забанил бота.",
}
