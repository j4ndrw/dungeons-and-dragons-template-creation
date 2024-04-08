package forms

import "github.com/charmbracelet/huh"

func New(groups ...*huh.Group) *huh.Form {
	return huh.
		NewForm(groups...).
		WithTheme(huh.ThemeCatppuccin())
}
