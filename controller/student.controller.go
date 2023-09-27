package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/johnldev/integrador-mvc/model"
	"gorm.io/gorm"
)

func StartHttpServer(db *gorm.DB) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fillUserList := func(w http.ResponseWriter, r *http.Request) {
		var students []model.Student
		db.Find(&students)

		err := template.Must(template.ParseFiles("view/studentList.html")).Execute(w, students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	r.Get("/", fillUserList)

	r.Get("/student-register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/studentRegister.html")
	})

	r.Post("/student", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		birthDate, _ := time.Parse("2006-01-02", r.FormValue("birthDate"))
		student := &model.Student{
			Name:      r.FormValue("name"),
			Gender:    r.FormValue("gender"),
			BirthDate: birthDate,
			Phone:     r.FormValue("phone"),
			ClassYear: r.FormValue("classYear"),
			TimeImdz:  r.FormValue("timeImdz"),
		}

		db.Create(student)

		fillUserList(w, r)
	})

	r.Get("/student/{id}/mother-info", func(w http.ResponseWriter, r *http.Request) {
		var motherInfo model.MotherInfo
		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

		db.Where(model.MotherInfo{StudentId: uint(id)}).First(&motherInfo)
		if motherInfo.StudentId == 0 {
			motherInfo.StudentId = uint(id)
		}

		fmt.Println(motherInfo)
		err := template.Must(template.ParseFiles("view/motherInfo.html")).Execute(w, motherInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Post("/student/{id}/mother-info", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		parseBoll := func(value string) bool {
			return value == "on"
		}

		var oldMotherInfo model.MotherInfo
		db.Where(model.MotherInfo{StudentId: uint(id)}).First(&oldMotherInfo)
		if oldMotherInfo.StudentId == 0 {
			oldMotherInfo.StudentId = uint(id)
		}

		if oldMotherInfo.CreatedAt.IsZero() {
			oldMotherInfo.CreatedAt = time.Now()
		}

		if oldMotherInfo.UpdatedAt.IsZero() {
			oldMotherInfo.UpdatedAt = time.Now()
		}
		helperName := r.FormValue("helperName")
		motherInfo := &model.MotherInfo{
			AuthorizedPeople: r.FormValue("authorizedPeople"),
			HelperName:       helperName,
			HasHelper:        helperName != "",
			ProjectName:      r.FormValue("projectName"),
			WorkOutside:      parseBoll(r.FormValue("workOutside")),
			ReceiveBenefit:   parseBoll(r.FormValue("receiveBenefit")),
			StudentId:        uint(id),
			Model: gorm.Model{
				ID:        oldMotherInfo.ID,
				CreatedAt: oldMotherInfo.CreatedAt,
				UpdatedAt: oldMotherInfo.UpdatedAt,
			},
		}

		fmt.Printf("%+v\n", motherInfo)

		db.Save(motherInfo)

		err = template.Must(template.ParseFiles("view/motherInfo.html")).Execute(w, motherInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Get("/student/{id}/aditional-info", func(w http.ResponseWriter, r *http.Request) {
		var aditionalInfo model.AditionalInfo
		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

		db.Where(model.AditionalInfo{StudentId: uint(id)}).First(&aditionalInfo)
		if aditionalInfo.StudentId == 0 {
			aditionalInfo.StudentId = uint(id)
		}
		fmt.Println(aditionalInfo)

		err := template.Must(template.ParseFiles("view/aditionalInfo.html")).Execute(w, aditionalInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Post("/student/{id}/aditional-info", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		parseBoll := func(value string) bool {
			return value == "on"
		}
		parseIncome := func(value string) int {
			response, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return 0
			}
			return int(response)
		}
		var oldAditionalInfo model.AditionalInfo
		db.Where(model.AditionalInfo{StudentId: uint(id)}).First(&oldAditionalInfo)
		if oldAditionalInfo.StudentId == 0 {
			oldAditionalInfo.StudentId = uint(id)
		}

		if oldAditionalInfo.CreatedAt.IsZero() {
			oldAditionalInfo.CreatedAt = time.Now()
		}

		if oldAditionalInfo.UpdatedAt.IsZero() {
			oldAditionalInfo.UpdatedAt = time.Now()
		}

		aditionalInfo := &model.AditionalInfo{
			Income:            parseIncome(r.FormValue("income")),
			ResponsableName:   r.FormValue("responsableName"),
			NisNumber:         r.FormValue("nisNumber"),
			ImageRight:        parseBoll(r.FormValue("imageRight")),
			GovernmentBenefit: parseBoll(r.FormValue("governmentBenefit")),
			HasBrother:        parseBoll(r.FormValue("hasBrother")),
			StudentId:         uint(id),
			Model: gorm.Model{
				ID:        oldAditionalInfo.ID,
				CreatedAt: oldAditionalInfo.CreatedAt,
				UpdatedAt: oldAditionalInfo.UpdatedAt,
			},
		}

		fmt.Printf("%+v\n", aditionalInfo)

		db.Save(aditionalInfo)

		err = template.Must(template.ParseFiles("view/aditionalInfo.html")).Execute(w, aditionalInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":3000", r)
}
