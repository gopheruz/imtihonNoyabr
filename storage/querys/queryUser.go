package querys

var UserCreateQuery string = `
INSERT INTO users(
	first_name,
	last_name,
	email,
	password,
	user_type,
	image_url
)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING id,  created_at
`

var UseGetQuery string = `
         SELECT
			first_name,
			last_name,
            email,
			image_url,
			user_type,
			created_at
		From users WHERE id = $1 and deleted_at is null
`

var UserUpdateQuey = `
	update users set 
			first_name=$1,
			last_name=$2,
			email=$3,
			image_url=$4,
			user_type=$5,
			updated_at=$6
		where id=$7
		returning updated_at
	`

var UserDeleteQuery string = `
	update users set
		deleted_at=$1
		where id=$2
`

var GetAllUserQuery string = `
	 SELECT 
			id,
			first_name,
			last_name,
			email,
			password,
			image_url,
			user_type,
			created_at
		FROM users 
`
