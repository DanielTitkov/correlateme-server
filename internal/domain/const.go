package domain

const (
	// Correlation

	PearsonCorrelationType  = "pearson"
	SpearmanCorrelationType = "spearman"
	AutoCorrelationType     = "auto"

	// Granularity

	GranularityDay   = "day"
	GranularityWeek  = "week"
	GranularityMonth = "month"

	// Scale types

	ScaleTypeNumeric = "numeric"
	ScaleTypeOrdinal = "ordinal"
	ScaleTypeNominal = "nominal"
	ScaleTypeBinary  = "binary"
)
