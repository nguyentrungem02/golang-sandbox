package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"trungem.com/shopping-cart/internal/utils"
)

func InitValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	_ = RegisterCustomValidation(v)

	return nil
}

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
					base := utils.PascalToSnake(part[:idx])
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = utils.PascalToSnake(part)
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
			case "email_advanced":
				errs[fieldPath] = fmt.Sprintf("%s %s này nằm trong danh sách đen", fieldPath, e.Value())
			case "password_strong":
				errs[fieldPath] = fmt.Sprintf("%s phải có ít nhất 8 ký tự bao gồm (chữ thường, chữ in hoa, số và ký tự đặc biệt)", fieldPath)
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
		"error":  "Yêu cầu không hợp lệ",
		"detail": err.Error(),
	}
}
