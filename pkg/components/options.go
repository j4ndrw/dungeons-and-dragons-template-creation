package components

import "github.com/j4ndrw/dungeons-and-dragons-template-creation/pkg/validators"

type Option struct {
	Validator   *validators.ValidatorType
	Suggestions *[]string
	Placeholder *string
}

func WithValidator(v validators.ValidatorType) Option {
	return Option{Validator: &v}
}

func WithSuggestions(s []string) Option {
	return Option{Suggestions: &s}
}

func WithPlaceholder(p string) Option {
	return Option{Placeholder: &p}
}
