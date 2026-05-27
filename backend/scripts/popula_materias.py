import os
import firebase_admin
from firebase_admin import credentials, firestore

pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

db = firestore.client(database_id="matriculas242")

def carregar_dados_existentes(db):
    print("Mapeando departamentos e matérias para garantir integridade...")
    
    deptos_ref = db.collection("departamentos").stream()
    mapa_deptos = {doc.to_dict()["departamentoId"]: doc.to_dict()["departamentoNome"] for doc in deptos_ref}
    
    materias_ref = db.collection("materias").stream()
    materias_existentes = {doc.id for doc in materias_ref}
    
    return mapa_deptos, materias_existentes

def get_proximo_codigo_materia(depto_id, materias_existentes):
    max_num = 0
    prefixo_busca = f"{depto_id}_"
    
    for mat in materias_existentes:
        if mat.startswith(prefixo_busca):
            num_str = mat.split('_')[1]
            try:
                num = int(num_str)
                if num > max_num:
                    max_num = num
            except (ValueError, IndexError):
                continue
                
    return f"{depto_id}_{(max_num + 1):04d}"

def parse_requisitos(entrada_str, materias_existentes):
    if not entrada_str.strip() or entrada_str.lower() == 'nenhum':
        return [], None
        
    matriz_requisitos = []
    grupos_ou = entrada_str.upper().split(' OU ')
    
    for grupo in grupos_ou:
        materias_and = [m.strip() for m in grupo.split(',')]
        for codigo in materias_and:
            if codigo not in materias_existentes:
                return None, codigo 
                
        # A MÁGICA: Empacota a lista num dicionário para o Firestore aceitar
        matriz_requisitos.append({"disciplinas": materias_and})
        
    return matriz_requisitos, None

def main():
    print("=== Cadastro On-Demand de Matérias (Graduação) ===")
    
    mapa_deptos, materias_existentes = carregar_dados_existentes(db)
    
    if not mapa_deptos:
        print("Erro crítico: Nenhum departamento no banco!")
        return

    print("\nINSTRUÇÕES:")
    print("- Digite o Código do departamento (ex: CIC01, ENE01).")
    print("- As matérias receberão IDs automáticos (ex: CIC01_0001).")
    print("- Digite 'ok' no nome da matéria para trocar de departamento.")
    print("- Digite 'sair' a qualquer momento para encerrar o script.\n")

    while True:
        print("=" * 60)
        depto_id = input("🔎 Digite o Código do Departamento (ou 'sair'): ").strip().upper()
        
        if depto_id == 'SAIR':
            print("\nEncerrando o script com segurança...")
            break
        if not depto_id:
            continue
            
        if depto_id not in mapa_deptos:
            print(f"❌ Departamento '{depto_id}' não existe no banco.")
            continue
            
        nome_depto = mapa_deptos[depto_id]
        
        print(f"\n📂 DEPARTAMENTO: {nome_depto} (ID: {depto_id})")
        print("-" * 60)
        
        while True:
            codigo_materia = get_proximo_codigo_materia(depto_id, materias_existentes)
            
            print(f"\n📝 Nova Matéria será: {codigo_materia}")
            nome_materia = input("Nome da Matéria (ou 'ok'/'sair'): ").strip()
            
            if nome_materia.lower() == 'sair':
                print("\nEncerrando o script com segurança...")
                return
            if nome_materia.lower() == 'ok':
                print(f">>> Saindo de {nome_depto}. Voltando à pesquisa de departamentos...")
                break
            if not nome_materia:
                continue

            try:
                carga_horaria = int(input("Carga Horária (ex: 60): ").strip())
            except ValueError:
                print("❌ Erro: Carga horária deve ser um número. Abortando esta matéria.")
                continue

            while True:
                pre_req_input = input("Pré-requisitos (separados por ',' ou 'OU', ex: CIC01_0001): ")
                matriz_pre, erro = parse_requisitos(pre_req_input, materias_existentes)
                if erro:
                    print(f"❌ Erro: Pré-requisito '{erro}' não cadastrado. Cadastre primeiro ou corrija.")
                else:
                    break 

            while True:
                co_req_input = input("Co-requisitos (separados por ',' ou 'OU'): ")
                matriz_co, erro = parse_requisitos(co_req_input, materias_existentes)
                if erro:
                    print(f"❌ Erro: Co-requisito '{erro}' não cadastrado. Cadastre primeiro ou corrija.")
                else:
                    break

            equivalencias = input("Equivalências (separadas por vírgula, ou vazio): ").strip().upper()
            lista_equiv = [e.strip() for e in equivalencias.split(',')] if equivalencias else []
            
            conteudo = input("Ementa/Conteúdo: ").strip()

            payload = {
                "codigo": codigo_materia,
                "nomeMateria": nome_materia,
                "departamentoId": depto_id,
                "preRequisitos": matriz_pre,
                "coRequisitos": matriz_co,
                "cargaHoraria": carga_horaria,
                "equivalencias": lista_equiv,
                "conteudo": conteudo,
                "nivelAcademico": "GRADUACAO" 
            }

            try:
                db.collection("materias").document(codigo_materia).set(payload)
                print(f"✅ [SUCESSO] {codigo_materia} - {nome_materia} salva no banco!")
                materias_existentes.add(codigo_materia)
                
            except Exception as e:
                print(f"❌ [ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()