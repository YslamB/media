package queries

var GetAdmin = "SELECT username, password FROM admins WHERE username = $1"

var CreateMusic = `
	insert into musics 
		(sub_category_id, language, title, description) 
		values ($1, $2, $3, $4) 
	returning id;
`

var CreateFilm = `
	insert into films 
		(sub_category_id, language, title, description) 
		values ($1, $2, $3, $4) 
	returning id;
`

var CreateBook = `
	insert into books 
		(sub_category_id, language, title, description) 
		values ($1, $2, $3, $4) 
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
	select id, sub_category_id, language, title, description, path, image_path from films where status=true offset $1 limit $2;
`

var GetBooks = `
	select id, sub_category_id, language, title, description, path, image_path from books offset $1 limit $2;
`

var GetMusics = `
	select id, sub_category_id, language, title, description, path, image_path from musics where status=true offset $1 limit $2;	
`

var CreateCategory = `
	insert into categories (name) 
    	values ($1) returning id;
    
`

var CreateSubCategory = `
	insert into sub_categories (category_id, name) 
    	values ($1, $2) returning id;
`

var GetCategories = `
	select 
    c.id, c.name, sc.sub_categories
	from categories c
	left join (
		select 
			sc.category_id, 
			json_agg(
				json_build_object(
					'id', sc.id,
					'name', sc.name
				) 
			) as sub_categories
		from sub_categories sc
		group by category_id
	) sc on sc.category_id = c.id
`

var UpdateBook = `
	update books set sub_category_id=$1, language=$2, title=$3, description=$4 where id=$5 returning path, image_path;
`

var UpdateFilm = `
	update films set sub_category_id=$1, language=$2, title=$3, description=$4 where id=$5 returning path, image_path;
`

var UpdateMusic = `
	update musics set sub_category_id=$1, language=$2, title=$3, description=$4 where id=$5 returning path, image_path;
`

var GetImageFilePathFilm = `
	select path, image_path, id from films where id=$1;
`

var GetImageFilePathBook = `
	select path, image_path, id from books where id=$1;
`

var GetImageFilePathMusic = `
	select path, image_path, id from musics where id=$1;
`

var UpdateFilmImage = `
	update films set image_path=$1 where id=$2
`

var UpdateFilmPath = `
	update films set path=$1 where id=$2
`

var UpdateMusicImage = `
	update musics set image_path=$1 where id=$2
`

var UpdateMusicPath = `
	update musics set path=$1 where id=$2
`

var UpdateBookImage = `
	update books set image_path=$1 where id=$2
`

var UpdateBookPath = `
	update books set path=$1 where id=$2
`
