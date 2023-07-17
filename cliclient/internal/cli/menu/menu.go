package menu

import (
	"fmt"
	"github.com/muesli/reflow/indent"
	"strconv"
)

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	switch m.Lvl {
	case 0:
		s = lvl1(m)
	case 1:
	case 2:

	}
	return indent.String("\n"+s+"\n\n", 2)
}

func lvl1(m model) string {
	c := m.Choice

	tpl := "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox("Войти", c == 0),
		checkbox("Зарегистрироваться", c == 1),
		//tpl,
	)

	return fmt.Sprintf(tpl, choices)
	//return choices

}

// The second view, after a task has been chosen
func chosenView(m model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("Очень хорошо")
		//case 1:
		//	msg = fmt.Sprintf("A trip to the market?\n\nOkay, then we should install %s and %s...", keyword("marketkit"), keyword("libshopping"))
		//case 2:
		//	msg = fmt.Sprintf("Reading time?\n\nOkay, cool, then we’ll need a library. Yes, an %s.", keyword("actual library"))
		//default:
		//	msg = fmt.Sprintf("It’s always good to see friends.\n\nFetching %s and %s...", keyword("social-skills"), keyword("conversationutils"))
	}

	label := "Downloading..."
	if m.Loaded {
		label = fmt.Sprintf("Downloaded. Exiting in %s seconds...", colorFg(strconv.Itoa(m.Ticks), "79"))
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.Progress) + "%"
}

// Sub-views

// The first view, where you're choosing a task
func choicesView(m model) string {
	//c := m.Choice
	//
	//header := "Выбери что хочешь\n\n"
	//footer := subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")
	//
	//choices := fmt.Sprintf(
	//	"%s\n%s\n\n",
	//	checkbox("Войти", c == 0),
	//	checkbox("Зарегистрироваться", c == 1),
	//)
	//return fmt.Sprintf(header, choices, footer)

	c := m.Choice

	tpl := "What to do today?\n\n"
	tpl += "%s\n\n"
	tpl += "Program quits in %s seconds\n\n"
	tpl += subtle("j/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Plant carrots", c == 0),
		checkbox("Go to the market", c == 1),
		checkbox("Read something", c == 2),
		checkbox("See friends", c == 3),
	)

	return fmt.Sprintf(tpl, choices, colorFg(strconv.Itoa(m.Ticks), "79"))
}
