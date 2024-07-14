package main

func (m Model) View() string {
	if m.Screen == "INIT_VIEW" {
		return GenerateInitView(m)
	}

	if m.Screen == "HOME_VIEW" {
		return GenerateHomeView(m)
	}

	return ""
}
