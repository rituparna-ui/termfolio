package main

func (m Model) View() string {
	switch m.Screen {
	case BootView:
		return GenerateBootView()
	default:
		return GenerateNotFoundView(&m)
	}
}

func GenerateBootView() string {
	return "Boot View"
}

func GenerateNotFoundView(m *Model) string {
	return CenterTextStyle(m).Render("404 Not Found !")
}
