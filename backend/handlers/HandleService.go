package handlers

import (
	"../model"
	"../dto"
	)

func Max() model.One {

	var one model.One
	db := &dto.Mysql_db{}
	db.Mysql_open()
	one =db.Mysql_select_Max("select * from one where f_id=(select max(f_id) from one)")
	db.Mysql_close()
	return one
}


func FindId(id int ) model.One {
	var one model.One
	db := &dto.Mysql_db{}
	db.Mysql_open()
	one =db.Mysql_select_id(id)
	db.Mysql_close()
	return one
}