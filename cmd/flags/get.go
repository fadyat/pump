package flags

type GetFlags struct {
	OnlyActive   bool
	OnlyInactive bool
}

func NewGetFlags() *GetFlags {
	return &GetFlags{
		OnlyActive:   false,
		OnlyInactive: false,
	}
}
