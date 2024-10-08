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
