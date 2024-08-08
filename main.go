package main

import (
    "database/sql"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "unicode"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", "user=postgres password=Dragonquest9 dbname=agnos_be sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    router := gin.Default()

    router.POST("/api/strong_password_steps", func(c *gin.Context) {
        var request struct {
            InitPassword string `json:"init_password"`
        }
        if err := c.BindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        numSteps := calculateSteps(request.InitPassword)
        response := gin.H{"num_of_steps": numSteps}

        // Log
        logRequestResponse(request.InitPassword, response)

        c.JSON(http.StatusOK, response)
    })

    router.Run(":8080")
}

func calculateSteps(password string) int {
    length := len(password)
    hasLower := false
    hasUpper := false
    hasDigit := false
    stepsToAdd := 0
    stepsToRemove := 0
    replacementSteps := 0


    repeats := make(map[int]int) 

    for i, ch := range password {
        if unicode.IsLower(ch) {
            hasLower = true
        } else if unicode.IsUpper(ch) {
            hasUpper = true
        } else if unicode.IsDigit(ch) {
            hasDigit = true
        }

        if i > 1 && password[i] == password[i-1] && password[i] == password[i-2] {
            lengthOfRepeat := 2
            for j := i; j < len(password) && password[j] == password[i]; j++ {
                lengthOfRepeat++
            }
            repeats[lengthOfRepeat]++
            i += lengthOfRepeat - 1
        }
    }

    missingTypes := 3 - (boolToInt(hasLower) + boolToInt(hasUpper) + boolToInt(hasDigit))

    // length case
    if length < 6 {
        stepsToAdd = 6 - length
    } else if length > 20 {
        stepsToRemove = length - 20

        for lengthOfRepeat, count := range repeats {
            if lengthOfRepeat > 2 {
                replacements := (lengthOfRepeat - 2) / 3
                replacementSteps += replacements * count
            }
        }

        replacementSteps = max(replacementSteps, missingTypes)

        return stepsToRemove + replacementSteps
    }

	
    for lengthOfRepeat, count := range repeats {
        if lengthOfRepeat > 2 {
            replacements := (lengthOfRepeat - 2) / 3
            replacementSteps += replacements * count
        }
    }

    return max(stepsToAdd, missingTypes) + replacementSteps
}

func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func logRequestResponse(request, response string) {
    _, err := db.Exec("INSERT INTO logs (request, response) VALUES ($1, $2)", request, response)
    if err != nil {
        log.Println("Error logging request/response:", err)
    }
}
