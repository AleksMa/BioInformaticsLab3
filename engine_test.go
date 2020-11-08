package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EngineTestSuite struct {
	suite.Suite
	nw *NeedlemanWunsch
}

func (s *EngineTestSuite) TestDNA() {
	for _, test := range []struct {
		seq1    string
		seq2    string
		result1 string
		result2 string
		score   int
	}{
		{
			"AT",
			"G",
			"AT",
			"G-",
			-14,
		},
		{
			"AATCG",
			"AACG",
			"AATCG",
			"AA-CG",
			10,
		},
		{
			"ACGGCTT",
			"ACGT",
			"ACGGCTT",
			"ACG---T",
			8,
		},
		{
			"GTACAACGTTA",
			"AATCGTAGCGA",
			"GTACAA-CGTT---A",
			"----AATCGTAGCGA",
			-9,
		},
	} {
		s.nw = NewNeedlemanWunsch(&Sequence{Value: test.seq1}, &Sequence{Value: test.seq2}, DNAFull, -10)
		r1, r2, sc := s.nw.Solve()
		s.Equal(r1, test.result1)
		s.Equal(r2, test.result2)
		s.Equal(sc, test.score)
	}
}

func TestVkAudienceCountersSuite(t *testing.T) {
	suite.Run(t, new(EngineTestSuite))
}
