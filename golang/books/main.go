package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// User 結構體，代表用戶
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Book 結構體，代表書籍
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// 連接 PostgreSQL 數據庫
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 測試數據庫連接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 創建 gin 路由器
	r := gin.Default()

	// 創建 /users 路由，用於管理用戶
	r.GET("/users", func(c *gin.Context) {
		// 查詢所有用戶
		rows, err := db.Query("SELECT id, username, password FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// 遍歷數據庫中的每一行用戶記錄，並添加到 users 切片中
		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username, &user.Password)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		// 將用戶切片作為 JSON 返回
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	// 創建 /users/:id 路由，用於查詢單個用戶
	r.GET("/users/:id", func(c *gin.Context) {
		// 從路由參數中獲取用戶 ID
		id := c.Param("id")

		// 查詢指定 ID 的用戶
		var user User
		row := db.QueryRow("SELECT id, username, password FROM users WHERE id=$1", id)
		err := row.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			log.Fatal(err)
		}

		// 將用戶結構體作為 JSON 返回
		c.JSON(http.StatusOK, user)
	})

	// 創建 /users 路由，用於添加新用戶
	r.POST("/users", func(c *gin.Context) {
		// 從 POST 請求中解析用戶 JSON 數
		var user User
		err := c.BindJSON(&user)
		if err != nil {
			log.Fatal(err)
		}

		// 在數據庫中插入新用戶記錄
		row := db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password)
		err = row.Scan(&user.ID)
		if err != nil {
			log.Fatal(err)
		}

		// 返回新建的用戶記錄
		c.JSON(http.StatusOK, user)
	})

	// 創建 /books 路由，用於管理書籍
	r.GET("/books", func(c *gin.Context) {
		// 查詢所有書籍
		rows, err := db.Query("SELECT id, title, author FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// 遍歷數據庫中的每一行書籍記錄，並添加到 books 切片中
		var books []Book
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.ID, &book.Title, &book.Author)
			if err != nil {
				log.Fatal(err)
			}
			books = append(books, book)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		// 將書籍切片作為 JSON 返回
		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	})

	// 創建 /books/:id 路由，用於查詢單個書籍
	r.GET("/books/:id", func(c *gin.Context) {
		// 從路由參數中獲取書籍 ID
		id := c.Param("id")

		// 查詢指定 ID 的書籍
		var book Book
		row := db.QueryRow("SELECT id, title, author FROM books WHERE id=$1", id)
		err := row.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}

		// 將書籍結構體作為 JSON 返回
		c.JSON(http.StatusOK, book)
	})

	// 創建 /books 路由，用於添加新書籍
	r.POST("/books", func(c *gin.Context) {
		// 從 POST 請求中解析書籍 JSON 數據
		var book Book
		err := c.BindJSON(&book)
		if err != nil {
			log.Fatal(err)
		}

		// 在數據庫中插入新書籍記錄
		row := db.QueryRow("INSERT INTO books (title, author) VALUES ($1, $2) RETURNING id", book.Title, book.Author)
		err = row.Scan(&book.ID)
		if err != nil {
			log.Fatal(err)
		}

		// 返回新建的書籍記錄
		c.JSON(http.StatusOK, book)
	})

	// 啟動 gin 服務
	if err := r.Run(":8080"); err != nil {
		fmt.Println("啟動 gin 服務失敗")
	}
}

/**
在上面的示例中，我們首先使用 `sql.Open` 創建了一個到本機 PostgreSQL 數據庫的連接。
然後，我們定義了 `User` 和 `Book` 兩個結構體，用於表示用戶和書籍的數據。
接著，我們使用 Gin 框架定義了幾個路由，用於管理用戶和書籍的數據。
最後，我們使用 `r.Run` 啟動了 Gin 服務，並指定了服務運行的端口號為 8080。

在使用本示例之前，您需要確保已經安裝了 `github.com/gin-gonic/gin` 和 `github.com/lib/pq` 兩個庫，
並且已經正確配置了本機 PostgreSQL 數據庫。在上面的代碼中，
我們使用 `postgres://postgres:password@localhost:5432/mydb?sslmode=disable`
作為連接字符串，其中 `postgres` 是用戶名，`password` 是密碼，`localhost` 是主機名，`5432` 是端口號，`mydb` 是數據庫名。
如果您的本機 PostgreSQL 數據庫使用不同的配置，請根據實際情況進行修改。
*/
