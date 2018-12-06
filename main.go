package main

import "github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/infrastructure"

func main() {
	router := infrastructure.NewRouter()
	router.Run()
}
