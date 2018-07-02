package model

import "github.com/zc-staff/dboj/util"

type DiscussionInfo struct {
	ID      int
	User    UserInfo
	Problem ProblemInfo
	Title   string
	Content string
}

type PostInfo struct {
	User       UserInfo
	Discussion DiscussionInfo
	Content    string
}

func CreateDiscussion(info DiscussionInfo) (int, error) {
	q := "insert into `discussion` (`problem`, `user`, `title`, `content`) values (?, ?, ?, ?);"
	res, err := db.Exec(q, info.Problem.ID, info.User.ID, info.Title, info.Content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func CountDiscussion(pid int) (int, error) {
	var ret int
	q := "select count(`id`) from `discussion_all` where ? = 0 or ? = `problem`;"

	err := db.QueryRow(q, pid, pid).Scan(&ret)
	return ret, err
}

func ListDiscussion(pid, page int) ([]DiscussionInfo, error) {
	q := "select `id`, `title`, `content`, `user`, `name`, `problem`, `problem_title` from  `discussion_all` where `problem` = ? or ? = 0 order by `id` desc limit ? offset ?;"
	rows, err := db.Query(q, pid, pid, util.PageSize, page*util.PageSize)
	if err != nil {
		return nil, err
	}

	ret := []DiscussionInfo{}
	defer rows.Close()
	for rows.Next() {
		var now DiscussionInfo
		err := rows.Scan(&now.ID, &now.Title, &now.Content, &now.User.ID, &now.User.Name, &now.Problem.ID, &now.Problem.Title)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
}

func GetDiscussionInfo(id int) (DiscussionInfo, error) {
	q := "select `id`, `title`, `content`, `user`, `name`, `problem`, `problem_title` from  `discussion_all` where `id` = ?;"
	var ret DiscussionInfo
	err := db.QueryRow(q, id).Scan(&ret.ID, &ret.Title, &ret.Content, &ret.User.ID, &ret.User.Name, &ret.Problem.ID, &ret.Problem.Title)
	return ret, err
}

func CreatePost(info PostInfo) (int, error) {
	q := "insert into `post` (`discussion`, `user`, `content`) values (?, ?, ?);"
	res, err := db.Exec(q, info.Discussion.ID, info.User.ID, info.Content)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func ListPost(id int) ([]PostInfo, error) {
	q := "select `post`.`content`, `post`.`user`, `user`.`name` from `post` left join `user` on `post`.`user` = `user`.`id` where `post`.`discussion` = ?;"
	rows, err := db.Query(q, id)
	if err != nil {
		return nil, err
	}

	ret := []PostInfo{}
	for rows.Next() {
		var now PostInfo
		err := rows.Scan(&now.Content, &now.User.ID, &now.User.Name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, now)
	}
	return ret, nil
}
