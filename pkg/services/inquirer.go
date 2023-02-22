package service

import (
	dto "calculation-estimate"
	form "calculation-estimate"
	"fmt"
)

type InquirerService struct{}

func NewInquirerService() *InquirerService {
	return &InquirerService{}
}

func (s *InquirerService) GetInquirer() (form.Inquirer, error) {

	question1 := dto.Question{Uid: "311111", Text: "Из чего ваш дом?",
		Options: []string{"кирпич", "панель", "монолитнокаркасный ?"}, Answer: ""}
	question2 := dto.Question{Uid: "3222", Text: "Есть ли проект расстановки освещения и мебели ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question3 := dto.Question{
		Uid: "33333", Text: "Балкон или лоджия.", Options: []string{"да", "нет"}, Answer: "",
		Subquery: dto.Subquery{Uid: "14", Text: "Планирую утеплять и делать теплый пол ?", Options: []string{"да", "нет"}, Answer: "", Conditions: "да"}}
	question4 := dto.Question{Uid: "3444", Text: "Планируете ли сплит систему ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question5 := dto.Question{Uid: "3555", Text: "Планируете установку кондиционера ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question6 := dto.Question{Uid: "3666", Text: "Планируете установку кондиционера ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question7 := dto.Question{Uid: "3777", Text: "Сколько телевизоров ?",
		Options: []string{}, Answer: ""}
	question8 := dto.Question{Uid: "38888", Text: "Есть ли тумба под телевизор в гостиной ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question10 := dto.Question{Uid: "3101010", Text: "Есть детская ?",
		Options: []string{"да", "нет"}, Answer: ""}

	question11 := dto.Question{Uid: "311111111111", Text: "Нужны розетки над рабочим местом или нет ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question12 := dto.Question{Uid: "312121212", Text: "Сколько шагов от щита до кухн###### ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question13 := dto.Question{Uid: "3131313131313", Text: "Духовой шкаф отдельно от варочной ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question14 := dto.Question{Uid: "314141414141414", Text: "Зеркало в санузле с подсветкой ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question15 := dto.Question{Uid: "315151515151515", Text: "Ввод воды и тепла ?",
		Options: []string{"в пластиковых трубах", "метал?"}}
	question16 := dto.Question{Uid: "316161616161616", Text: "Отличное решение со светодиодной подсветкой на кухне, сделаем ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question17 := dto.Question{Uid: "317171717171717", Text: "Кртуо когда над обеденным столом отдельный светильник, считаемм ввод ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question18 := dto.Question{Uid: "31818181818181818", Text: "Сколько рабочих столов будет в квартире, над каждым нужно сделать блок розеток, и под столом также ?",
		Options: []string{}, Answer: ""}
	question19 := dto.Question{Uid: "31919191991919", Text: "Сколько информационных розеток или точек вайфай ?",
		Options: []string{}, Answer: ""}
	question20 := dto.Question{Uid: "320202020202", Text: "Оштукатуренны стены или нет ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question21 := dto.Question{Uid: "21212121212121212", Text: "Планируется водонагреватель ?",
		Options: []string{"проточный", "накопительный", "нет"}}
	question22 := dto.Question{Uid: "32222222222", Text: "Обязательно надо предусмотреть подсветку шкафа сделаем ?",
		Options: []string{"да", "нет"}, Answer: ""}
	question23 := dto.Question{Uid: "3232323232332", Text: "Можно оюъединить несколько линий это сохранит место в щитке и сэкономить средсва ?",
		Options: []string{"да", "нет"}, Answer: ""}

	group1 := dto.Groups{
		Uid:   "21",
		Title: "Группа 1",
		Questions: []dto.Question{
			question1, question2, question3, question4, question5, question6, question7, question8, question10,
		},
	}

	group2 := dto.Groups{
		Uid:   "22",
		Title: "Группа 2",
		Questions: []dto.Question{
			question11, question12, question13, question14, question15, question16, question17, question18, question19,
			question20, question21, question22, question23,
		},
	}

	inquirer := dto.Inquirer{
		Uid: "1",
		Groups: []dto.Groups{
			group1, group2,
		},
	}

	return inquirer, nil
}

func (s *InquirerService) СalculationEstimate(inquirer dto.Inquirer) ([]dto.Estimate, error) {

	fmt.Println("services")

	estimates, err := RunPdfReport(inquirer)

	if err != nil {
		panic(err)
	}

	return estimates, err
}
