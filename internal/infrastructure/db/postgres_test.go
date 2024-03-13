package db

// TODO: once we have the GitHub-Action PostgreSQL service ready, uncomment the code below.
// Otherwise, it throws the following error: dial tcp [::1]:5432: connect: connection refused

// func TestNewPgConn(t *testing.T) {
// 	// Ensure you have the pgdb container up and running
// 	cfg := config.PostgreSQL{
// 		Host:   "localhost",
// 		Port:   5432,
// 		User:   "camgouser",
// 		Passwd: "camgop4s5W0rD",
// 		DBName: "camgo",
// 	}

// 	// TODO: add more test cases

// 	conn, err := NewDBPgConn(cfg)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, conn)
// }
