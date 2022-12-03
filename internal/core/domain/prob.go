package domain

import (
	"errors"
	"math/rand"
	"reflect"
)

type (

	// discrete random variable
	DiscreteRandVar struct {
		Prob   float32
		Choice interface{}
	}

	// probability density function
	ProbDensityFunc struct {
		Vars []DiscreteRandVar
	}
)

func NewPDF(choices interface{}, probs []float32) (*ProbDensityFunc, error) {
	if reflect.TypeOf(choices).Kind() != reflect.Slice {
		return nil, errors.New("invalid choices")
	}
	choicesConv := reflect.ValueOf(choices)
	if choicesConv.Len() == 0 {
		return nil, errors.New("empty choices")
	}
	if choicesConv.Len() != len(probs) {
		return nil, errors.New("unmatched choices and probs")
	}

	pdf := ProbDensityFunc{}
	for i := 0; i < choicesConv.Len(); i++ {
		pdf.Vars = append(pdf.Vars, DiscreteRandVar{
			Prob:   probs[i],
			Choice: choicesConv.Index(i).Interface(),
		})
	}

	if !pdf.IsValid() {
		return nil, errors.New("provided probabilities are invalid")
	}
	return &pdf, nil
}

func NewPDFFromValues(values interface{}) (*ProbDensityFunc, error) {
	if reflect.TypeOf(values).Kind() != reflect.Slice {
		return nil, errors.New("invalid values")
	}
	valuesConv := reflect.ValueOf(values)
	if valuesConv.Len() == 0 {
		return nil, errors.New("empty values")
	}
	freq := make(map[interface{}]int)
	for i := 0; i < valuesConv.Len(); i++ {
		freq[valuesConv.Index(i).Interface()]++
	}

	pdf := ProbDensityFunc{}
	pdfIdx := 0
	for k, v := range freq {
		pdf.Vars = append(pdf.Vars, DiscreteRandVar{
			Prob:   float32(v) / float32(valuesConv.Len()),
			Choice: k,
		})
		pdfIdx++
	}

	if !pdf.IsValid() {
		return nil, errors.New("provided probabilities are invalid")
	}
	return &pdf, nil
}

// TODO: add unit test
// check if sum of provided probabilities is equal to 1
func (p *ProbDensityFunc) IsValid() bool {
	if len(p.Vars) == 0 {
		return false
	}
	var sumProb float32 = 0
	for i := range p.Vars {
		if p.Vars[i].Prob > 1 || p.Vars[i].Prob < 0 {
			return false
		}
		sumProb += p.Vars[i].Prob
	}
	return !(sumProb > 1)
}

// compute cdf & find the correct bucket: https://stackoverflow.com/questions/50507513/golang-choice-number-from-slice-array-with-given-probability
// solution from Patwie
func (p *ProbDensityFunc) Rand() (int, error) {
	if !p.IsValid() {
		return 0, errors.New("provided probabilities are invalid")
	}

	// init cdf
	cdf := make([]float32, len(p.Vars))
	cdf[0] = p.Vars[0].Prob
	for i := 1; i < len(cdf); i++ {
		cdf[i] = cdf[i-1] + p.Vars[i].Prob
	}

	rValue := rand.Float32()
	// find the random index according to the cdf[i]
	randIdx := 0
	for rValue > cdf[randIdx] {
		randIdx++
	}
	return randIdx, nil
}
