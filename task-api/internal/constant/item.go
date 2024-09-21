package constant

type ItemStatus string

const (
	ItemPendingStatus  ItemStatus = "PENDING"
	ItemApprovedStatus ItemStatus = "APPROVED"
	ItemRejectedStatus ItemStatus = "REJECTED"
)
