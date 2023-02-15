package enviroment

type Env struct {
	//DocumentDBUri i a monogodb connection string for Document database
	DocumentDBUri string `env:"DOCUMENT_DB_URI"`
	//DocumentDB is the databse for wordDocumennt
	DocumentDB string `env:"DOCUMENT_DB"`
}
