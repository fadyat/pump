package flags

type GetFlags struct {
	OnlyActive bool
}

func NewGetFlags() *GetFlags {
	return &GetFlags{
		OnlyActive: false,
	}
}
