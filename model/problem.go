package model

import "github.com/zc-staff/dboj/util"

// ProblemInfo stores problem information
type ProblemInfo struct {
	ID      int
	Owner   UserInfo
	Title   string
	Desc    string
	Submits int
	Accepts int
}

// DatasetInfo stores dataset information
type DatasetInfo struct {
	ID     int
	Score  int
	Input  string
	Answer string
}

// CreateProblem inserts a problem
func CreateProblem(info ProblemInfo) (int, error) {
	q := "insert into `problem` (`owner`, `title`, `description`) values (?, ?, ?);"
	res, err := db.Exec(q, info.Owner.ID, info.Title, info.Desc)
	if err != nil {
		return 0, err
	}
	pid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(pid), nil
}

// CountProblem count problems
func CountProblem(uid int) (int, error) {
	var ret int
	q := "select count(`problem`.`id`) from `problem` where ? = 0 or ? = `problem`.`owner`;"

	err := db.QueryRow(q, uid, uid).Scan(&ret)
	return ret, err
}

// ListProblem list problems, only returns id & title
func ListProblem(page, uid int) ([]ProblemInfo, error) {
	q := "select `id`, `title`, `owner`, `name`, `submits`, `accepts` from `problem_all` where 0 = ? or `owner` = ? order by `id` asc limit ? offset ?;"
	rows, err := db.Query(q, uid, uid, util.PageSize, util.PageSize*page)
	if err != nil {
		return nil, err
	}

	ret := []ProblemInfo{}
	defer rows.Close()
	for rows.Next() {
		prob := ProblemInfo{}
		err := rows.Scan(&prob.ID, &prob.Title, &prob.Owner.ID, &prob.Owner.Name, &prob.Submits, &prob.Accepts)
		if err != nil {
			return nil, err
		}
		ret = append(ret, prob)
	}
	return ret, nil
}

// GetProblemInfo returns problem information
func GetProblemInfo(pid int) (ProblemInfo, error) {
	var ret ProblemInfo
	q := "select `id`, `title`, `description`, `owner`, `name`, `submits`, `accepts` from `problem_all` where `id` = ?;"
	err := db.QueryRow(q, pid).Scan(&ret.ID, &ret.Title, &ret.Desc, &ret.Owner.ID, &ret.Owner.Name, &ret.Submits, &ret.Accepts)
	return ret, err
}

// GetProblemOwner returns the problem owner
func GetProblemOwner(pid int) (int, error) {
	var ret int
	q := "select `owner` from `problem` where `id` = ?;"
	err := db.QueryRow(q, pid).Scan(&ret)
	return ret, err
}

// UpdateProblem updates a problem
func UpdateProblem(info ProblemInfo) error {
	q := "update `problem` set `title` = ?, `description` = ? where `id` = ?;"
	_, err := db.Exec(q, info.Title, info.Desc, info.ID)
	return err
}

// DeleteProblem deletes a problem
func DeleteProblem(pid int) error {
	q := "delete from `problem` where `id` = ?;"
	_, err := db.Exec(q, pid)
	return err
}

// AddDataset adds a dataset to a problem
func AddDataset(pid int, data DatasetInfo) error {
	q := "insert into `dataset` (`problem`, `score`, `input`, `answer`) values (?, ?, ?, ?);"
	_, err := db.Exec(q, pid, data.Score, data.Input, data.Answer)
	return err
}

func GetDatasetInfo(id int) (DatasetInfo, error) {
	var ret DatasetInfo
	q := "select `input`, `answer` from `dataset` where `id` = ?;"
	err := db.QueryRow(q, id).Scan(&ret.Input, &ret.Answer)
	return ret, err
}
