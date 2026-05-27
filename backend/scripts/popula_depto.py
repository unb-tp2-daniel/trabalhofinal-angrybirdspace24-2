import os
import firebase_admin
from firebase_admin import credentials, firestore
import datetime

# 1. Autenticação usando o arquivo JSON com caminho dinâmico e seguro
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

# 2. Conecta direto no Firestore
db = firestore.client(database_id="matriculas242")

def gerar_sigla_base(nome):
    # Extrai apenas as maiúsculas para formar a base (ex: "CIência da Computação" -> "CIC")
    sigla = "".join([char for char in nome if char.isupper()])
    if not sigla:
        sigla = nome[:2].upper()
    return sigla

def get_depto_id_unico(db, sigla_base):
    # Trava 1: Garante que o ID do departamento não exista no banco.
    # Se CIC01 existir, ele tenta CIC02, CIC03, etc.
    contador = 1
    while True:
        # Formata o contador com 2 dígitos (01, 02, etc.)
        depto_id = f"{sigla_base}{contador:02d}"
        doc_ref = db.collection("departamentos").document(depto_id)
        
        # Se o documento NÃO existir, achamos um ID livre!
        if not doc_ref.get().exists:
            return depto_id
        
        contador += 1

def get_proximo_coordenador_id(db):
    # Trava 2: Varre o banco para achar o maior coordenadorId já cadastrado
    deptos = db.collection("departamentos").stream()
    maior_id = 0
    
    for doc in deptos:
        dados = doc.to_dict()
        if "coordenadorId" in dados:
            try:
                c_id = int(dados["coordenadorId"])
                if c_id > maior_id:
                    maior_id = c_id
            except ValueError:
                continue
                
    # Retorna o maior encontrado + 1
    return maior_id + 1

def formatar_nome_departamento(nome):
    resultado = ""
    for i, char in enumerate(nome):
        if i == 0 or nome[i-1] == ' ':
            resultado += char
        else:
            resultado += char.lower()
    return resultado

def main():
    print("=== Populando Departamentos UnB (Via Admin SDK) ===")
    print("Buscando integridade atual do banco...\n")
    
    # Descobre automaticamente de onde começar a contar
    coordenador_serial = get_proximo_coordenador_id(db)
    print(f"Próximo Coordenador ID disponível: {coordenador_serial}\n")

    while True:
        nome_cru = input("Nome do Departamento (ou 'sair'): ").strip()

        if nome_cru.lower() == 'sair':
            break
        if not nome_cru:
            continue
            
        nome_formatado = formatar_nome_departamento(nome_cru)
        
        # Gera a sigla base e busca um ID livre no banco
        sigla_base = gerar_sigla_base(nome_cru)
        depto_id = get_depto_id_unico(db, sigla_base)
        
        coordenador_id = str(coordenador_serial)
        
        payload = {
            "departamentoId": depto_id,
            "departamentoNome": nome_formatado,
            "coordenadorId": coordenador_id,
            "created": datetime.datetime.utcnow().isoformat() + "Z"
        }

        try:
            db.collection("departamentos").document(depto_id).set(payload)
            print(f"[SUCESSO] '{nome_formatado}' salvo! (Depto ID: {depto_id} | Coord ID: {coordenador_id})")
            
            # Só incrementa o serial se a gravação no banco deu certo
            coordenador_serial += 1
            
        except Exception as e:
            print(f"[ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()