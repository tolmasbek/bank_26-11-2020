package db

const CreateUsersAccount = `Create table if not exists users (
	id 			integer 	primary key	autoincrement,
	name 		text 		not null,
	surname 	text 		not null,
	age 		integer 	not null,
	gender 		text 		not null,
	login 		text 		not null 	unique,
	password 	text 		not null,
	remove 		boolean 	not null 	default false
);`