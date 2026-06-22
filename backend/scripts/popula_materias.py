import os
import sys
import firebase_admin
from firebase_admin import credentials, firestore

# --- CONFIGURAÇÃO DO FIREBASE ---
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

db = firestore.client(database_id="matriculas242")

# --- FUNÇÕES DE VALIDAÇÃO EM TEMPO REAL ---

def verificar_departamento_live(depto_id):
    """
    Busca no banco se o departamento existe no momento exato da digitação.
    Retorna o nome do departamento se existir, ou None.
    """
    # Baseado no seu script anterior que busca pela chave interna 'departamentoId'
    query = db.collection("departamentos").where("departamentoId", "==", depto_id).limit(1).stream()
    docs = list(query)
    
    if not docs:
        return None
    return docs[0].to_dict().get("departamentoNome", "Nome Não Definido")

def get_proximo_codigo_materia_live(depto_id):
    """
    Consulta todas as matérias do departamento no banco para garantir 
    que o gerador de ID nunca sofra descompasso.
    """
    docs = db.collection("materias").where("departamentoId", "==", depto_id).stream()
    max_num = 0
    prefixo = f"{depto_id}_"
    
    for doc in docs:
        mat_id = doc.id
        if mat_id.startswith(prefixo):
            try:
                num = int(mat_id.split('_')[1])
                if num > max_num:
                    max_num = num
            except (ValueError, IndexError):
                continue
                
    return f"{depto_id}_{(max_num + 1):04d}"

def parse_e_validar_requisitos_live(entrada_str):
    """
    Lê a string de requisitos, formata corretamente e vai no Firestore
    checar se CADA matéria citada realmente existe no banco.
    """
    if not entrada_str.strip() or entrada_str.lower() == 'nenhum':
        return [], None
        
    matriz_requisitos = []
    grupos_ou = entrada_str.upper().split(' OU ')
    
    for grupo in grupos_ou:
        # Filtra espaços vazios caso o usuário digite vírgulas sobrando (ex: "CIC01, , CIC02")
        materias_and = [m.strip() for m in grupo.split(',') if m.strip()]
        
        if not materias_and:
            continue

        for codigo in materias_and:
            # Live Check: A matéria realmente existe agora?
            doc = db.collection("materias").document(codigo).get()
            if not doc.exists:
                return None, codigo # Retorna qual matéria falhou na validação
                
        matriz_requisitos.append({"disciplinas": materias_and})
        
    return matriz_requisitos, None

def parse_e_validar_equivalencias_live(entrada_str):
    """
    Garante que matérias declaradas como equivalentes também existam no banco.
    """
    if not entrada_str.strip() or entrada_str.lower() == 'nenhum':
        return [], None
        
    equivalencias = [e.strip() for e in entrada_str.upper().split(',') if e.strip()]
    
    for codigo in equivalencias:
        doc = db.collection("materias").document(codigo).get()
        if not doc.exists:
            return None, codigo
            
    return equivalencias, None

# --- FLUXO PRINCIPAL ---

def main():
    print("=" * 60)
    print("=== Cadastro On-Demand Robusto de Matérias (Graduação) ===")
    print("=" * 60)
    print("\nINSTRUÇÕES:")
    print("- Digite o Código do departamento (ex: CIC01, ENE01).")
    print("- As matérias receberão IDs automáticos consultados em tempo real.")
    print("- Digite 'ok' no nome da matéria para trocar de departamento.")
    print("- Digite 'sair' a qualquer momento para encerrar.\n")

    try:
        while True:
            print("-" * 60)
            depto_id = input("🔎 Digite o Código do Departamento (ou 'sair'): ").strip().upper()
            
            if depto_id == 'SAIR':
                print("\nEncerrando o script com segurança...")
                break
            if not depto_id:
                continue
                
            nome_depto = verificar_departamento_live(depto_id)
            if not nome_depto:
                print(f"❌ Erro: Departamento '{depto_id}' não localizado no banco de dados.")
                continue
                
            print(f"\n📂 DEPARTAMENTO: {nome_depto} (ID: {depto_id})")
            print("=" * 60)
            
            while True:
                codigo_materia = get_proximo_codigo_materia_live(depto_id)
                print(f"\n📝 Próximo ID reservado: {codigo_materia}")
                
                nome_materia = input("Nome da Matéria (ou 'ok'/'sair'): ").strip()
                
                if nome_materia.lower() == 'sair':
                    print("\nEncerrando o script com segurança...")
                    return
                if nome_materia.lower() == 'ok':
                    print(f">>> Saindo do departamento {nome_depto}...")
                    break
                if not nome_materia:
                    print("⚠️ O nome da matéria não pode ficar vazio.")
                    continue

                # --- Validação de Carga Horária ---
                while True:
                    try:
                        carga_input = input("Carga Horária (ex: 60): ").strip()
                        carga_horaria = int(carga_input)
                        if carga_horaria <= 0:
                            print("⚠️ A carga horária deve ser maior que zero.")
                            continue
                        break
                    except ValueError:
                        print("❌ Erro: Digite apenas números inteiros para a carga horária.")

                # --- Validação de Pré-requisitos ---
                while True:
                    pre_req_input = input("Pré-requisitos (separados por ',' ou 'OU', ou 'nenhum'): ")
                    matriz_pre, erro = parse_e_validar_requisitos_live(pre_req_input)
                    if erro:
                        print(f"❌ Erro de Integridade: O pré-requisito '{erro}' NÃO existe no banco de dados. Cadastre-o primeiro ou corrija a digitação.")
                    else:
                        break 

                # --- Validação de Co-requisitos ---
                while True:
                    co_req_input = input("Co-requisitos (separados por ',' ou 'OU', ou 'nenhum'): ")
                    matriz_co, erro = parse_e_validar_requisitos_live(co_req_input)
                    if erro:
                        print(f"❌ Erro de Integridade: O co-requisito '{erro}' NÃO existe no banco de dados. Cadastre-o primeiro ou corrija a digitação.")
                    else:
                        break

                # --- Validação de Equivalências ---
                while True:
                    equiv_input = input("Equivalências (separadas por vírgula, ou 'nenhum'): ")
                    lista_equiv, erro = parse_e_validar_equivalencias_live(equiv_input)
                    if erro:
                        print(f"❌ Erro de Integridade: A equivalência '{erro}' NÃO existe no banco de dados. Cadastre-a primeiro ou corrija a digitação.")
                    else:
                        break

                # --- Validação de Ementa ---
                while True:
                    conteudo = input("Ementa/Conteúdo: ").strip()
                    if not conteudo:
                        print("⚠️ A ementa não pode ficar vazia. Descreva o conteúdo da matéria.")
                        continue
                    break

                # --- Montagem do Payload e Persistência ---
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
                    print(f"✅ [SUCESSO] {codigo_materia} - '{nome_materia}' cadastrada com integridade total!")
                except Exception as e:
                    print(f"❌ [ERRO CRÍTICO] Falha de comunicação ao salvar no Firebase: {e}")

    except KeyboardInterrupt:
        # Permite que o usuário feche com Ctrl+C sem estourar aquele bloco gigante de erro na tela
        print("\n\n⚠️ Operação interrompida pelo usuário (Ctrl+C). Encerrando segurança...")
        sys.exit(0)

if __name__ == "__main__":
    main()