package main

import (
	"log"
)

func getOne(id int) (a student, err error) {
	a = student{}
	log.Println(db)
	err = db.QueryRow("Select id,user_name,user_sex,user_age from test.student where id=?",
		id).Scan(&a.id, &a.userName, &a.userSex, &a.userAge)
	return
}

func getMany(id int) (students []student, err error) {
	rows, err := db.Query("Select id,user_name,user_sex,user_age from test.student where id > ?", id)
	for rows.Next() {
		a := student{}
		err = rows.Scan(&a.id, &a.userName, &a.userSex, &a.userAge)
		if err != nil {
			log.Fatalln(rows.Err())
		}
		students = append(students, a)
	}
	return
}

func (a *student) update() (err error) {
	_, err = db.Exec("update test.student set user_name=?,user_sex=?,user_age=? where id=?", a.userName, a.userSex, a.userAge, a.id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
