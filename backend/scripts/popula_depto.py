import os
import firebase_admin
from firebase_admin import credentials, firestore
import datetime

# 1. Autenticação usando o arquivo JSON com caminho dinâmico e seguro
# Descobre a pasta exata onde este script Python está salvo (backend/script)
pasta_do_script = os.path.dirname(os.path.abspath(__file__))

# Constrói o caminho: sai de 'script' (..), sai de 'backend' (..), entra em 'json'
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

# Passa o caminho calculado para o Firebase
cred = credentials.Certificate(caminho_json)
firebase_admin.initialize_app(cred)

# 2. Conecta direto no Firestore
db = firestore.client(database_id="matriculas242")

def gerar_id_departamento(nome):
    sigla = "".join([char for char in nome if char.isupper()])
    if not sigla:
        sigla = nome[:2].upper()
    return f"{sigla}01"

def main():
    print("=== Populando Departamentos UnB (Via Admin SDK) ===")
    
    # Imprime o caminho que está usando para você ter certeza que ele achou o arquivo certo
    print(f"Usando chave em: {os.path.normpath(caminho_json)}\n")
    
    coordenador_serial = 1

    while True:
        nome_depto = input("Nome do Departamento (ou 'sair'): ").strip()

        if nome_depto.lower() == 'sair':
            break
        if not nome_depto:
            continue

        depto_id = gerar_id_departamento(nome_depto)
        coordenador_id = str(coordenador_serial)
        
        # O payload continua sendo o que o seu banco espera lá no Go/Nuxt
        payload = {
            "departamentoId": depto_id,
            "departamentoNome": nome_depto,
            "coordenadorId": coordenador_id,
            "created": datetime.datetime.utcnow().isoformat() + "Z"
        }

        try:
            # 3. Salva DIRETAMENTE na coleção departamentos
            db.collection("departamentos").document(depto_id).set(payload)
            
            print(f"[SUCESSO] {nome_depto} salvo direto no Firestore! (ID: {depto_id})")
            coordenador_serial += 1
            
        except Exception as e:
            print(f"[ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()