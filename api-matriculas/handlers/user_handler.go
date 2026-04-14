package handlers
import ("encoding/json"; "net/http"; "servidor-api/models")
var Users []models.User
func HandleFunc(w http.ResponseWriter, r *http.Request) 
{
		if r.Method == http.MethodGet {
			msg := "Olá, " + users.Nome + "! Bem-vindo ao servidor API."
			w.Write([]byte(msg))
		} else if r.Method == http.MethodPost {
			err := json.NewDecoder(r.Body).Decode(&users)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Error decoding JSON"))
				return
			}
			w.Write([]byte("POST request received"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
}