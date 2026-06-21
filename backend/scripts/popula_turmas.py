import os
import re
import firebase_admin
from firebase_admin import credentials, firestore

pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

db = firestore.client(database_id="matriculas242")

def carregar_dados_existentes(db):
    print("🔐 Mapeando o banco de dados para garantir integridade...")
    
    # 1. Mapeia Matérias
    materias_ref = db.collection("materias").stream()
    mapa_materias = {doc.id: doc.to_dict().get("nomeMateria") for doc in materias_ref}
    
    # 2. Mapeia Professores (Garante que o professor existe)
    prof_ref = db.collection("professores").stream()
    mapa_professores = {doc.id: doc.to_dict().get("nomeProfessor", doc.to_dict().get("professorNome")) for doc in prof_ref}
    
    # 3. Mapeia Turmas Existentes para controle de ID e choque de horário
    turmas_ref = db.collection("turmas").stream()
    lista_turmas = []
    turmas_ids = set()
    
    for doc in turmas_ref:
        dados = doc.to_dict()
        lista_turmas.append({
            "semestre": dados.get("semestre"),
            "professorid": dados.get("professorid"),
            "horario": dados.get("horario")
        })
        turmas_ids.add(doc.id)
    
    return mapa_materias, mapa_professores, turmas_ids, lista_turmas

def validar_e_parsear_horario(horario_str):
    """
    Valida o formato universitário clássico (Dias + Turno + Horas)
    Exemplos válidos: 24M12, 35T456, 6N12
    Retorna um dicionário com os conjuntos de dias, turno e horas se válido.
    """
    padrao = r"^([2-7]+)([MNT])([1-6]+)$"
    match = re.match(padrao, horario_str.upper().strip())
    if not match:
        return None
        
    return {
        "dias": set(match.group(1)),
        "turno": match.group(2),
        "horas": set(match.group(3))
    }

def ha_choque_de_horario(horario_novo, horario_existente):
    """
    Compara duas strings de horário e diz se elas colidem no mesmo turno.
    """
    h1 = validar_e_parsear_horario(horario_novo)
    h2 = validar_e_parsear_horario(horario_existente)
    
    if not h1 or not h2:
        return False
    
    # Se forem turnos diferentes (ex: um é Manhã 'M' e outro é Tarde 'T'), não choca
    if h1["turno"] != h2["turno"]:
        return False
        
    # Se for o mesmo turno, choca se compartilhar o mesmo dia E a mesma hora/período
    dias_compartilhados = h1["dias"].intersection(h2["dias"])
    horas_compartilhadas = h1["horas"].intersection(h2["horas"])
    
    return len(dias_compartilhados) > 0 and len(horas_compartilhadas) > 0

def get_proximo_codigo_turma(materia_id, turmas_existentes):
    max_num = 0
    prefixo_busca = f"{materia_id}_"
    for turma in turmas_existentes:
        if turma.startswith(prefixo_busca):
            partes = turma.split('_')
            if len(partes) >= 3:
                try:
                    num = int(partes[-1])
                    if num > max_num: max_num = num
                except ValueError: continue
    return f"{materia_id}_{(max_num + 1):02d}"

def main():
    mapa_materias, mapa_professores, turmas_ids, lista_turmas = carregar_dados_existentes(db)
    
    if not mapa_materias:
        print("❌ Erro crítico: Cadastre as matérias primeiro.")
        return
    if not mapa_professores:
        print("❌ Erro crítico: Cadastre os professores na coleção 'professores' antes de rodar.")
        return

    print("\n🚀 VALIDAÇÕES ATIVAS: Formato de horário estrito e prevenção de choque de professores.")

    while True:
        print("=" * 60)
        materia_id = input("ID da Matéria (ou 'sair'): ").strip().upper()
        if materia_id == 'SAIR': break
        if not materia_id: continue
            
        if materia_id not in mapa_materias:
            print(f"❌ Matéria '{materia_id}' não cadastrada.")
            continue
            
        semestre = input("Semestre (ex: 2026.1): ").strip()
        if not semestre: continue

        # --- VALIDAÇÃO DO PROFESSOR ---
        prof_id = input("ID do Professor: ").strip()
        if prof_id not in mapa_professores:
            print(f"❌ Erro: Professor ID '{prof_id}' não existe no banco. Cadastre-o primeiro.")
            continue
        prof_nome = mapa_professores[prof_id]
        print(f"➡️ Professor selecionado: {prof_nome}")

        # --- VALIDAÇÃO DO HORÁRIO ---
        while True:
            horario = input("Horário (padrão ex: 24M12): ").strip().upper()
            if horario.lower() == 'cancelar': break
            
            parsed = validar_e_parsear_horario(horario)
            if not parsed:
                print("❌ Formato INVÁLIDO! Use o padrão de dias, turno e horas (Ex: 24M12, 35T123). Tente novamente ou digite 'cancelar'.")
                continue
                
            # --- VALIDAÇÃO DE CONFLITO/CHOQUE ---
            conflito_detectado = False
            for t in lista_turmas:
                if t["semestre"] == semestre and t["professorid"] == prof_id:
                    if ha_choque_de_horario(horario, t["horario"]):
                        print(f"❌ [CONFLITO] O professor {prof_nome} já está dando aula no horário {t['horario']} neste semestre!")
                        conflito_detectado = True
                        break
            
            if not conflito_detectado:
                break # Horário válido e sem conflitos

        if horario.lower() == 'cancelar':
            print("⚠️ Operação cancelada.")
            continue

        local = input("Local (Sala/Prédio): ").strip()
        try:
            vagas = int(input("Vagas Totais: ").strip())
        except ValueError:
            print("❌ Vagas inválidas. Cancelando registro.")
            continue

        codigo_turma = get_proximo_codigo_turma(materia_id, turmas_ids)

        payload = {
            "codigoTurma": codigo_turma,
            "materiaId": materia_id,
            "nomeMateria": mapa_materias[materia_id],
            "semestre": semestre,
            "vagasTotais": vagas,
            "vagasOcupadas": 0,
            "ativo": True,
            "local": local,
            "horario": horario,
            "vagasExclusivas": {},
            "professorid": prof_id,
            "professorNome": prof_nome,
            "prioridades": {}
        }

        try:
            db.collection("turmas").document(codigo_turma).set(payload)
            print(f"✅ [SUCESSO] Turma {codigo_turma} criada com sucesso!")
            
            # Atualiza o cache local do script para os próximos inputs do loop
            turmas_ids.add(codigo_turma)
            lista_turmas.append({"semestre": semestre, "professorid": prof_id, "horario": horario})
        except Exception as e:
            print(f"❌ Erro ao salvar: {e}")

if __name__ == "__main__":
    main()