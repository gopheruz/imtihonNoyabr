package querys

var HotelCreateQuey string = `
	INSERT INTO hotels(
		hotel_,name,
		adress,
		descriortion
	)
	VALUES($1, $2, $3)
	RETURNING id
`