package search

import (
	"encoding/json"
	"math"
	"sort"
	"strconv"
)

type Document struct {
	id         string
	name       string
	words      []string
	wordsCount map[string]int
	invRate    map[string]float64
	rate       map[string]float64
}

func (d *Document) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name": d.name,
	})
}

type DocumentService struct {
	documents  []Document
	termsCount map[string]int
}

type DocumentArg struct {
	Words []string
	Name  string
}

func BuildDocumentService(documents []DocumentArg) DocumentService {
	termsCount := countDocsWords(documents)
	resDocuments := make([]Document, len(documents))
	for i, doc := range documents {
		invRate := make(map[string]float64, 0)
		wordsCount := make(map[string]int, 0)
		for _, t := range doc.Words {
			x := float64(len(documents)) / float64(termsCount[t])
			invRate[t] = math.Log2(x)
			wordsCount[t]++
		}
		resDocuments[i].id = strconv.Itoa(i)
		resDocuments[i].name = doc.Name
		resDocuments[i].words = doc.Words
		resDocuments[i].invRate = invRate
		resDocuments[i].wordsCount = wordsCount
	}
	for i, doc := range documents {
		rate := make(map[string]float64, 0)
		for _, t := range doc.Words {
			resDoc := resDocuments[i]
			idf := float64(resDoc.wordsCount[t]) / float64(len(resDoc.words))
			rate[t] = resDoc.invRate[t] * idf
		}
		resDocuments[i].rate = rate
	}
	return DocumentService{
		documents:  resDocuments,
		termsCount: termsCount,
	}
}

func countDocsWords(docs []DocumentArg) map[string]int {
	count := make(map[string]int, 0)
	for _, doc := range docs {
		for _, t := range doc.Words {
			count[t]++
		}
	}
	return count
}

func countDocWords(doc []string) map[string]int {
	count := make(map[string]int)
	for _, w := range doc {
		count[w]++
	}
	return count
}

type DocResult struct {
	SimilarityRate float64  `json:"similarityRate"`
	Doc            Document `json:"doc"`
}

type Result struct {
	Docs              []DocResult
	MetricsAggregator *MetricsAggregator `json:"metricsAggregator,omitempty"`
}

func calculateInvRate(docsAmount, termsCount int) float64 {
	x := float64(docsAmount) / float64(termsCount)
	return math.Log2(x)
}

func (s *DocumentService) Search(words []string) Result {
	count := countDocWords(words)
	invRate := make(map[string]float64, 0)
	for _, w := range words {
		invRate[w] = calculateInvRate(len(s.documents), s.termsCount[w])
	}
	rate := make(map[string]float64, 0)
	for _, w := range words {
		rate[w] = invRate[w] * float64(count[w]) / float64(len(words))
	}
	return s.buildSimilarities(words, rate)
}

func (s *DocumentService) buildSimilarities(words []string, rate map[string]float64) Result {
	similarity := make([]DocResult, 0)
	for _, d := range s.documents {
		var numerator float64
		var docLen float64
		for word, docRate := range d.rate {
			docLen += docRate * docRate
			searchRate, ok := rate[word]
			if !ok {
				continue
			}
			numerator += docRate * searchRate
		}
		docLen = math.Sqrt(docLen)

		var searchRate float64
		for _, r := range rate {
			searchRate += r * r
		}
		searchRate = math.Sqrt(searchRate)
		denominator := searchRate * docLen
		similarity = append(similarity, DocResult{
			SimilarityRate: numerator / denominator,
			Doc:            d,
		})
	}
	sort.Slice(similarity, func(i, j int) bool {
		return similarity[i].SimilarityRate > similarity[j].SimilarityRate
	})
	return Result{
		Docs:              similarity,
		MetricsAggregator: addMetricsOnSpecificWords(words),
	}
}

func addMetricsOnSpecificWords(words []string) *MetricsAggregator {
	return &MetricsAggregator{
		DocSetMetrics: BuildDocSetMetrics(1, 0, 2, 1),
		DocSeqMetrics: BuildDocSeqMetrics(4, 1, []float64{0.2, 0.002, 0.0006, 0}),
		AccuracyGraph: BuildAccuracyGraph(),
	}
}
