package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrdering(test *testing.T) {
	const originalText = `Загорские штрукли (хорв. Zagorski štrukli) — популярное традиционное хорватское блюдо родом из Хорватского Загорья и Загреба . Популярность штруклей у местного населения переросла пределы Крапинско-Загорской и Вараждинской жупаний, и поэтому их сегодня зачастую готовят с целью популяризации хорватской гастрономии в мире[1]. Штрукли представляют собой мучное блюдо с сыром и яйцами, которое поливается сливочным маслом, сливками или соусом. Это очень нежное блюдо нужно есть вилкой[2]. Загорские штрукли родственны традиционным словенским «штруклям» (словен. Štruklji)[3].
	Приготовление
	
	Штрукли можно приготовить двумя способами: путём варки и выпекания[1]. Тесто, наполненное коровьим сыром, используется в обоих методах. Отваренные в подсоленной воде штрукли поливают растопленным жиром и посыпают поджаренными хлебными крошками, а также подают в ароматном супе. Запечённые штрукли перед термической обработкой дополнительно поливают сливками.
	
	Для приготовления штруклей необходимо замесить тесто из муки, соли, одного яйца, тёплой воды, добавить уксуса и масла, дать немного полежать, потом хорошо раскатать его и наполнить смесью коровьего сыра, соли, яиц и сметаны с добавлением сахара для сладких штруклей. Наполненное тесто быстро заворачивают с помощью скатерти, а затем разделяют на подушечки, которые затем готовят выбранным способом.
	Культурное наследие
	
	В 2007 году загорские штрукли были внесены в список нематериального культурного наследия Хорватии, который ведёт Министерство культуры Хорватии[4].
	
	В городке Кумровец с 2009 года проводится Штруклиада — специальное мероприятие, организованное с целью продвижения и сохранения традиционного метода приготовления штруклей, которое из года в год привлекает всё больше местных жителей и зарубежных гостей[1].`

	const text = `Загорские штрукли (хорв. Zagorski štrukli) — популярное традиционное хорватское блюдо родом из Хорватского Загорья и Загреба . Популярность штруклей у местного населения переросла пределы Крапинско-Загорской и Вараждинской жупаний, и поэтому их сегодня зачастую готовят с целью популяризации хорватской гастрономии в мире[strukli official]. Штрукли представляют собой мучное блюдо с сыром и яйцами, которое поливается сливочным маслом, сливками или соусом. Это очень нежное блюдо нужно есть вилкой[ТОП смачних страв, які треба спробувати у країнах Європи]. Загорские штрукли родственны традиционным словенским «штруклям» (словен. Štruklji)[Bogataj, Janez, 2007].
	Приготовление
	
	Штрукли можно приготовить двумя способами: путём варки и выпекания[strukli official]. Тесто, наполненное коровьим сыром, используется в обоих методах. Отваренные в подсоленной воде штрукли поливают растопленным жиром и посыпают поджаренными хлебными крошками, а также подают в ароматном супе. Запечённые штрукли перед термической обработкой дополнительно поливают сливками.
	
	Для приготовления штруклей необходимо замесить тесто из муки, соли, одного яйца, тёплой воды, добавить уксуса и масла, дать немного полежать, потом хорошо раскатать его и наполнить смесью коровьего сыра, соли, яиц и сметаны с добавлением сахара для сладких штруклей. Наполненное тесто быстро заворачивают с помощью скатерти, а затем разделяют на подушечки, которые затем готовят выбранным способом.
	Культурное наследие
	
	В 2007 году загорские штрукли были внесены в список нематериального культурного наследия Хорватии, который ведёт Министерство культуры Хорватии[Croatian Ministry of Culture].
	
	В городке Кумровец с 2009 года проводится Штруклиада — специальное мероприятие, организованное с целью продвижения и сохранения традиционного метода приготовления штруклей, которое из года в год привлекает всё больше местных жителей и зарубежных гостей[strukli official].`

	var sources = []string{
		"strukli official",
		"Croatian Ministry of Culture",
		"ТОП смачних страв, які треба спробувати у країнах Європи",
		"Bogataj, Janez, 2007",
	}
	var ordering = NewSourceOrdering(sources)
	for _, chunk := range strings.Split(text, "\n") {
		ordering.ScanText(chunk)
	}
	var restoredText = text
	for i, source := range ordering.OrderedSources() {
		var textRef = fmt.Sprintf("[%s]", source)
		var numRef = fmt.Sprintf("[%d]", i+1)
		restoredText = strings.ReplaceAll(restoredText, textRef, numRef)
	}
	if restoredText != originalText {
		var diff = cmp.Diff(originalText, restoredText)
		test.Fatal(diff)
	}
}
