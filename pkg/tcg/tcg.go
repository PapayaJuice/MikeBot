package tcg

import "fmt"

const (
	tcgRespTemplate = "```\nName: %s\nRarity: %s\nOracle Text: %s\n```\n"
)

// SearchTCG ...
func SearchTCG(text string) (string, string, error) {
	cardID, err := searchCard(text)
	if err != nil {
		return "", "", err
	}
	cardInfo, err := requestCard(cardID)
	if err != nil {
		return "", "", err
	}

	rarity := "Not Found"
	oracle := "Not Found"
	for _, ex := range cardInfo.ExtendedData {
		switch ex.Name {
		case "Rarity":
			rarity = ex.Value
		case "OracleText":
			oracle = ex.Value
		}
	}
	return fmt.Sprintf(tcgRespTemplate, cardInfo.Name, rarity, oracle), cardInfo.ImageURL, nil
}
