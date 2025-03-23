package main

import "github.com/charmbracelet/lipgloss"

const AVeryLongString = `Rituparna 1
Rituparna 2
Rituparna 3
Rituparna 4
Rituparna 5
Rituparna 6
Rituparna 7
Rituparna 8
Rituparna 9
Rituparna 10
Rituparna 11
Rituparna 12
Rituparna 13
Rituparna 14
Rituparna 15
Rituparna 16
Rituparna 17
Rituparna 18
Rituparna 19
Rituparna 20
Rituparna 21
Rituparna 22
Rituparna 23
Rituparna 24
Rituparna 25
Rituparna 26
Rituparna 27
Rituparna 28
Rituparna 29
Rituparna 30
Rituparna 31
Rituparna 32
Rituparna 33
Rituparna 34
Rituparna 35
Rituparna 36
Rituparna 37
Rituparna 38
Rituparna 39
Rituparna 40
Rituparna 41
Rituparna 42
Rituparna 43
Rituparna 44
Rituparna 45
Rituparna 46
Rituparna 47
Rituparna 48
Rituparna 49
Rituparna 50
Rituparna 51
Rituparna 52
Rituparna 53
Rituparna 54
Rituparna 55
Rituparna 56
Rituparna 57
Rituparna 58
Rituparna 59
Rituparna 60
Rituparna 61
Rituparna 62
Rituparna 63
Rituparna 64
Rituparna 65
Rituparna 66
Rituparna 67
Rituparna 68
Rituparna 69
Rituparna 70
Rituparna 71
Rituparna 72
Rituparna 73
Rituparna 74
Rituparna 75
Rituparna 76
Rituparna 77
Rituparna 78
Rituparna 79
Rituparna 80
Rituparna 81
Rituparna 82
Rituparna 83
Rituparna 84
Rituparna 85
Rituparna 86
Rituparna 87
Rituparna 88
Rituparna 89
Rituparna 90
Rituparna 91
Rituparna 92
Rituparna 93
Rituparna 94
Rituparna 95
Rituparna 96
Rituparna 97
Rituparna 98
Rituparna 99
Rituparna 100`

const AboutPageTextHeader = ` ____
|  _ \
| |_) |
|  _ <
|_| \_\ ituparna Warwatkar

Yet another noob in this professional world !`

func getAboutPageTextBody(m *Model) string {
	return lipgloss.
		NewStyle().
		Border(lipgloss.ThickBorder(), true, false).
		Width(m.Dim.Width*3/4 - 8).
		Render("🚀 Building EC2 @ Amazon Web Services 🌍 Berlin, Germany")
}

const aboutPageFooter = `I am

1. a Backend Developer
2. a FullStack Developer
3. a linux enthusiast
4. a ME(F/A/R/V)N Stack Developer
5. a N(e/u)xt Developer`

func getAboutPageText(m *Model) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		AboutPageTextHeader,
		getAboutPageTextBody(m),
		aboutPageFooter,
	)
}
