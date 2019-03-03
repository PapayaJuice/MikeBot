package tcg

func (s *TCGTestSuite) TestSearchCard() {
	r := s.Require()

	cID, err := searchCard("black cat")
	r.NoError(err)
	r.Equal(57806, cID)
}

func (s *TCGTestSuite) TestRequestCard() {
	r := s.Require()

	cID := 57806
	card, err := requestCard(cID)
	r.NoError(err)
	r.NotNil(card)
}
