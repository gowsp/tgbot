package tgbot

// Message is bot command
func (m *Message) IsBotCmd() bool {
	for _, v := range m.Entities {
		if v.Type == "bot_command" {
			return true
		}

	}
	return false
}
