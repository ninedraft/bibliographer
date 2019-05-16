package main

import (
	"strings"
	"testing"
	"time"
)

func TestTextBibliography(test *testing.T) {
	var text = strings.Split(`Загорские штрукли (хорв. Zagorski štrukli) — популярное традиционное хорватское блюдо родом из Хорватского Загорья и Загреба . Популярность штруклей у местного населения переросла пределы Крапинско-Загорской и Вараждинской жупаний, и поэтому их сегодня зачастую готовят с целью популяризации хорватской гастрономии в мире[strukli official]. Штрукли представляют собой мучное блюдо с сыром и яйцами, которое поливается сливочным маслом, сливками или соусом. Это очень нежное блюдо нужно есть вилкой[ТОП смачних страв, які треба спробувати у країнах Європи]. Загорские штрукли родственны традиционным словенским «штруклям» (словен. Štruklji)[Bogataj, Janez, 2007].
	Приготовление
	
	Штрукли можно приготовить двумя способами: путём варки и выпекания[strukli official]. Тесто, наполненное коровьим сыром, используется в обоих методах. Отваренные в подсоленной воде штрукли поливают растопленным жиром и посыпают поджаренными хлебными крошками, а также подают в ароматном супе. Запечённые штрукли перед термической обработкой дополнительно поливают сливками.
	
	Для приготовления штруклей необходимо замесить тесто из муки, соли, одного яйца, тёплой воды, добавить уксуса и масла, дать немного полежать, потом хорошо раскатать его и наполнить смесью коровьего сыра, соли, яиц и сметаны с добавлением сахара для сладких штруклей. Наполненное тесто быстро заворачивают с помощью скатерти, а затем разделяют на подушечки, которые затем готовят выбранным способом.
	Культурное наследие
	
	В 2007 году загорские штрукли были внесены в список нематериального культурного наследия Хорватии, который ведёт Министерство культуры Хорватии[Croatian Ministry of Culture].
	
	В городке Кумровец с 2009 года проводится Штруклиада — специальное мероприятие, организованное с целью продвижения и сохранения традиционного метода приготовления штруклей, которое из года в год привлекает всё больше местных жителей и зарубежных гостей[strukli official].`, "\n")
	var bib = NewBibliography(
		"Источники",
		Source{
			Indent:       "strukli official",
			Title:        "Загорские штрукли",
			Type:         ruSources.Digital,
			URL:          "https://croatia.hr/ru-RU/experiences/gastronomy-and-enology/central-croatia/zagorje-strukli",
			RefferalDate: time.Now(),
			Desc:         "Официальный сайт Хорватского туристического общества",
		},
		Source{
			Indent:       "Croatian Ministry of Culture",
			Title:        "Lista zaštićenih nematerijalnih kulturnih dobara",
			Type:         ruSources.Digital,
			URL:          "http://www.min-kulture.hr/default.aspx?id=3650",
			RefferalDate: time.Now(),
			Desc:         "Официальный сайт Министерства Культуры Хорватии",
		},
		Source{
			Indent:       "ТОП смачних страв, які треба спробувати у країнах Європи",
			Title:        "ТОП смачних страв, які треба спробувати у країнах Європи",
			URL:          "https://maximum.fm/top-smachnih-strav-yaki-treba-sprobuvati-u-krayinah-yevropi_n118657",
			Type:         ruSources.Blog,
			RefferalDate: time.Now(),
			Desc:         "Блог радио МАКСИМУМ",
		},
		Source{
			Indent:  "Bogataj, Janez, 2007",
			Type:    "книга",
			Title:   "Okusiti Slovenijo",
			Authors: []string{"Bogataj, Janez."},
			Desc:    "Ljubljana : Darila Rokus, 2007",
		},
	)

	var textWithBibliography, err = AddBibliographyToText(text, bib, GOST)
	if err != nil {
		test.Fatal(err)
	}
	test.Log(strings.Join(textWithBibliography, "\n"))
}
