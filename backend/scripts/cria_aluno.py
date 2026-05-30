import requests

# A URL da sua Cloud Function
url = "https://criaraluno-nudpqdvptq-rj.a.run.app"

# Os dados que você quer enviar para a função
dados = {
	"matricula":"123456",
	"nomeAluno": "João da Silva",
	"cursoId": "CURSO001",
	"ativo": True,
	"semestre": "2023.1",
	"materiasConcluidas": {
		"DISC001": "Aprovada",
		"DISC002": "Aprovada"
	},
	"ira": 8.5,
	"prioridades": {
		"DISC003": "Alta",
		"DISC004": "Média"
	},
	"nivelAcademico": "Graduacao",
	"horas": 120
}

try:
    # Fazendo a requisição POST com os dados em formato JSON
    resposta = requests.post(url, json=dados)
    print(resposta.text)
    # Verifica se a requisição deu certo (Status 200)
    resposta.raise_for_status()
    
    # Se a função retornar um JSON, você captura assim:
    resultado = resposta.json()
    print("Sucesso! Resposta da Function:", resultado)

except requests.exceptions.HTTPError as err:
    print(status_code=f"Erro HTTP: {err}")
except Exception as e:
    print(f"Ocorreu um erro: {e}")