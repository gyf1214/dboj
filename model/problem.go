package model

import "github.com/gyf1214/dboj/util"

// ProblemInfo stores problem information
type ProblemInfo struct {
	ID        int
	Owner     int
	OwnerName string
	Title     string
	Desc      string
}

// DatasetInfo stores dataset information
type DatasetInfo struct {
	ID     int
	Score  int
	Input  string
	Answer string
}

// CreateProblem inserts a problem
func CreateProblem(uid int, title, desc string) (int, error) {
	q := "insert into `problem` (`owner`, `title`, `description`) values (?, ?, ?);"
	res, err := db.Exec(q, uid, title, desc)
	if err != nil {
		return 0, err
	}
	pid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(pid), nil
}

// ListProblem list problems, only returns id & title
func ListProblem(page, uid int) ([]ProblemInfo, error) {
	q := "select `id`, `title` from `problem` where ? = 0 or ? = `owner` order by `id` desc limit ? offset ?;"
	rows, err := db.Query(q, uid, uid, util.PageSize, util.PageSize*page)
	if err != nil {
		return nil, err
	}

	ret := []ProblemInfo{}
	defer rows.Close()
	for rows.Next() {
		prob := ProblemInfo{}
		err := rows.Scan(&prob.ID, &prob.Title)
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
	q := "select `problem`.`id`, `problem`.`title`, `problem`.`description`, `user`.`id`, `user`.`name` from `problem` left join `user` on `problem`.owner = `user`.id where `problem`.`id` = ?;"
	err := db.QueryRow(q, pid).Scan(&ret.ID, &ret.Title, &ret.Desc, &ret.Owner, &ret.OwnerName)
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
