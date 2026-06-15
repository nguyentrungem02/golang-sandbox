package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		errs := make(map[string]string)

		for _, e := range validationError {
			// Convert Field PascalCase to snake_case
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")

			for i, part := range parts {
				if strings.Contains(part, "[") {
					idx := strings.Index(part, "[")
					base := pascalToSnake(part[:idx])
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = pascalToSnake(part)
				}
			}

			fieldPath := strings.Join(parts, ".")

			switch e.Tag() {
			case "gt":
				errs[fieldPath] = fmt.Sprintf("%s phải lớn hơn %s", fieldPath, e.Param())
			case "gte":
				errs[fieldPath] = fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", fieldPath, e.Param())
			case "lt":
				errs[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn %s", fieldPath, e.Param())
			case "lte":
				errs[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", fieldPath, e.Param())
			case "uuid":
				errs[fieldPath] = fmt.Sprintf("%s phải là UUID hợp lệ", fieldPath)
			case "slug":
				errs[fieldPath] = fmt.Sprintf("%s chỉ được chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm", fieldPath)
			case "min":
				errs[fieldPath] = fmt.Sprintf("%s phải nhiều hơn %s ký tự", fieldPath, e.Param())
			case "max":
				errs[fieldPath] = fmt.Sprintf("%s phải ít hơn %s ký tự", fieldPath, e.Param())
			case "min_int":
				errs[fieldPath] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", fieldPath, e.Param())
			case "max_int":
				errs[fieldPath] = fmt.Sprintf("%s phải có giá trị bé hơn %s", fieldPath, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errs[fieldPath] = fmt.Sprintf("%s phải là một trong các giá trị: %s", fieldPath, allowedValues)
			case "required":
				errs[fieldPath] = fmt.Sprintf("%s là bắt buộc", fieldPath)
			case "search":
				errs[fieldPath] = fmt.Sprintf("%s chỉ được chứa chữ thường, in hoa, số và khoảng trắng", fieldPath)
			case "email":
				errs[fieldPath] = fmt.Sprintf("%s phải đúng định dạng emai", fieldPath)
			case "datetime":
				errs[fieldPath] = fmt.Sprintf("%s phải theo đúng định dạng YYYY-MM-DD", fieldPath)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ", ")
				errs[fieldPath] = fmt.Sprintf("%s chỉ cho phép những file có extension: %s", fieldPath, allowedValues)
			}
		}

		return gin.H{"error": errs}
	}

	return gin.H{
		"error": "Yêu cầu không hợp lệ " + err.Error(),
	}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
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
