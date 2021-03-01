package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ActivityLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		write := fmt.Sprintf("%v - %v \t from %v \n", req.Method, req.RequestURI, req.RemoteAddr)

		path := "log/logData.txt"

		// deteksi apakah file sudah ada
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		// buat file baru jika belum ada
		if err != nil || os.IsNotExist(err) {
			var _, err = os.Create(path)
			if err != nil {
				log.Println("Error CREATE", err)
			}
		}
		//tulis logActivity ke file
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		_, err = file.WriteString(write)
		if err != nil {
			log.Println(err)
		}

		defer file.Close()
		//simpan perubahan
		err = file.Sync()
		if err != nil {
			log.Println(err)
		}
		next.ServeHTTP(resp, req)
	})
}
