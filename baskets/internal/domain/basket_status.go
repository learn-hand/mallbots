package domain

import "errors"

type BasketStatus struct {
	slug string
}

func (b BasketStatus) String() string {
	return b.slug
}

var (
	BasketUnknown      = BasketStatus{""}
	BasketIsOpen       = BasketStatus{"open"}
	BasketIsCanceled   = BasketStatus{"canceled"}
	BasketIsCheckedOut = BasketStatus{"checked_out"}
)

func FromString(s string) (BasketStatus, error) {
	switch s {
	case BasketIsOpen.slug:
		return BasketIsOpen, nil
	case BasketIsCanceled.slug:
		return BasketIsCanceled, nil
	case BasketIsCheckedOut.slug:
		return BasketIsCheckedOut, nil
	}
	return BasketUnknown, errors.New("unknown Basket status: " + s)
}
