package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custon form struct, embeds an object value object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise it return false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initialisez a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

func (f *Form) MinLength(fields string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("this field is not long enough, must be at least %d characters long", length))
		return false
	}
	return true
}
