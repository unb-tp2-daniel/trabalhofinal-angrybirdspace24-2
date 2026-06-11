import requests

url = "https://listaralunos-nudpqdvptq-rj.a.run.app/"

response = requests.get(url)
response.raise_for_status()  # Lança exceção se houver erro HTTP

dados = response.json()

for aluno in dados:
    nome = aluno.get("nomeAluno")
    matricula = aluno.get("matricula")
    senha = matricula
    nome = nome.replace("á", "a").replace("ã", "a").replace("é", "e").replace("ê", "e").replace("í", "i").replace("ó", "o").replace("ô", "o")
    email = nome.lower().replace(" ", ".") + "@email.com"

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
            f"Aluno: {nome} ({matricula}) - "
            f"Status: {resposta.status_code}"
        )

        print(resposta.text)

    except requests.exceptions.RequestException as e:
        print(f"Erro ao criar usuário {matricula}: {e}")