package search

type MetricsAggregator struct {
	DocSetMetrics *DocSetMetrics `json:"docSetMetrics"`
	DocSeqMetrics *DocSeqMetrics `json:"docSeqMetrics"`
	AccuracyGraph *AccuracyGraph `json:"accuracyGraph"`
}

type DocSetMetrics struct {
	Recall    float64 `json:"recall"`
	Precision float64 `json:"precision"`
	Accuracy  float64 `json:"accuracy"`
	Error     float64 `json:"error"`
	FMeasure  float64 `json:"fMeasure"`
}

func BuildDocSetMetrics(foundRelevant, notFoundRelevant, foundNotRelevant, notFoundNotRelevant int) *DocSetMetrics {
	m := DocSetMetrics{
		Recall:    float64(foundNotRelevant) / float64(foundRelevant+notFoundRelevant),
		Precision: float64(foundNotRelevant) / float64(foundRelevant+foundNotRelevant),
		Accuracy:  float64(foundNotRelevant+notFoundNotRelevant) / float64(foundRelevant+foundNotRelevant+notFoundRelevant+notFoundNotRelevant),
		Error:     float64(foundNotRelevant+notFoundRelevant) / float64(foundRelevant+foundNotRelevant+notFoundRelevant+notFoundNotRelevant),
	}
	m.FMeasure = float64(2) / (1/m.Precision + 1/m.Recall)
	return &m
}

type DocSeqMetrics struct {
	Precision        float64 `json:"precision"`
	RPrecision       float64 `json:"rPrecision"`
	AveragePrecision float64 `json:"averagePrecision"`
}

func BuildDocSeqMetrics(totalDocuments, relevantDocuments int, precisions []float64) *DocSeqMetrics {
	return &DocSeqMetrics{
		Precision:        float64(relevantDocuments) / float64(totalDocuments),
		RPrecision:       average(precisions),
		AveragePrecision: average(precisions),
	}
}

type AccuracyGraph struct {
	Points []Point `json:"points"`
}

type Point struct {
	Recall    float64 `json:"recall"`
	Precision float64 `json:"precision"`
}

func BuildAccuracyGraph() *AccuracyGraph {
	points := []Point{
		{
			Precision: 1,
			Recall:    0,
		},
		{
			Precision: 0.8,
			Recall:    0.1,
		},
		{
			Precision: 0.7,
			Recall:    0.2,
		},
		{
			Precision: 0.6,
			Recall:    0.3,
		},
		{
			Precision: 0.55,
			Recall:    0.4,
		},
		{
			Precision: 0.51,
			Recall:    0.5,
		},
		{
			Precision: 0.5,
			Recall:    0.6,
		},
		{
			Precision: 0.45,
			Recall:    0.7,
		},
		{
			Precision: 0.4,
			Recall:    0.8,
		},
		{
			Precision: 0.2,
			Recall:    0.9,
		},
		{
			Precision: 0.3,
			Recall:    1,
		},
	}
	return &AccuracyGraph{
		Points: points,
	}
}

func average(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	return sum / float64(len(s))
}
