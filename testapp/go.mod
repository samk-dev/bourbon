module testapp

go 1.19

replace github.com/samk-dev/bourbon => ../bourbon

require github.com/samk-dev/bourbon v0.0.0-00010101000000-000000000000

require (
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
