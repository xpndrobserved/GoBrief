package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `
 ________  ________  ________  ________  ___  _______   ________ ___
|\   ____\|\   __  \|\   __  \|\   __  \|\  \|\  ___ \ |\  _____\\  \
\ \  \___|\ \  \|\  \ \  \|\ /\ \  \|\  \ \  \ \   __/|\ \  \__/\ \  \
 \ \  \  __\ \  \\\  \ \   __  \ \   _  _\ \  \ \  \_|/_\ \   __\\ \  \
  \ \  \|\  \ \  \\\  \ \  \|\  \ \  \\  \\ \  \ \  \_|\ \ \  \_| \ \__\
   \ \_______\ \_______\ \_______\ \__\\ _\\ \__\ \_______\ \__\   \|__|
    \|_______|\|_______|\|_______|\|__|\|__|\|__|\|_______|\|__|       ___
                                                                      |\__\
                                                                      \|__| 

`

const (
	subtitle   = `GoBrief! - The Pilot's Terminal Weather Briefing System`
	navigation = `[yellow]Ctrl-N[-]: Next Page [yellow]Ctrl-P[-]: Previous Page [yellow]Ctrl-C[-]: Exit`
	mouse      = `(or use your mouse)`
)

// Function adopted from the examples in tview.

func Cover(nextSlide func()) (title string, content tview.Primitive) {
	// What's the size of the logo?
	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText(version, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mouse, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Home", flex
}
