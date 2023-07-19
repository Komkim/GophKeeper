package menu

import (
	"cliclient/internal/cli/creditcard"
	"cliclient/internal/cli/inputs"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
)

// The main view, which just calls the appropriate sub-view
func (m *model) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	switch m.Choice + m.Lvl {

	// стартовая страница
	case 0:
		s, _ = lvl0(m)
	case 1:
		s, _ = lvl0(m)
	case 2:
		m.Choice = 0
		s, _ = lvl0(m)

	// Ввод логина и пароля для входа\регистрациии
	case 10:
		l, _ := tea.NewProgram(inputs.LoginPassword()).Run()
		s = l.View()
		m.Lvl += 10
		m.Choice = 0
	case 11:
		l, _ := tea.NewProgram(inputs.LoginPassword()).Run()
		s = l.View()
		m.Lvl -= 10
		m.Choice = 0

	// Выбор типа информации
	case 20:
		s, _ = lvl2(m)
	case 21:
		s, _ = lvl2(m)
	case 22:
		s, _ = lvl2(m)
	case 23:
		s, _ = lvl2(m)
	case 24:
		m.Choice = 0
		s, _ = lvl2(m)

	//Работа с логином\паролем
	case 30:
		m.Choice = 0
		m.Lvl = 100
	case 100:
		s, _ = logPass(m)
	case 101:
		s, _ = logPass(m)
	case 102:
		s, _ = logPass(m)
	case 103:
		s, _ = logPass(m)
	case 104:
		m.Choice = 0
		s, _ = logPass(m)

	//Работа с картой
	case 31:
		m.Choice = 0
		m.Lvl = 200
	case 200:
		s, _ = card(m)
	case 201:
		s, _ = card(m)
	case 202:
		s, _ = card(m)
	case 203:
		s, _ = card(m)
	case 204:
		m.Choice = 0
		s, _ = card(m)
	//добавить карту
	case 210:
		l, _ := tea.NewProgram(creditcard.InitCard()).Run()
		s = l.View()
		m.Lvl -= 10
		m.Choice = 0

	//Работа с файлами
	case 32:
		m.Choice = 0
		m.Lvl = 300
	case 300:
		s, _ = file(m)
	case 301:
		s, _ = file(m)
	case 302:
		s, _ = file(m)
	case 303:
		s, _ = file(m)
	case 304:
		m.Choice = 0
		s, _ = file(m)

	//Работа с записями
	case 33:
		m.Choice = 0
		m.Lvl = 400
	case 400:
		s, _ = note(m)
	case 401:
		s, _ = note(m)
	case 402:
		s, _ = note(m)
	case 403:
		s, _ = note(m)
	case 404:
		m.Choice = 0
		s, _ = note(m)
	}
	return indent.String("\n"+s+"\n\n", 2)
}

func lvl0(m *model) (string, *model) {
	c := m.Choice

	tpl := "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s",
		checkbox("Войти", c == 0),
		checkbox("Зарегистрироваться", c == 1),
		//tpl,
	)

	return fmt.Sprintf(tpl, choices), m
	//return choices

}

func lvl2(m *model) (string, *model) {
	c := m.Choice

	tpl := "Выберите одно из полей\n\n"
	tpl = "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Логин/пароль", c == 0),
		checkbox("Карта", c == 1),
		checkbox("Файл", c == 2),
		checkbox("Записи", c == 3),
		//tpl,
	)

	return fmt.Sprintf(tpl, choices), m
	//return choices
}

func logPass(m *model) (string, *model) {
	c := m.Choice

	tpl := "Выберите одно из полей\n\n"
	tpl = "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Добавить", c == 0),
		checkbox("Список всех доступных", c == 1),
		checkbox("Получить пароль по логину", c == 2),
		checkbox("Изменить пароль по логину", c == 3),
		//tpl,
	)

	return fmt.Sprintf(tpl, choices), m
	//return choices
}

func card(m *model) (string, *model) {
	c := m.Choice

	tpl := "Выберите одно из полей\n\n"
	tpl = "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Добавить", c == 0),
		checkbox("Список всех доступных", c == 1),
		checkbox("Получить карту по номеру", c == 2),
		checkbox("Изменить карту по номеру", c == 3),
	)

	return fmt.Sprintf(tpl, choices), m
}

func file(m *model) (string, *model) {
	c := m.Choice

	tpl := "Выберите одно из полей\n\n"
	tpl = "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Добавить", c == 0),
		checkbox("Список всех доступных", c == 1),
		checkbox("Получить файл по названию", c == 2),
		checkbox("Изменить название файла", c == 3),
	)

	return fmt.Sprintf(tpl, choices), m
}

func note(m *model) (string, *model) {
	c := m.Choice

	tpl := "Выберите одно из полей\n\n"
	tpl = "%s\n\n"
	tpl += subtle("j/k, up/down, стрелки: вверх/вниз") + dot + subtle("enter: выбрать с: назад") + dot + subtle("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Добавить", c == 0),
		checkbox("Список всех доступных", c == 1),
		checkbox("Получить запись по названию", c == 2),
		checkbox("Изменить содержимое записи", c == 3),
	)

	return fmt.Sprintf(tpl, choices), m
}
