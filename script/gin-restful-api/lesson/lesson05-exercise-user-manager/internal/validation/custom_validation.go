package validation

import (
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"trungem.com/user-manager/internal/utils"
)

func RegisterCustomValidation(v *validator.Validate) error {
	var blockedDomains = map[string]bool{
		"blacklist.com": true,
		"abc.com":       true,
	}
	if err := v.RegisterValidation("email_advanced", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()

		parts := strings.Split(email, "@")
		if len(parts) != 2 {
			return false
		}

		domain := utils.NormalizeString(parts[1])

		return !blockedDomains[domain]
	}); err != nil {
		return err
	}

	if err := v.RegisterValidation("password_strong", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 {
			return false
		}

		hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};:'",.<>?/\\|]`).MatchString(password)

		return hasLower && hasUpper && hasDigit && hasSpecial

	}); err != nil {
		return err
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	if err := v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	}); err != nil {
		return err
	}

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	if err := v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	}); err != nil {
		return err
	}

	if err := v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minInt, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minInt
	}); err != nil {
		return err
	}

	if err := v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxInt, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() <= maxInt
	}); err != nil {
		return err
	}

	if err := v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()
		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if strings.ToLower(allowed) == ext {
				return true
			}
		}

		return false
	}); err != nil {
		return err
	}

	return nil
}
