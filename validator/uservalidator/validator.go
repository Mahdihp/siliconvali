package uservalidator

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	"siliconvali/dto"
	richerror "siliconvali/pkg"
	"siliconvali/pkg/errmsg"
)

const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface {
	IsMobileUnique(ctx context.Context, mobile string) (bool, error)
	GetByMobile(ctx context.Context, mobile string) (dto.UserInfo, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{repo: repo}
}

func (v Validator) ValidateRegisterRequest(req dto.UserInsertRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateRegisterRequest"

	if err := validation.ValidateStruct(&req,
		// TODO - add 3 to config
		validation.Field(&req.FirstName,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.LastName,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.NationalCode,
			validation.Required,
			validation.Length(1, 11)),

		validation.Field(&req.Address,
			validation.Required,
			validation.Length(1, 500)),

		validation.Field(&req.RoleId,
			validation.Required),

		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`))),

		validation.Field(&req.Mobile,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errmsg.ErrorMsgPhoneNumberIsNotValid),
			validation.By(v.checkPhoneNumberUniqueness)),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)

	if isUnique, err := v.repo.IsMobileUnique(context.Background(), phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(errmsg.ErrorMsgPhoneNumberIsNotUnique)
		}
	}

	return nil
}
