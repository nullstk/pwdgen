package core

import (
	"fmt"
	"regexp"
	"strings"
)

// Validator checks values against rules and collects errors.
type Validator struct {
	errors []string
}

// NewValidator creates an empty validator.
func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) add(err string) {
	v.errors = append(v.errors, err)
}

// Required ensures the value is not empty.
func (v *Validator) Required(name, value string) *Validator {
	if strings.TrimSpace(value) == "" {
 v.add(fmt.Sprintf("%s is required", name))
	}
	return v
}

// MinLength ensures the value is long enough.
func (v *Validator) MinLength(name, value string, min int) *Validator {
	if len(value) < min {
 v.add(fmt.Sprintf("%s must be at least %d characters", name, min))
	}
	return v
}

// MaxLength ensures the value is short enough.
func (v *Validator) MaxLength(name, value string, max int) *Validator {
	if len(value) > max {
 v.add(fmt.Sprintf("%s must be at most %d characters", name, max))
	}
	return v
}

// Slug ensures the value is a valid slug.
func (v *Validator) Slug(name, value string) *Validator {
	re := regexp.MustCompile(`^[a-z0-9-]+$`)
	if !re.MatchString(value) {
 v.add(fmt.Sprintf("%s must be a lowercase slug", name))
	}
	return v
}

// Valid returns true when no errors were collected.
func (v *Validator) Valid() bool {
	return len(v.errors) == 0
}

// Errors returns the collected errors.
func (v *Validator) Errors() []string {
	return v.errors
}