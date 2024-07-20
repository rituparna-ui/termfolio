package main

func (m Model) View() string {
	if m.Screen == EntryScreen {
		return GenerateEntryView(&m)
	}

	if m.Screen == HomeScreen {
		return GenerateHomeView(&m)
	}

	return "Loading..."
}
