import os
import json
import random
import firebase_admin
from firebase_admin import credentials, firestore

# --- CONFIGURAÇÃO DO FIREBASE ---
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

db = firestore.client(database_id="matriculas242")

def carregar_dados_locais():
    """Lê os arquivos JSON e organiza os professores por departamento para busca em O(1)"""
    print("📂 Carregando dados locais...")
    
    with open("materias.json", "r", encoding="utf-8") as f:
        materias = json.load(f)
        
    with open("professores.json", "r", encoding="utf-8") as f:
        professores = json.load(f)
        
    profs_por_depto = {}
    todos_profs = []
    
    for prof in professores:
        depto = prof.get("departamentoid")
        if depto not in profs_por_depto:
            profs_por_depto[depto] = []
        profs_por_depto[depto].append(prof)
        todos_profs.append(prof)
        
    return materias, profs_por_depto, todos_profs

def gerar_horario_universitario():
    """Gera combinações de horários respeitando os limites de 1 a 8"""
    dias = ["24", "35", "46", "234", "56", "2", "3", "4", "5", "6"]
    turnos = ["M", "T", "N"]
    horas = ["12", "34", "56", "78", "1234", "5678"]
    
    dia = random.choice(dias)
    turno = random.choice(turnos)
    hora = random.choice(horas)
    
    return f"{dia}{turno}{hora}"

def rodar_seeder_turmas():
    materias, profs_por_depto, todos_profs = carregar_dados_locais()
    
    print(f"🚀 Iniciando geração de turmas para {len(materias)} matérias...")
    
    batch = db.batch()
    operacoes_no_batch = 0
    total_turmas_geradas = 0
    
    colecao_turmas = db.collection("turmas")
    
    for materia in materias:
        materia_id = materia.get("codigo")
        depto_id = materia.get("departamentoId")
        # Fallback caso o JSON das matérias não tenha o nome embutido
        nome_materia = materia.get("nomeMateria", f"Materia {materia_id}") 
        
        # Filtra professores do mesmo departamento. Se não houver, pega de qualquer departamento.
        opcoes_professores = profs_por_depto.get(depto_id, todos_profs)
        
        for i in range(1, 11): # Gera exatamente 10 turmas (01 a 10)
            codigo_turma = f"{materia_id}_{i:02d}"
            
            prof_escolhido = random.choice(opcoes_professores)
            
            # Vagas: múltiplos de 5 entre 5 e 60
            vagas = random.choice(range(5, 65, 5)) 
            
            # Formatação do Payload seguindo estritamente as tags `firestore:"..."` da struct Go
            payload = {
                "codigoTurma": codigo_turma,
                "materiaId": materia_id,
                "nomeMateria": nome_materia,
                "semestre": "2026.2",
                "vagasTotais": vagas,
                "vagasOcupadas": 0,
                "ativo": True,
                "local": f"Sala {random.randint(1, 600)}",
                "horario": gerar_horario_universitario(),
                "vagasExclusivas": {},
                "professorid": prof_escolhido.get("professorid"),
                "professorNome": prof_escolhido.get("professorNome"),
                "prioridades": {}
            }
            
            # Adiciona a operação ao lote (Batch)
            doc_ref = colecao_turmas.document(codigo_turma)
            batch.set(doc_ref, payload)
            operacoes_no_batch += 1
            total_turmas_geradas += 1
            
            # O Firestore aceita no máximo 500 operações por Batch. 
            # Quando chegar em 500, comita tudo e abre um lote novo.
            if operacoes_no_batch == 500:
                batch.commit()
                print(f"📦 Batch commitado... ({total_turmas_geradas} turmas salvas até agora)")
                batch = db.batch()
                operacoes_no_batch = 0
                
    # Comita o restinho de turmas que sobrou no último lote (se houver)
    if operacoes_no_batch > 0:
        batch.commit()
        
    print("-" * 50)
    print(f"🎉 SUCESSO! {total_turmas_geradas} turmas foram criadas e injetadas no banco de dados.")

if __name__ == "__main__":
    rodar_seeder_turmas()