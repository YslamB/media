package queries

var GetAdmin = "SELECT username, password FROM admins WHERE username = $1"

var CreateMusic = `
	insert into musics 
		(sub_category_id, language, title, description, path, image_path) 
		values ($1, $2, $3, $4, $5, $6) 
	returning id;
`

var CreateFilm = `
	insert into films 
		(sub_category_id, language, title, description, path, image_path) 
		values ($1, $2, $3, $4, $5, $6) 
	returning id;
`

var CreateBook = `
	insert into books 
		(sub_category_id, language, title, description, path, image_path) 
		values ($1, $2, $3, $4, $5, $6) 
	returning id;
`

var DeleteMusic = `
	delete from musics where id = $1 returning path;
`

var DeleteFilm = `
	delete from films where id = $1 returning path;
`

var DeleteBook = `	
	delete from books where id = $1 returning path;
`

var GetFilms = `
	select id, sub_category_id, language, title, description, path, image_path from films offset $1 limit $2;
`

var GetBooks = `
	select id, sub_category_id, language, title, description, path, image_path from books offset $1 limit $2;
`

var GetMusics = `
	select id, sub_category_id, language, title, description, path, image_path from musics offset $1 limit $2;	
`
