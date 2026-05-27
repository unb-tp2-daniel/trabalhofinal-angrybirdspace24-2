import os
import firebase_admin
from firebase_admin import credentials, firestore

# 1. Autenticação usando o mesmo esquema dinâmico
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

# 2. Conecta no Firestore 
db = firestore.client(database_id="matriculas242")

def main():
    print("=== Populando Coordenadores da UnB ===")
    print("Verificando status do banco de dados...\n")
    
    # 3. Busca os coordenadores que JÁ EXISTEM para criar a trava
    coords_ref = db.collection("coordenadores").stream()
    coords_cadastrados = set() # Usamos um Set (conjunto) para busca ultra-rápida
    
    for doc in coords_ref:
        # Como salvamos o documento com o próprio ID do coordenador, pegamos o doc.id
        coords_cadastrados.add(doc.id)
    
    # 4. Busca os departamentos
    deptos_ref = db.collection("departamentos").stream()
    lista_deptos_pendentes = []
    
    for doc in deptos_ref:
        depto = doc.to_dict()
        coord_id = depto.get("coordenadorId")
        
        # A MÁGICA AQUI: Só adiciona na lista se o ID NÃO estiver nos cadastrados
        if coord_id not in coords_cadastrados:
            lista_deptos_pendentes.append(depto)
            
    # 5. Verifica se a lista de pendentes está vazia
    if not lista_deptos_pendentes:
        print("✅ Todos os departamentos cadastrados já possuem um coordenador!")
        input("Pressione ENTER ou digite 'sair' para encerrar o script: ")
        return # Encerra o programa na hora

    # 6. Ordena a lista pendente numericamente
    lista_deptos_pendentes.sort(key=lambda x: int(x["coordenadorId"]))
    
    print(f"Faltam {len(lista_deptos_pendentes)} departamentos para preencher. Vamos lá:\n")

    # 7. Itera APENAS pela lista pendente
    for depto in lista_deptos_pendentes:
        nome_depto = depto["departamentoNome"]
        coord_id = depto["coordenadorId"]
        depto_id = depto["departamentoId"]
        
        print(f"[{coord_id}] Departamento: {nome_depto} (ID: {depto_id})")
        nome_coord = input("Nome do(a) Coordenador(a) (ou 'sair'): ").strip()
        
        if nome_coord.lower() == 'sair':
            print("\nEncerrando o script...")
            break
            
        if not nome_coord:
            print(">>> Nome vazio ignorado. Pulando para o próximo...\n")
            continue
            
        payload = {
            "coordenadorId": coord_id,
            "departamentoId": depto_id,
            "coordenadorNome": nome_coord
        }
        
        try:
            db.collection("coordenadores").document(coord_id).set(payload)
            print(f"[SUCESSO] {nome_coord} salvo(a) como coordenador(a) de {nome_depto}!\n")
        except Exception as e:
            print(f"[ERRO] Falha ao salvar no banco: {e}\n")

if __name__ == "__main__":
    main()