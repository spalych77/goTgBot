package main

const (
	EMOJI_HALO         = "\U0001F607" // 😇
	EMOJI_ANGRY        = "\U0001F624" // 😤
	EMOJI_WOW          = "\U0001F604" // 😄
	EMOJI_SMILE        = "\U0001F642" // 🙂
	EMOJI_REP          = "\U0001F91D" // 🤝
	EMOJI_SUNGLASSES   = "\U0001F60E" // 😎
	EMOJI_DOWN         = "\U0001F44E" // 👎
	EMOJI_BICEPS       = "\U0001F4AA" // 💪
	EMOJI_SAD          = "\U0001F615" // 😕
	EMOJI_BUTTON_START = "\U000025B6" // ▶
	EMOJI_BUTTON_END   = "\U000025C0" // ◀

	BUTTON_TEXT_PRINT_INTRO       = EMOJI_BUTTON_START + "Показать интро" + EMOJI_BUTTON_END
	BUTTON_TEXT_SKIP_INTRO        = EMOJI_BUTTON_START + "Скипнуть интро" + EMOJI_BUTTON_END
	BUTTON_TEXT_BALANCE           = EMOJI_BUTTON_START + "Очки респекта" + EMOJI_BUTTON_END
	BUTTON_TEXT_USEFUL_ACTIVITIES = EMOJI_BUTTON_START + "Действия" + EMOJI_BUTTON_END
	BUTTON_TEXT_REWARDS           = EMOJI_BUTTON_START + "Дроп" + EMOJI_BUTTON_END
	BUTTON_TEXT_PRINT_MENU        = EMOJI_BUTTON_START + "ДОМОЙ" + EMOJI_BUTTON_END

	BUTTON_CODE_PRINT_INTRO       = "print_intro"
	BUTTON_CODE_SKIP_INTRO        = "skip_intro"
	BUTTON_CODE_BALANCE           = "show_balance"
	BUTTON_CODE_USEFUL_ACTIVITIES = "show_useful_activities"
	BUTTON_CODE_REWARDS           = "show_rewards"
	BUTTON_CODE_PRINT_MENU        = "print_menu"

	TOKEN_NAME_IN_OS             = "KTSstudent_bot "
	UPDATE_CONFIG_TIMEOUT        = 60
	MAX_USER_COINS        uint16 = 500
)
