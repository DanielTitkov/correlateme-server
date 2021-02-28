package domain

type (
	CreateIndicatorArgs struct {
		Username    string
		Title       string
		Description string
		ScaleType   string
	}
	FilterIndicatorsArgs struct {
		ID             []int
		Code           []string
		Title          []string
		Active         *bool
		BuiltIn        *bool
		AuthorUsername *string
		ScaleType      *string
		External       *bool // not accesible via API, only for intenal use
	}
)
