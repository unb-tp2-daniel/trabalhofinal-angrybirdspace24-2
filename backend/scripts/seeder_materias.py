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

# --- DADOS ESTÁTICOS ---
DEPARTAMENTOS = [
    {"departamentoNome":"Filosofia","departamentoId":"FIL01"}, {"departamentoNome":"Saúde Coletiva","departamentoId":"SC01"},
    {"departamentoNome":"Agronomia","departamentoId":"AGRO01"}, {"departamentoNome":"Educação Fisica","departamentoId":"EDFI01"},
    {"departamentoNome":"Fitopatologia","departamentoId":"FIT01"}, {"departamentoNome":"Línguas Estrangeiras e Tradução","departamentoId":"LET01"},
    {"departamentoNome":"Fisica","departamentoId":"FIS01"}, {"departamentoNome":"Serviço Social","departamentoId":"SS01"},
    {"departamentoNome":"Engenharia Civil e Ambiental","departamentoId":"ECA01"}, {"departamentoNome":"Arquitetura","departamentoId":"ARQ01"},
    {"departamentoNome":"Ecologia","departamentoId":"ECO01"}, {"departamentoNome":"Teoria Literária e Literaturas","departamentoId":"TLL01"},
    {"departamentoNome":"Medicina","departamentoId":"MED01"}, {"departamentoNome":"Veterinária","departamentoId":"VET01"},
    {"departamentoNome":"História","departamentoId":"HIST01"}, {"departamentoNome":"Antropologia","departamentoId":"ANT01"},
    {"departamentoNome":"Administração","departamentoId":"ADM01"}, {"departamentoNome":"Enfermagem","departamentoId":"ENFE01"},
    {"departamentoNome":"Ciências Contábeis","departamentoId":"CC01"}, {"departamentoNome":"Ciência da Informação","departamentoId":"CI01"},
    {"departamentoNome":"Estatística","departamentoId":"EST01"}, {"departamentoNome":"Direito","departamentoId":"DIR01"},
    {"departamentoNome":"Geografia","departamentoId":"GEO01"}, {"departamentoNome":"Mussica","departamentoId":"MUS01"},
    {"departamentoNome":"Sociologia","departamentoId":"SOCIO01"}, {"departamentoNome":"Zoologia","departamentoId":"ZOO01"},
    {"departamentoNome":"Odontologia","departamentoId":"ODONTO01"}, {"departamentoNome":"Matemática","departamentoId":"MAT01"},
    {"departamentoNome":"Educação","departamentoId":"EDU01"}, {"departamentoNome":"Farmacia","departamentoId":"FARMA01"},
    {"departamentoNome":"Botânica","departamentoId":"BOT01"}, {"departamentoNome":"Ciência da Computação","departamentoId":"CIC01"},
    {"departamentoNome":"Desenho Industrial","departamentoId":"DI01"}, {"departamentoNome":"Quimica","departamentoId":"QUI01"},
    {"departamentoNome":"Engenharia de Produção","departamentoId":"ENP01"}, {"departamentoNome":"Genética e Morfologia","departamentoId":"GM01"},
    {"departamentoNome":"Engenharia Elétrica","departamentoId":"ENE01"}, {"departamentoNome":"Engenharia Florestal","departamentoId":"ENF01"},
    {"departamentoNome":"Engenharia Mecânica","departamentoId":"ENM01"}, {"departamentoNome":"Linguística, Português, e Línguas Clássicas","departamentoId":"LPLC01"},
    {"departamentoNome":"Artes Visuais","departamentoId":"AV01"}, {"departamentoNome":"Jornalismo","departamentoId":"JOR01"},
    {"departamentoNome":"Biologia Celular","departamentoId":"BIC01"}, {"departamentoNome":"Artes Cênicas","departamentoId":"AC01"}
]

# --- FUNÇÕES DE APOIO ---

def carregar_dicionario():
    caminho_dic = os.path.join(pasta_do_script, "dicionario_deptos.json")
    try:
        with open(caminho_dic, 'r', encoding='utf-8') as f:
            return json.load(f)
    except FileNotFoundError:
        print("⚠️ Dicionário não encontrado. Usando palavras genéricas para todos.")
        return {"geral": ["Tópicos", "Introdução", "Estudos", "Análise", "Fundamentos", "Prática", "Avançado", "Laboratório", "Seminário"]}

def gerar_nome_materia(depto_id, dicionario):
    gerais = dicionario.get("geral", ["Estudos"])
    especificas = dicionario.get(depto_id, ["Gerais", "Aplicados", "Básicos", "I", "II", "Avançados", "Teóricos"])
    
    # Exemplo: "Fundamentos de Algoritmos" ou "Introdução a Redes Avançados"
    prefixo = random.choice(gerais)
    tema = random.sample(especificas, min(2, len(especificas)))
    
    # Se o prefixo pedir preposição
    conector = " de " if prefixo not in ["Prática", "Estudos"] else " em "
    
    return f"{prefixo}{conector}{' '.join(tema)}"

def gerar_requisitos(ids_base):
    """
    Gera uma lógica aleatória de pré-requisitos usando as 25 matérias base.
    Possibilidades:
    - 1 matéria simples
    - (MatA E MatB)
    - (MatA E MatB) OU (MatC)
    """
    qtd_grupos_ou = random.randint(1, 2)
    requisitos = []
    
    for _ in range(qtd_grupos_ou):
        qtd_and = random.randint(1, 2)
        disciplinas = random.sample(ids_base, qtd_and)
        requisitos.append({"disciplinas": disciplinas})
        
    return requisitos

def rodar_seeder():
    dicionario = carregar_dicionario()
    
    print("🚀 Iniciando o Seeder Massivo de Matérias...")
    total_gerado = 0

    for depto in DEPARTAMENTOS:
        depto_id = depto["departamentoId"]
        nome_depto = depto["departamentoNome"]
        print(f"📂 Gerando 50 matérias para {nome_depto} ({depto_id})...")
        
        # O Batch agrupa as 50 escritas em uma única viagem ao servidor
        batch = db.batch()
        colecao_materias = db.collection("materias")
        
        ids_fase_1 = [] # Guarda os códigos das primeiras 25 para usar como pré-requisitos

        # --- FASE 1: 25 Matérias Base (Sem pré-requisitos) ---
        for i in range(1, 26):
            codigo = f"{depto_id}_{i:04d}"
            ids_fase_1.append(codigo)
            
            payload = {
                "codigo": codigo,
                "nomeMateria": gerar_nome_materia(depto_id, dicionario),
                "departamentoId": depto_id,
                "preRequisitos": [],
                "coRequisitos": [],
                "cargaHoraria": random.choice([30, 45, 60, 60, 90, 120]), # 60 tem mais chance
                "equivalencias": [],
                "conteudo": f"Ementa gerada automaticamente para a disciplina {codigo}.",
                "nivelAcademico": "GRADUACAO"
            }
            
            doc_ref = colecao_materias.document(codigo)
            batch.set(doc_ref, payload)

        # --- FASE 2: 25 Matérias Avançadas (Com pré-requisitos lógicos) ---
        for i in range(26, 51):
            codigo = f"{depto_id}_{i:04d}"
            
            payload = {
                "codigo": codigo,
                "nomeMateria": gerar_nome_materia(depto_id, dicionario) + " Avançado",
                "departamentoId": depto_id,
                "preRequisitos": gerar_requisitos(ids_fase_1),
                "coRequisitos": [], # Podem ficar vazios para simplificar, ou aplicar a mesma lógica
                "cargaHoraria": random.choice([60, 90, 120]),
                "equivalencias": [],
                "conteudo": f"Conteúdo avançado que exige conhecimento prévio das disciplinas base.",
                "nivelAcademico": "GRADUACAO"
            }
            
            doc_ref = colecao_materias.document(codigo)
            batch.set(doc_ref, payload)

        # Comita o batch inteiro do departamento de uma vez
        batch.commit()
        total_gerado += 50
        
    print("-" * 50)
    print(f"🎉 SUUUUCESSO! {total_gerado} matérias foram criadas no banco de dados.")

if __name__ == "__main__":
    rodar_seeder()