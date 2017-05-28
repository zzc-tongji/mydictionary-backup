package mydictionary

import (
	"strings"

	"github.com/zzc-tongji/bingdictionary4mydictionary"
	"github.com/zzc-tongji/icibacollins4mydictionary"
	"github.com/zzc-tongji/merriamwebster4mydictionary"
	"github.com/zzc-tongji/vocabulary4mydictionary"
)

// quary vocabulary online
func requestOnline(vocabularyAsk vocabulary4mydictionary.VocabularyAskStruct) (vocabularyAnswerList []vocabulary4mydictionary.VocabularyAnswerStruct) {
	var (
		vocabularyAnswerChannel chan vocabulary4mydictionary.VocabularyAnswerStruct
		vocabularyAnswer        vocabulary4mydictionary.VocabularyAnswerStruct
	)
	// prepare
	vocabularyAnswerChannel = make(chan vocabulary4mydictionary.VocabularyAnswerStruct, setting.Online.length)
	// query
	if setting.Online.Service.BingDictionary {
		go func() {
			vocabularyAnswerChannel <- bingdictionary4mydictionary.Request(vocabularyAsk)
		}()
	}
	if setting.Online.Service.IcibaCollins {
		go func() {
			vocabularyAnswerChannel <- icibacollins4mydictionary.Request(vocabularyAsk)
		}()
	}
	if setting.Online.Service.MerriamWebster {
		go func() {
			vocabularyAnswerChannel <- merriamwebster4mydictionary.Request(vocabularyAsk)
		}()
	}
	// add to answer list
	for i := 0; i < setting.Online.length; i++ {
		vocabularyAnswer = <-vocabularyAnswerChannel
		if setting.Online.Debug {
			vocabularyAnswerList = append(vocabularyAnswerList, vocabularyAnswer)
		} else if strings.Compare(vocabularyAnswer.Status, vocabulary4mydictionary.Basic) == 0 {
			vocabularyAnswerList = append(vocabularyAnswerList, vocabularyAnswer)
		}
	}
	return
}