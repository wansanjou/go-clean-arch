package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"

	"github.com/bxcodec/go-clean-arch/internal/repository"
	"github.com/bxcodec/go-clean-arch/internal/rest"
	service "github.com/bxcodec/go-clean-arch/pdf"

	"github.com/joho/godotenv"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//prepare database
	// dbHost := os.Getenv("DATABASE_HOST")
	// dbPort := os.Getenv("DATABASE_PORT")
	// dbUser := os.Getenv("DATABASE_USER")
	// dbPass := os.Getenv("DATABASE_PASS")
	// dbName := os.Getenv("DATABASE_NAME")
	// connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// val := url.Values{}
	// val.Add("parseTime", "1")
	// val.Add("loc", "Asia/Jakarta")
	// dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	// dbConn, err := sql.Open(`mysql`, dsn)
	// if err != nil {
	// 	log.Fatal("failed to open connection to database", err)
	// }
	// err = dbConn.Ping()
	// if err != nil {
	// 	log.Fatal("failed to ping database ", err)
	// }

	// defer func() {
	// 	err := dbConn.Close()
	// 	if err != nil {
	// 		log.Fatal("got error when closing the DB connection", err)
	// 	}
	// }()
	// prepare echo

	// e := echo.New()
	e := fiber.New()
	// e.Use(middleware.CORS)
	// timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	// timeout, err := strconv.Atoi(timeoutStr)
	// if err != nil {
	// 	log.Println("failed to parse timeout, using default timeout")
	// 	timeout = defaultTimeout
	// }
	// timeoutContext := time.Duration(timeout) * time.Second
	// e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Prepare Repository
	// authorRepo := mysqlRepo.NewAuthorRepository(dbConn)
	// articleRepo := mysqlRepo.NewArticleRepository(dbConn)
	pdfRepo := repository.NewPDFRepository()

	// Build service Layer
	// svc := article.NewService(articleRepo, authorRepo )
	// rest.NewArticleHandler(e, svc)
	svc := service.NewService(pdfRepo)
	rest.NewPDFHandler(e, svc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Listen(address)) //nolint
}
