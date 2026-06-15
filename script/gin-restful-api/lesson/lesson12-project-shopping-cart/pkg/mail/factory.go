package mail

import (
	"fmt"

	"trungem.com/shopping-cart/internal/utils"
)

type ProviderType string

const (
	ProviderMailtrap ProviderType = "mailtrap"
)

type ProviderFactory interface {
	CreateProvider(config *EmailConfig) (EmailProviderService, error)
}

type MailtrapProviderFactory struct{}

func (f *MailtrapProviderFactory) CreateProvider(config *EmailConfig) (EmailProviderService, error) {
	return NewMailtrapProvider(config)
}

func NewProviderFactory(providerType ProviderType) (ProviderFactory, error) {
	switch providerType {
	case ProviderMailtrap:
		return &MailtrapProviderFactory{}, nil
	default:
		return nil, utils.NewError(fmt.Sprintf("Unsuppported provider type: %s", utils.ErrorCode(providerType)), utils.ErrCodeInternal)
	}
}
