package tcg

func (s *TCGTestSuite) TestGetBearer() {
	r := s.Require()
	bt, err := requestToken()
	r.NoError(err)
	r.NotNil(bt)
}
