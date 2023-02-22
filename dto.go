package calculation_estimate

type Inquirer struct {
	Uid    string   `json:"uid"`
	Groups []Groups `json:"groups"`
}

type Groups struct {
	Uid       string     `json:"uid"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Uid      string   `json:"uid"`
	Text     string   `json:"text"`
	Options  []string `json:"options"`
	Answer   string   `json:"answer"`
	Subquery Subquery `json:"subquery"`
}

type Subquery struct {
	Uid        string    `json:"uid"`
	Text       string    `json:"text"`
	Options    []string  `json:"options"`
	Answer     string    `json:"answer"`
	Conditions string    `json:"conditions"`
	Subquery   *Subquery `json:"subquery"`
}

type Estimate struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

//1 добавить подаопрос в модель
//2 исправить фронт
//3 сгенерить пдф
//4 вернуть пдф
