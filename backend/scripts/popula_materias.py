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
    # Para o batch, guardamos a lista inteira de dicts (já ordenada pelo nome para facilitar)
    lista_deptos = [doc.to_dict() for doc in deptos_ref]
    lista_deptos.sort(key=lambda x: x.get("departamentoNome", ""))
    
    materias_ref = db.collection("materias").stream()
    materias_existentes = {doc.id for doc in materias_ref}
    
    return lista_deptos, materias_existentes

def get_proximo_codigo_materia(depto_id, materias_existentes):
    """
    Busca o maior número cadastrado para um depto_id (ex: CIC01) e retorna o próximo (ex: CIC01_0001).
    """
    max_num = 0
    prefixo_busca = f"{depto_id}_"
    
    for mat in materias_existentes:
        # Busca matérias que começam com "CIC01_"
        if mat.startswith(prefixo_busca):
            # Extrai apenas a parte depois do underscore
            num_str = mat.split('_')[1]
            try:
                num = int(num_str)
                if num > max_num:
                    max_num = num
            except (ValueError, IndexError):
                continue
                
    # Retorna o ID do depto + "_" + serial formatado com 4 casas
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
        matriz_requisitos.append(materias_and)
        
    return matriz_requisitos, None

def main():
    print("=== Cadastro em Batch de Matérias (Graduação) ===")
    
    lista_deptos, materias_existentes = carregar_dados_existentes(db)
    
    if not lista_deptos:
        print("Erro crítico: Nenhum departamento no banco!")
        return

    print("\nINSTRUÇÕES:")
    print("- As matérias receberão IDs automáticos (ex: CIC01_0001).")
    print("- Digite 'ok' no nome da matéria para pular para o próximo departamento.")
    print("- Digite 'sair' a qualquer momento para encerrar o script de vez.\n")

    # --- Loop Externo: Itera pelos Departamentos ---
    for depto in lista_deptos:
        nome_depto = depto["departamentoNome"]
        depto_id = depto["departamentoId"]
        
        print(f"\n{'='*60}")
        print(f"🏢 DEPARTAMENTO: {nome_depto} (ID: {depto_id})")
        print(f"{'='*60}")
        
        # --- Loop Interno: Cadastro de Matérias ---
        while True:
            # Calcula o serial dinamicamente (ex: CIC01_0001)
            codigo_materia = get_proximo_codigo_materia(depto_id, materias_existentes)
            
            print(f"\n📝 Nova Matéria será: {codigo_materia}")
            nome_materia = input("Nome da Matéria (ou 'ok'/'sair'): ").strip()
            
            if nome_materia.lower() == 'sair':
                print("\nEncerrando o script com segurança...")
                return
            if nome_materia.lower() == 'ok':
                print(f">>> Fechando {nome_depto}. Indo para o próximo departamento...")
                break
            if not nome_materia:
                continue

            try:
                carga_horaria = int(input("Carga Horária (ex: 60): ").strip())
            except ValueError:
                print("❌ Erro: Carga horária deve ser um número. Abortando esta matéria.")
                continue

            # Validação de Pré-requisitos
            while True:
                pre_req_input = input("Pré-requisitos (separados por ',' ou 'OU', ex: CIC01_0001): ")
                matriz_pre, erro = parse_requisitos(pre_req_input, materias_existentes)
                if erro:
                    print(f"❌ Erro: Pré-requisito '{erro}' não cadastrado. Cadastre primeiro ou corrija.")
                else:
                    break 

            # Validação de Co-requisitos
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
                
                # Adiciona na memória para a próxima iteração achar o próximo número
                materias_existentes.add(codigo_materia)
                
            except Exception as e:
                print(f"❌ [ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()