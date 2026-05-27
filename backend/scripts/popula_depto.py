import os
import firebase_admin
from firebase_admin import credentials, firestore
import datetime

# 1. Autenticação usando o arquivo JSON com caminho dinâmico e seguro
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
firebase_admin.initialize_app(cred)

# 2. Conecta direto no Firestore
db = firestore.client(database_id="matriculas242")

def gerar_id_departamento(nome):
    # Usa o nome cru (ex: "CIência da Computação") para pescar as maiúsculas
    sigla = "".join([char for char in nome if char.isupper()])
    if not sigla:
        sigla = nome[:2].upper()
    return f"{sigla}01"

def formatar_nome_departamento(nome):
    # Mantém a letra original se for o índice 0 ou se vier logo após um espaço.
    # Força todo o resto a ser minúsculo.
    resultado = ""
    for i, char in enumerate(nome):
        if i == 0 or nome[i-1] == ' ':
            resultado += char
        else:
            resultado += char.lower()
    return resultado

def main():
    print("=== Populando Departamentos UnB (Via Admin SDK) ===")
    print(f"Usando chave em: {os.path.normpath(caminho_json)}\n")
    
    coordenador_serial = 1

    while True:
        # Pega a string exatamente como você digitou (o "hack" das maiúsculas)
        nome_cru = input("Nome do Departamento (ou 'sair'): ").strip()

        if nome_cru.lower() == 'sair':
            break
        if not nome_cru:
            continue

        # 1º Passo: Gera o ID extraindo as maiúsculas do nome cru -> "CIC01"
        depto_id = gerar_id_departamento(nome_cru)
        
        # 2º Passo: Limpa o nome corrigindo as letras no meio da palavra -> "Ciência da Computação"
        nome_formatado = formatar_nome_departamento(nome_cru)
        
        coordenador_id = str(coordenador_serial)
        
        # O payload agora usa a variável nome_formatado
        payload = {
            "departamentoId": depto_id,
            "departamentoNome": nome_formatado,
            "coordenadorId": coordenador_id,
            "created": datetime.datetime.utcnow().isoformat() + "Z"
        }

        try:
            # 3. Salva DIRETAMENTE na coleção departamentos
            db.collection("departamentos").document(depto_id).set(payload)
            
            print(f"[SUCESSO] '{nome_formatado}' salvo direto no Firestore! (ID: {depto_id})")
            coordenador_serial += 1
            
        except Exception as e:
            print(f"[ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()