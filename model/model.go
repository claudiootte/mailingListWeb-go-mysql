package model

import (
	"html/template"
	"net/http"
)

var View = template.Must(template.ParseGlob("view/*"))

// Conecta-se com o banco de dados, recebe-se todos os contatos existentes e expõe no client

func Home(w http.ResponseWriter, r *http.Request) {
	connectionEstablished := ConnectDB()
	record, err := connectionEstablished.Query(" SELECT * FROM contacts")

	if err != nil {
		panic(err.Error())
	}

	contact := Contact{}
	sliceContact := []Contact{}

	for record.Next() {
		var id int
		var name, email string

		err := record.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		contact.Id = id
		contact.Name = name
		contact.Email = email

		sliceContact = append(sliceContact, contact)
	}

	View.ExecuteTemplate(w, "home", sliceContact)
}

// Função utilizada no template create para a informação de dados para posteriormente serem anexados no banco de dados

func Create(w http.ResponseWriter, r *http.Request) {
	View.ExecuteTemplate(w, "create", nil)
}

// Após passado pelo formulário do create, verifica-se se as informações são válidas e a conexão ainda está ativa, após isso inser-se as informações no banco de dados

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		connectionEstablished := ConnectDB()
		instertRecord, err := connectionEstablished.Prepare("INSERT INTO contacts (name,email) VALUES(?,?) ")

		if err != nil {
			panic(err.Error())
		}

		instertRecord.Exec(name, email)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

// Verifica-se se o id é valido, checa-se as informações relativas ao banco de dados e todos os usuários,

func Edit(w http.ResponseWriter, r *http.Request) {
	idContact := r.URL.Query().Get("id")

	connectionEstablished := ConnectDB()
	record, err := connectionEstablished.Query("SELECT * FROM contacts WHERE id=?", idContact)

	if err != nil {
		panic(err.Error())
	}

	contact := Contact{}

	for record.Next() {
		var id int
		var name, email string

		err := record.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}

		contact.Id = id
		contact.Name = name
		contact.Email = email

	}

	View.ExecuteTemplate(w, "edit", contact)

}

// Após  ter passado pelo template edit, verifica-se se o formulário foi preenchido corretamente e altera-se as devidas informações

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		connectionEstablished := ConnectDB()
		updateRecord, err := connectionEstablished.Prepare(" UPDATE contacts SET name=?,email=? WHERE id=? ")

		if err != nil {
			panic(err.Error())
		}

		updateRecord.Exec(name, email, id)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

// Verifica-se se o id é valido e ainda há conexão com o banco de dados ativos, caso sim exclue-se do banco de dados e redireciona para a página home

func Delete(w http.ResponseWriter, r *http.Request) {
	idContact := r.URL.Query().Get("id")

	connectionEstablished := ConnectDB()
	deleteRecord, err := connectionEstablished.Prepare("DELETE FROM contacts WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteRecord.Exec(idContact)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
