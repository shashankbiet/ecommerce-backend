package model

type ProductStatus int64

const (
	ProductStatusDraft    ProductStatus = 0
	ProductStatusActive   ProductStatus = 1
	ProductStatusArchived ProductStatus = 2
)

func (status ProductStatus) IsValid() bool {
	return status == ProductStatusDraft ||
		status == ProductStatusActive ||
		status == ProductStatusArchived
}
