import os
import firebase_admin
from firebase_admin import credentials, firestore

# --- CONFIGURAÇÃO DO FIREBASE ---
# Mantendo a mesma estrutura de pastas do seu script original
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

# Usando o mesmo database_id que você definiu antes
db = firestore.client(database_id="matriculas242")

def aplicar_curativo():
    print("🔍 Mapeando professores...")
    
    # 1. Carrega todos os professores e extrai o 'professorNome'
    prof_ref = db.collection("professores").stream()
    mapa_professores = {}
    
    for doc in prof_ref:
        dados = doc.to_dict()
        # Conforme seu model Go: firestore:"professorNome"
        nome = dados.get("professorNome") 
        
        if nome:
            # Assumindo que o ID do documento é o mesmo salvo em professorid
            mapa_professores[doc.id] = nome
            
            # Por garantia, também mapeia usando a chave de dentro do documento
            prof_id_interno = dados.get("professorid")
            if prof_id_interno:
                mapa_professores[prof_id_interno] = nome

    print(f"✅ {len(set(mapa_professores.values()))} professores carregados para o cache.")
    print("\n🛠️ Aplicando band-aid nas turmas...\n")

    # 2. Busca todas as turmas
    turmas_ref = db.collection("turmas").stream()
    
    atualizadas = 0
    ignoradas = 0

    for doc in turmas_ref:
        dados = doc.to_dict()
        turma_id = doc.id
        
        # Conforme seu model Go: firestore:"professorid"
        prof_id = dados.get("professorid")
        nome_atual = dados.get("professorNome")
        
        if not prof_id:
            print(f"⚠️ Turma {turma_id} não tem 'professorid' definido. Pulando.")
            ignoradas += 1
            continue
            
        nome_correto = mapa_professores.get(prof_id)
        
        if not nome_correto:
            print(f"⚠️ Não encontrei o professor de ID '{prof_id}' no banco. Pulando turma {turma_id}.")
            ignoradas += 1
            continue
            
        # 3. Verifica se precisa de correção (está null, vazio ou errado)
        if nome_atual != nome_correto:
            try:
                # Faz o update apenas do campo específico sem alterar o resto da turma
                db.collection("turmas").document(turma_id).update({
                    "professorNome": nome_correto
                })
                print(f"🩹 Corrigido {turma_id}: {nome_atual} -> {nome_correto}")
                atualizadas += 1
            except Exception as e:
                print(f"❌ Erro ao atualizar turma {turma_id}: {e}")
        else:
            ignoradas += 1 # O nome já está certinho

    print("-" * 40)
    print(f"🎉 Resumo da operação:")
    print(f"   Turmas corrigidas: {atualizadas}")
    print(f"   Turmas ignoradas (já estavam OK): {ignoradas}")

if __name__ == "__main__":
    aplicar_curativo()