import requests

url = "https://southamerica-east1-matriculas242.cloudfunctions.net/ListarCoordenadores"

response = requests.get(url)
response.raise_for_status()  # Lança exceção se houver erro HTTP

dados = response.json()

for aluno in dados:
    nome = aluno.get("coordenadorNome")
    matricula = "coord"+aluno.get("coordenadorId")
    senha = matricula
    nome = nome.replace("á", "a").replace("ã", "a").replace("é", "e").replace("ê", "e").replace("í", "i").replace("ó", "o").replace("ô", "o")
    nome = nome.split()
    email = nome[0].lower()+"."+ nome[-1].lower() + "@email.com"

    payload = {
        "matricula": matricula,
        "email": email,
        "senha": senha
    }

    try:
        resposta = requests.post(
            "https://southamerica-east1-matriculas242.cloudfunctions.net/CriarUsuario",
            json=payload,
            headers={
                "Content-Type": "application/json"
            },
            timeout=30
        )

        print(
            f"Coordenador: {nome} ({matricula}) - "
            f"Status: {resposta.status_code}"
        )

        print(resposta.text)

    except requests.exceptions.RequestException as e:
        print(f"Erro ao criar usuário {matricula}: {e}")