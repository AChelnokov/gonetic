package gonetic

import (
	codes "github.com/AChelnokov/gonetic/codes"
	"strings"
)

type PhoneticText struct {
	Source       string
	SourceLength int
	IsCyrillic   bool
}

func (p *PhoneticText) DMCodes() map[string]string {
	refString := p.Source

	output := map[string]string{}
	dm := codes.GetDaitchMokotoffMap()

	splitSource := strings.Fields(refString)
	for _, word := range splitSource {
		_, ok := output[word]
		if ok {
			continue
		}

		if codes.NeedTranslit(word) {
			output[word] = dm.Run(codes.Translit(word), true)
			continue
		}
		output[word] = dm.Run(word, false)
	}

	return output
}

func SetText(text string) *PhoneticText {
	var p = PhoneticText{}
	p.Source = text
	p.SourceLength = len([]rune(p.Source))

	return &p
}
