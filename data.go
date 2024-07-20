package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func GenerateExitView(m *Model) string {
	exitStyle := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 8).
		Height(m.Window.Height - 5).
		AlignVertical(lipgloss.Center).
		AlignHorizontal(lipgloss.Center)

	thankYou := "Thank you for visiting my Termfolio !!!\n"
	thankYou = lipgloss.NewStyle().Blink(true).Foreground(lipgloss.Color("63")).Render(thankYou)

	openSourceBold := lipgloss.NewStyle().Bold(true).Render("Open Source")
	github := "It's " + openSourceBold + "! check it out at\n"
	github += "https://github.com/rituparna-ui/termfolio"

	githubRender := lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#FF00FF")).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(github)

	confirmation := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 8).
		Foreground(lipgloss.Color("#FF0000")).
		Render("Are you sure you want to exit ? (y/n)")

	return exitStyle.Render(
		lipgloss.JoinVertical(lipgloss.Center, thankYou, githubRender, "\n", confirmation),
	)
}

var contactInfo = []string{
	"Portfolio", "https://rituparnawarwatkar.com/",
	"LinkedIn", "https://www.linkedin.com/in/rituparna-warwatkar/",
	"GitHub", "https://github.com/rituparna-ui/",
	"WhatsApp", "+91 77985 16764",
	"Email", "rwarwatkar@gmail.com",
	"Instagram", "https://www.instagram.com/rituparna_warwatkar/",
}

func GenerateContactView(m *Model) string {
	contactStyle := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 8).
		Height(m.Window.Height - 5)

	contact := "Contact Me"
	contact = lipgloss.
		NewStyle().
		Bold(true).
		Background(lipgloss.Color("63")).
		PaddingLeft(1).
		PaddingRight(1).
		Render(contact)

	infoList := ""

	for i := 0; i < len(contactInfo); i += 2 {
		info := contactInfo[i]
		link := contactInfo[i+1]

		infoLink := lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.NewStyle().Bold(true).Render(info),
			lipgloss.NewStyle().Faint(true).Render(link),
		)

		infoList += infoLink + "\n\n"
	}

	return contactStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			contact,
			"\n\n",
			infoList,
		),
	)
}

func GetFiglet() string {
	figlet := ""
	figlet += " ____  \n"
	figlet += "|  _ \\ \n"
	figlet += "| |_) |\n"
	figlet += "|  _ < \n"
	figlet += "|_| \\_\\ ITUPARNA WARWATKAR\n"

	return figlet
}

func GenerateAboutView(m *Model) string {
	figlet := GetFiglet()
	building := lipgloss.
		NewStyle().
		Width(m.Window.Width*3/4 - 10).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderTop(true).
		BorderBottom(true).
		BorderLeft(false).
		BorderRight(false).
		PaddingLeft(1).
		Render("🚀 Building Veritus.ai 🌎 Japan")

	tech := []string{
		"a FullStack Developer.",
		"a Flutter Developer.",
		"a ME(F/A/R)N Stack Developer.",
		"Flutter/Angular/React.",
		"a linux enthusiast.",
		"also learning GenAI.",
	}

	stack := lipgloss.NewStyle().Bold(true).Render("I am\n")

	for i, t := range tech {
		stack += fmt.Sprintf("\n%d. %s", i+1, t)
	}

	visitorCount := fmt.Sprintf("You are visitor number: %d", m.Visitors)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		figlet,
		"A high-performing engineer, passionate about tech and startups !",
		building,
		stack,
		"\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Bold(true).Render(visitorCount),
	)
}

func GetNavItems() []string {
	return []string{"About", "Projects", "Contact Me", "Exit"}
}

func GenerateTabView(m *Model) string {
	switch cursor := m.HomeView.LeftPane.cursor; cursor {
	case 0:
		return GenerateAboutView(m)
	case 1:
		return GenerateProjectsView(m)
	case 2:
		return GenerateContactView(m)
	case 3:
		return GenerateExitView(m)
	}
	return "404 | Not Found"
}

type Project struct {
	Name        string
	Description string
	Link        string
	Tags        []string
}

var projects = []Project{
	{
		Name:        "Termfolio",
		Description: "A terminal based portfolio app, served over SSH",
		Link:        "https://github.com/rituparna-ui/termfolio",
		Tags:        []string{"Go", "Bubbletea", "Lipgloss"},
	},
	{
		Name:        "HealthCare Portal",
		Description: "One-stop platform for better virtual healthcare in a hassle-free manner and eliminates any stigma associated with conventional treatment.",
		Link:        "https://github.com/rituparna-ui/health-care-portal",
		Tags:        []string{"Express", "NodeJS", "MongoDB", "React", "FastAPI"},
	},
	{
		Name:        "Eligram-X",
		Description: "A robust social media platform built with the MEAN stack, featuring advanced security measures like OAuth, session management, two-factor authentication along with a fully featured admin panel",
		Link:        "https://github.com/rituparna-ui/eligram-x",
		Tags:        []string{"MongoDB", "Express", "Angular", "NodeJS", "Firebase", "Redis"},
	},
	{
		Name:        "IIIT-Guwahati I-Card Portal",
		Description: "This application processes around 500 ID card requests annually, benefiting the Indian Institute of Information Technology Guwahati.",
		Link:        "https://github.com/rituparna-ui/iiitg-icard",
		Tags:        []string{"MongoDB", "Express", "NodeJS"},
	},
	{
		Name:        "Multi-Threaded Kernel (On-going)",
		Description: "A 32-bit multi-threaded kernel with FAT16 filesystem, utilizing Intel's memory protection mechanisms to safeguard the kernel from user program damage. Memory virtualization, virtual filesystem layer, FAT16 filesystem driver etc.",
		Link:        "https://github.com/rituparna-ui/kernel",
		Tags:        []string{"C", "Assembly"},
	},
}

func GenerateProjectsList(m *Model) string {
	projectsList := ""

	for i, project := range projects {
		title := lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("%d. %s", i+1, project.Name))
		description := lipgloss.NewStyle().Faint(true).Render(project.Description)
		link := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render("GitHub: " + project.Link)
		tags := lipgloss.NewStyle().Foreground(lipgloss.Color("63")).Render("Tags: " + strings.Join(project.Tags[:], ", "))

		projectsList = lipgloss.JoinVertical(
			lipgloss.Left,
			projectsList,
			title,
			description,
			tags,
			link,
			"\n",
		)
	}

	return projectsList
}

func GenerateProjectsView(m *Model) string {
	projects := lipgloss.
		NewStyle().
		Bold(true).
		Background(lipgloss.Color("63")).
		PaddingLeft(1).
		PaddingRight(1).
		Render("Projects")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		projects,
		GenerateProjectsList(m),
	)
}
