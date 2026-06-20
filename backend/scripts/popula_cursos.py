import json
import requests
import time

URL = "https://criarcurso-nudpqdvptq-rj.a.run.app"
ARQUIVO_JSON = "cursos_gerados_v2.json"

with open(ARQUIVO_JSON, "r", encoding="utf-8") as f:
    cursos = json.load(f)

sucessos = 0
falhas = 0

for curso in cursos:
    try:
        resposta = requests.post(
            URL,
            json=curso,
            headers={"Content-Type": "application/json"},
            timeout=15
        )

        if resposta.status_code in (200, 201):
            print(f"[OK] {curso['cursoid']}")
            sucessos += 1
        else:
            print(f"[ERRO] {curso['cursoid']} -> {resposta.status_code}: {resposta.text}")
            falhas += 1

    except Exception as e:
        print(f"[EXCEPTION] {curso['cursoid']} -> {e}")
        falhas += 1

    time.sleep(0.1)

print()
print(f"Sucessos: {sucessos}")
print(f"Falhas: {falhas}")