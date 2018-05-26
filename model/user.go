package model

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/gyf1214/dboj/util"
)

// UserInfo is user information returned
type UserInfo struct {
	ID        int
	Name      string
	Group     int
	GroupName string
}

func random() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", b), nil
}

func hash(a string) string {
	h := md5.New()
	io.WriteString(h, a)
	io.WriteString(h, util.Salt1)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Login process user login
func Login(name, passwd string) (string, error) {
	passwd = hash(passwd)

	var uid int
	q := "select `id` from `user` where `name` = ? and `passwd` = ?;"
	err := db.QueryRow(q, name, passwd).Scan(&uid)
	if err != nil {
		return "", err
	}

	sid, err := random()
	if err != nil {
		return "", err
	}
	q = "update `user` set `session` = ?, `activity` = now() where `id` = ?;"
	_, err = db.Exec(q, sid, uid)
	if err != nil {
		return "", err
	}

	return sid, nil
}

// Authenticate process session authenticate
func Authenticate(sid string) (int, error) {
	var (
		uid int
		act time.Time
	)

	q := "select `id`, `activity` from `user` where `session` = ?;"
	err := db.QueryRow(q, sid).Scan(&uid, &act)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	dur := time.Now().Sub(act)
	if dur > util.SessionExpire {
		return 0, nil
	}

	q = "update `user` set `activity` = now() where `id` = ?;"
	_, err = db.Exec(q, uid)
	if err != nil {
		return 0, err
	}

	return uid, nil
}

// GetUserInfo returns user information
func GetUserInfo(uid int) (UserInfo, error) {
	var ret UserInfo
	q := "select `user`.`id`, `user`.`name`, coalesce(`group`.`id`, 0), coalesce(`group`.`name`, '') from `user` left join `group` on `user`.`group` = `group`.`id` where `user`.`id` = ?;"
	err := db.QueryRow(q, uid).Scan(&ret.ID, &ret.Name, &ret.Group, &ret.GroupName)
	if err != nil {
		return UserInfo{}, err
	}
	return ret, nil
}

// UpdateUserGroup changes the group of user
func UpdateUserGroup(uid, gid int) error {
	q := "update `user` set `group` = ? where `id` = ?;"
	_, err := db.Exec(q, gid, uid)

	return err
}

// Register deals with new user
func Register(name, passwd string) (int, string, error) {
	passwd = hash(passwd)
	sid, err := random()
	if err != nil {
		return 0, "", err
	}

	q := "insert into `user` (`name`, `passwd`, `session`, `activity`) values (?, ?, ?, now());"
	res, err := db.Exec(q, name, passwd, sid)
	if err != nil {
		return 0, "", err
	}

	uid, err := res.LastInsertId()
	if err != nil {
		return 0, "", err
	}

	return int(uid), sid, nil
}

// Logout clears the session
func Logout(sid string) error {
	q := "update `user` set `session` = null where `session` = ?;"
	_, err := db.Exec(q, sid)
	return err
}
