package tcg

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type TCGTestSuite struct {
	suite.Suite
}

func TestTCGTestSuite(t *testing.T) {
	httpmock.ActivateNonDefault(&c)
	defer httpmock.DeactivateAndReset()

	mockBearerResp := `{"access_token":"iOq8lffH-EpAumgZg68o3tP_sgtJAV-WJzPyE-S6k4Tr5Su55FTdBqNdEwKVGVfv6EWmzXc2W1Igl_kz3FR7dSw2K61MwtLqYC4bh1WsMFQ0XHzdDKJhyOuvMctYZ4jszS7vbzJz3c_mGuCE2plPGhoB7ulaLKmFBAhDPIQqX1sS4IEHiO1N9pz7-aKVtUJw8oR3QdzmATilMI-k-sNEG3DzpUfLF2bxIGmEVmcG7BJ6zzQg8XHLQhiDjAQrbLA_bg0fsPN3TRLI4h4LaixJLRihBnqlYfHRQZZtDc4Yxrc2t7VY3fVk5HZPFqe8nm0DB4qo8Q","token_type":"bearer","expires_in":1209599,"userName":"96d18aa7-6404-302a-91f8-95c5df43b146",".issued":"Sat, 02 Mar 2019 23:03:43 GMT",".expires":"Sat, 16 Mar 2099 23:03:43 GMT"}`
	httpmock.RegisterResponder(http.MethodPost, tcgpURI, httpmock.NewStringResponder(200, mockBearerResp))

	mockSearchResp := `{"totalItems":2,"success":true,"errors":[],"results":[{"productId":57806,"name":"Black Cat","cleanName":"Black Cat","imageUrl":"https://6d4be195623157e28848-7697ece4918e0a73861de0eb37d08968.ssl.cf1.rackcdn.com/57806_200w.jpg","categoryId":1,"groupId":125,"url":"https://store.tcgplayer.com/magic/dark-ascension/black-cat","modifiedOn":"2017-09-27T17:23:36.867"},{"productId":91086,"name":"Black Cat","cleanName":"Black Cat","imageUrl":"https://6d4be195623157e28848-7697ece4918e0a73861de0eb37d08968.ssl.cf1.rackcdn.com/91086_200w.jpg","categoryId":1,"groupId":1293,"url":"https://store.tcgplayer.com/magic/magic-2015-m15/black-cat","modifiedOn":"2018-02-03T15:03:04.767"}]}`
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(mtgSearchURI, "black+cat"), httpmock.NewStringResponder(200, mockSearchResp))

	mockRequestResp := `{"success":true,"errors":[],"results":[{"productId":57806,"name":"Black Cat","cleanName":"Black Cat","imageUrl":"https://6d4be195623157e28848-7697ece4918e0a73861de0eb37d08968.ssl.cf1.rackcdn.com/57806_200w.jpg","categoryId":1,"groupId":125,"url":"https://store.tcgplayer.com/magic/dark-ascension/black-cat","modifiedOn":"2017-09-27T17:23:36.867","imageCount":1,"presaleInfo":{"isPresale":false,"releasedOn":null,"note":null},"extendedData":[{"name":"Rarity","displayName":"Rarity","value":"C"},{"name":"Number","displayName":"#","value":"54"},{"name":"SubType","displayName":"Creature Type or Sub Type","value":"Creature Zombie Cat"},{"name":"P","displayName":"P","value":"1"},{"name":"T","displayName":"T","value":"1"},{"name":"OracleText","displayName":"Rules Text Contains","value":"When Black Cat dies, target opponent discards a card at random."},{"name":"FlavorText","displayName":"Flavor","value":"Its last life is spent tormenting your dreams."}]}]}`
	httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf(mtgReqURI, 57806), httpmock.NewStringResponder(200, mockRequestResp))

	suite.Run(t, new(TCGTestSuite))
}

func (s *TCGTestSuite) TestSearchTCG() {
	r := s.Require()

	_, _, err := SearchTCG("black cat")
	r.NoError(err)
}
