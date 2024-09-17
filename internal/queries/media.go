package queries

var GetAdmin = "SELECT username, password FROM admins WHERE username = $1"
