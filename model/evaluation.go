package model

// SubmitInfo stores info of submition
type SubmitInfo struct {
	ID       int
	Problem  int
	User     int
	Code     string
	Language string
}

type EvaluationInfo struct {
	ID      int
	Submit  int
	Dataset int
	Judge   int
	Status  int
	Message string
}

// CreateSubmit creates a submit
func CreateSubmit(info SubmitInfo) (int, error) {
	q := "insert into `submition` (`problem`, `user`, `code`, `language`) values (?, ?, ?, ?);"
	res, err := db.Exec(q, info.Problem, info.User, info.Code, info.Language)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

// ListLanguage returns all the language supported
func ListLanguage() ([]string, error) {
	q := "select distinct `language` from `judge`;"
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for rows.Next() {
		var now string
		err = rows.Scan(&now)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
}

// ListJudge returns judge with language
func ListJudge(language string) ([]int, error) {
	q := "select `id` from `judge` where `language` = ?;"
	rows, err := db.Query(q, language)
	if err != nil {
		return nil, err
	}

	ret := []int{}
	for rows.Next() {
		var now int
		err = rows.Scan(&now)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
}

func ListEvalution(submit int) ([]EvaluationInfo, error) {
	q := "select `id`, `status` from `evaluation` where `submition` = ?;"
	rows, err := db.Query(q, submit)
	if err != nil {
		return nil, err
	}

	ret := []EvaluationInfo{}
	for rows.Next() {
		var now EvaluationInfo
		err = rows.Scan(&now.ID, &now.Status)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
}

func UpdateEvaluation(info EvaluationInfo) error {
	q := "update `evaluation` set `judge` = ?, `status` = ?, `message` = ? where `id` = ?;"
	_, err := db.Exec(q, info.Judge, info.Status, info.Message, info.ID)
	return err
}
