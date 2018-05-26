package model

import (
	"database/sql"

	"github.com/gyf1214/dboj/util"
)

// SubmitInfo stores info of submition
type SubmitInfo struct {
	ID       int
	Problem  ProblemInfo
	User     UserInfo
	Code     string
	Language string
	Score    int
}

type JudgeInfo struct {
	ID      int
	Name    string
	Address string
}

type EvaluationInfo struct {
	ID      int
	Submit  int
	Dataset int
	Judge   JudgeInfo
	Status  int
	Message string
}

// CreateSubmit creates a submit
func CreateSubmit(info SubmitInfo) (int, error) {
	q := "insert into `submition` (`problem`, `user`, `code`, `language`) values (?, ?, ?, ?);"
	res, err := db.Exec(q, info.Problem.ID, info.User.ID, info.Code, info.Language)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func GetSubmitInfo(id int) (SubmitInfo, error) {
	ret := SubmitInfo{ID: id}
	q := "select `problem`, `title`, `user`, `name`, `code`, `language`, `score` from `submition_all` where `id` = ?;"
	err := db.QueryRow(q, id).Scan(&ret.Problem.ID, &ret.Problem.Title, &ret.User.ID, &ret.User.Name, &ret.Code, &ret.Language, &ret.Score)
	return ret, err
}

func CountSubmit(uid, pid int) (int, int, error) {
	var ret, ac int
	q := "select count(`id`), coalesce(sum(`accepted`), 0) from `submition_all` where (`user` = ? or 0 = ?) and (`problem` = ? or 0 = ?);"
	err := db.QueryRow(q, uid, uid, pid, pid).Scan(&ret, &ac)
	return ret, ac, err
}

func ListSubmit(uid int, pid int, page int) ([]SubmitInfo, error) {
	q := "select `id`, `problem`, `title`, `score`, `language` from `submition_all` where (`user` = ? or 0 = ?) and (`problem` = ? or 0 = ?) order by `id` desc limit ? offset ?;"
	rows, err := db.Query(q, uid, uid, pid, pid, util.PageSize, util.PageSize*page)
	if err != nil {
		return nil, err
	}

	ret := []SubmitInfo{}
	for rows.Next() {
		var now SubmitInfo
		err := rows.Scan(&now.ID, &now.Problem.ID, &now.Problem.Title, &now.Score, &now.Language)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
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
func ListJudge(language string) ([]JudgeInfo, error) {
	q := "select `id`, `address` from `judge` where `language` = ?;"
	rows, err := db.Query(q, language)
	if err != nil {
		return nil, err
	}

	ret := []JudgeInfo{}
	for rows.Next() {
		var now JudgeInfo
		err = rows.Scan(&now.ID, &now.Address)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
}

func ListEvalution(submit int) ([]EvaluationInfo, error) {
	q := "select `evaluation`.`id`, `evaluation`.`status`, `evaluation`.`dataset`, `judge`.`name` from `evaluation` left join `judge` on `evaluation`.`judge` = `judge`.`id` where `evaluation`.`submition` = ?;"
	rows, err := db.Query(q, submit)
	if err != nil {
		return nil, err
	}

	ret := []EvaluationInfo{}
	for rows.Next() {
		var judge sql.NullString
		var now EvaluationInfo
		err = rows.Scan(&now.ID, &now.Status, &now.Dataset, &judge)
		if err != nil {
			return nil, err
		}
		if judge.Valid {
			now.Judge.Name = judge.String
		}
		ret = append(ret, now)
	}
	return ret, nil
}

func UpdateEvaluation(info EvaluationInfo) error {
	q := "update `evaluation` set `judge` = ?, `status` = ?, `message` = ? where `id` = ?;"
	_, err := db.Exec(q, info.Judge.ID, info.Status, info.Message, info.ID)
	return err
}
