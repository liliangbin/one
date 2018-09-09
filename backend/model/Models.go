package model

/*
CREATE TABLE one
(
id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
text varchar(255) NOT NULL,
img_url varchar(255) NOT NULL,
day varchar(255),
month varchar(255),
time varchar(255),
f_id  int not  null
);
*/

type One struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	ImgUrl string `json:"imgUrl"`
	Day    string `json:"day"`
	Month  string `json:"month"`
	Time   string `json:"time"`
	FullId int    `json:"f_id"`
}

type Ones [] One
