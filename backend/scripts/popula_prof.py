import os
import firebase_admin
from firebase_admin import credentials, firestore

# 1. Autenticação 
pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

# 2. Conecta no Firestore
db = firestore.client(database_id="matriculas242")

def get_proximo_professor_id(db):
    # Trava de Segurança: Descobre o maior ID de professor que já existe no banco
    profs = db.collection("professores").stream()
    maior_id = 0
    
    for doc in profs:
        dados = doc.to_dict()
        if "professorid" in dados:
            try:
                p_id = int(dados["professorid"])
                if p_id > maior_id:
                    maior_id = p_id
            except ValueError:
                continue
                
    return maior_id + 1

def main():
    print("=== Cadastro On-Demand de Professores da UnB ===")
    print("Mapeando departamentos no banco de dados...\n")
    
    # 3. Busca a lista oficial e cria um DICIONÁRIO para busca e validação rápida
    deptos_ref = db.collection("departamentos").stream()
    mapa_deptos = {}
    
    for doc in deptos_ref:
        dados = doc.to_dict()
        # Cria a chave (ex: "CIC01") apontando para o nome (ex: "Ciência da Computação")
        mapa_deptos[dados["departamentoId"]] = dados["departamentoNome"]
    
    if not mapa_deptos:
        print("Nenhum departamento encontrado! Cadastre os departamentos primeiro.")
        return
        
    prof_serial = get_proximo_professor_id(db)
    
    print(f"Próximo Professor ID disponível: {prof_serial}\n")
    print("INSTRUÇÕES:")
    print("- Digite o CÓDIGO do departamento (ex: CIC01) para selecioná-lo.")
    print("- Dentro do departamento, digite os nomes para salvar.")
    print("- Digite 'voltar' (ou 'ok') para escolher outro departamento.")
    print("- Digite 'sair' a qualquer momento para fechar o script.\n")

    # 4. Loop Externo: Escolha do Departamento por ID
    while True:
        # Pede o ID e já converte para maiúsculo para evitar erros de digitação (cic01 -> CIC01)
        depto_id_input = input("\n🔎 Digite o Código do Departamento (ou 'sair'): ").strip().upper()
        
        if depto_id_input == 'SAIR':
            print("\nEncerrando o script com segurança...")
            break
            
        if not depto_id_input:
            continue
            
        # Trava de Integridade: Verifica se o código digitado existe no banco
        if depto_id_input not in mapa_deptos:
            print(f"❌ Erro: O departamento '{depto_id_input}' não existe no banco de dados.")
            print("Verifique se digitou corretamente (ex: MAT01, ADM01).")
            continue
            
        nome_depto = mapa_deptos[depto_id_input]
        
        print(f"\n{'='*60}")
        print(f"🏢 DEPARTAMENTO SELECIONADO: {nome_depto} (ID: {depto_id_input})")
        print(f"{'='*60}")
        
        # 5. Loop Interno: Pede os professores deste departamento específico
        while True:
            nome_prof = input(f"Nome do Professor (ou 'voltar'/'sair'): ").strip()
            
            # Se quiser abortar o script inteiro
            if nome_prof.lower() == 'sair':
                print("\nEncerrando o script com segurança...")
                return 
                
            # Se quiser trocar de departamento
            if nome_prof.lower() in ['voltar', 'ok']:
                print(f">>> Saindo de {nome_depto}...")
                break 
                
            if not nome_prof:
                print(">>> Nome vazio ignorado. Tente novamente.")
                continue
                
            prof_id_str = str(prof_serial)
            
            # 6. Payload com as tags EXATAS do seu Go model
            payload = {
                "professorid": prof_id_str,
                "professorNome": nome_prof,
                "departamentoid": depto_id_input
            }
            
            try:
                # Salva o professor na coleção dele usando o ID global
                db.collection("professores").document(prof_id_str).set(payload)
                print(f"  [SUCESSO] {nome_prof} salvo(a)! (ID: {prof_id_str})")
                
                # Incrementa o serial 
                prof_serial += 1
                
            except Exception as e:
                print(f"  [ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()