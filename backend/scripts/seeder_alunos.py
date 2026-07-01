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

# --- BANCO DE NOMES ALEATÓRIOS ---
NOMES = ["Ana", "Bruno", "Carlos", "Daniela", "Eduardo", "Fernanda", "Gabriel", "Helena", "Igor", "Julia", "Lucas", "Mariana", "Nicolas", "Olivia", "Pedro", "Quintino", "Rafael", "Sofia", "Tiago", "Ursula", "Vitor", "Yasmin", "João", "Marcela", "Danilo"]
SOBRENOMES = ["Silva", "Santos", "Oliveira", "Souza", "Rodrigues", "Ferreira", "Alves", "Pereira", "Lima", "Gomes", "Costa", "Ribeiro", "Martins", "Carvalho", "Almeida", "Lopes", "Soares", "Fernandes", "Vieira", "Barbosa"]
MENCOES_APROVACAO = ["SS", "MS", "MM"]

def gerar_nome_aleatorio():
    return f"{random.choice(NOMES)} {random.choice(SOBRENOMES)} {random.choice(SOBRENOMES)}"

def carregar_dados_robusto():
    print("📂 Carregando bases de dados...")
    
    with open("departamentos.json", "r", encoding="utf-8") as f:
        deptos_raw = json.load(f)
        
    if isinstance(deptos_raw, str):
        try:
            deptos_raw = json.loads(deptos_raw)
        except Exception:
            pass

    if isinstance(deptos_raw, dict):
        for key, value in deptos_raw.items():
            if isinstance(value, list):
                deptos_raw = value
                break
        else:
            deptos_raw = list(deptos_raw.values())

    with open("materias.json", "r", encoding="utf-8") as f:
        materias_raw = json.load(f)

    if isinstance(materias_raw, str):
        try:
            materias_raw = json.loads(materias_raw)
        except Exception:
            pass

    if isinstance(materias_raw, dict):
        for key, value in materias_raw.items():
            if isinstance(value, list):
                materias_raw = value
                break
        else:
            materias_raw = list(materias_raw.values())

    materias_por_depto = {}
    for mat in materias_raw:
        if not isinstance(mat, dict):
            continue 
            
        depto_id = mat.get("departamentoId")
        if depto_id not in materias_por_depto:
            materias_por_depto[depto_id] = []
        materias_por_depto[depto_id].append(mat)
        
    return deptos_raw, materias_por_depto

def pode_cursar(materia, concluidas_set):
    reqs = materia.get("preRequisitos")
    if not reqs:
        return True
        
    for grupo in reqs:
        disciplinas_do_and = grupo.get("disciplinas", [])
        if all(disc in concluidas_set for disc in disciplinas_do_and):
            return True
            
    return False

def simular_historico_academico(materias_curso):
    concluidas_map = {}
    reprovadas_map = {}
    concluidas_set = set() 
    
    qtd_semestres = random.randint(1, 10) 
    
    for _ in range(qtd_semestres):
        disponiveis = [m for m in materias_curso if m.get("codigo") not in concluidas_set and pode_cursar(m, concluidas_set)]
        
        if not disponiveis:
            break
            
        qtd_pegar = random.randint(1, min(6, len(disponiveis)))
        cursando_agora = random.sample(disponiveis, qtd_pegar)
        
        for mat in cursando_agora:
            codigo = mat.get("codigo")
            if not codigo: continue
            
            if random.random() <= 0.85:
                concluidas_set.add(codigo)
                concluidas_map[codigo] = random.choice(MENCOES_APROVACAO) 
            else:
                reprovadas_map[codigo] = reprovadas_map.get(codigo, 0) + 1
                
    return concluidas_map, reprovadas_map

def main():
    print("=" * 60)
    print("=== Gerador Automático Massivo de Alunos ===")
    print("=" * 60)
    
    deptos, materias_por_depto = carregar_dados_robusto()
    colecao_alunos = db.collection("alunos")
    
    batch = db.batch()
    operacoes = 0
    total_gerados = 0
    
    for depto in deptos:
        if not isinstance(depto, dict):
            continue
            
        depto_id = depto.get("departamentoId")
        if not depto_id:
            continue
            
        materias_curso = materias_por_depto.get(depto_id, [])
        print(f"🎓 Processando: {depto.get('departamentoNome', depto_id)} ({depto_id}) - Injetando 10 alunos...")
        
        # --- CASO ESPECIAL: ALUNO FIXO DO CIC ---
        if depto_id == "CIC01":
            matricula_esp = "Unb_242010102"
            concluidas_esp, reprovadas_esp = simular_historico_academico(materias_curso)
            
            payload_esp = {
                "matricula": matricula_esp,
                "nomeAluno": "Guilherme Silva Cavalcante",
                "cursoId": depto_id,
                "ativo": True,
                "semestre": "26.2",
                "materiasConcluidas": concluidas_esp,
                "materiasReprovadas": reprovadas_esp,
                "ira": round(random.uniform(4.0, 5.0), 4),
                "prioridades": {
                    "EhDACES": "nao",
                    "integralizado": str(random.randint(40, 90)),
                    "turno": "Matutino"
                },
                "nivelAcademico": "GRADUACAO",
                "horas": random.randint(1000, 3210)
            }
            
            batch.set(colecao_alunos.document(matricula_esp), payload_esp)
            operacoes += 1
            total_gerados += 1
            print("   🌟 Aluno específico do CIC criado com sucesso!")

        # --- LOOP DE GERAÇÃO NORMAL (10 ALUNOS) ---
        for _ in range(5):
            # Gera os 5 dígitos finais aleatórios e concatena no prefixo 2420
            numero_aleatorio = random.randint(10000, 99999)
            matricula = f"Unb_2420{numero_aleatorio}"
            
            # Prevenção simples para não colidir acidentalmente com a matrícula fixa
            while matricula == "Unb_242010102":
                numero_aleatorio = random.randint(10000, 99999)
                matricula = f"Unb_2420{numero_aleatorio}"
            
            concluidas, reprovadas = simular_historico_academico(materias_curso)
            
            payload = {
                "matricula": matricula,
                "nomeAluno": gerar_nome_aleatorio(),
                "cursoId": depto_id,
                "ativo": True,
                "semestre": "26.2",
                "materiasConcluidas": concluidas,
                "materiasReprovadas": reprovadas,
                "ira": round(random.uniform(0.0, 5.0), 4),
                "prioridades": {
                    "EhDACES": random.choice(["sim", "nao"]),
                    "integralizado": str(random.randint(0, 100)),
                    "turno": random.choice(["Matutino", "Vespertino", "Noturno"])
                },
                "nivelAcademico": "GRADUACAO",
                "horas": random.randint(0, 3210)
            }
            
            doc_ref = colecao_alunos.document(matricula)
            batch.set(doc_ref, payload)
            
            operacoes += 1
            total_gerados += 1
            
            if operacoes >= 500:
                batch.commit()
                batch = db.batch()
                operacoes = 0
                
    if operacoes > 0:
        batch.commit()
        
    print("-" * 60)
    print(f"✅ [OPERAÇÃO CONCLUÍDA] {total_gerados} alunos foram matriculados no banco de dados!")

if __name__ == "__main__":
    main()