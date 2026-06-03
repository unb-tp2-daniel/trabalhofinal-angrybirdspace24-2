import os
import firebase_admin
from firebase_admin import credentials, firestore

pasta_do_script = os.path.dirname(os.path.abspath(__file__))
caminho_json = os.path.join(pasta_do_script, "..", "..", "serviceAccountKey.json")

cred = credentials.Certificate(caminho_json)
if not firebase_admin._apps:
    firebase_admin.initialize_app(cred)

db = firestore.client(database_id="matriculas242")

def main():
    print("=== Cadastro On-Demand de Alunos ===")
    print("DICA: Para pular a lista de matérias concluídas/reprovadas, basta dar ENTER com a linha vazia.\n")
    
    while True:
        print("=" * 60)
        matricula = input("Matrícula (ex: 20260001) ou 'sair': ").strip()
        
        if matricula.lower() == 'sair':
            print("\nEncerrando o script com segurança...")
            break
        if not matricula:
            continue
            
        nome = input("Nome do Aluno: ").strip()
        curso_id = input("ID do Curso (ex: CIC01): ").strip().upper()
        semestre = input("Semestre atual (ex: 2026.1): ").strip()
        
        try:
            ira = float(input("IRA (ex: 4.5): ").strip())
            horas = int(input("Horas cursadas (ex: 120): ").strip())
        except ValueError:
            print("❌ Erro: IRA e Horas devem ser números. Cancelando este aluno.")
            continue
            
        # --- MATÉRIAS CONCLUÍDAS (Loop dinâmico) ---
        print("\n📚 MATÉRIAS CONCLUÍDAS")
        print("Digite no formato 'ID_MATERIA MENÇÃO' (ex: CIC01_0001 SS). Deixe vazio (ENTER) para finalizar.")
        materias_concluidas = {}
        while True:
            entrada = input("> ").strip().upper()
            if not entrada:
                break
            partes = entrada.split()
            if len(partes) >= 2:
                materia = partes[0]
                mencao = " ".join(partes[1:])
                materias_concluidas[materia] = mencao
            else:
                print("⚠️ Formato inválido. Use: ID_MATERIA MENÇÃO")
                
        # --- MATÉRIAS REPROVADAS (Loop dinâmico) ---
        print("\n❌ MATÉRIAS REPROVADAS")
        print("Digite no formato 'ID_MATERIA QTD' (ex: CIC01_0002 2). Deixe vazio (ENTER) para finalizar.")
        materias_reprovadas = {}
        while True:
            entrada = input("> ").strip().upper()
            if not entrada:
                break
            partes = entrada.split()
            if len(partes) >= 2:
                materia = partes[0]
                try:
                    qtd = int(partes[1])
                    materias_reprovadas[materia] = qtd
                except ValueError:
                    print("⚠️ A quantidade deve ser um número inteiro. Ignorando...")
            else:
                print("⚠️ Formato inválido. Use: ID_MATERIA QTD")

        # --- PRIORIDADES DE MATRÍCULA ---
        print("\n⭐ PRIORIDADES DE MATRÍCULA")
        daces = input("Possui DACES? (Sim/Nao): ").strip()
        integralizado = input("Porcentagem Integralizada (ex: 85): ").strip()
        turno = input("Turno do Aluno (Matutino/Vespertino/Noturno): ").strip()
        
        prioridades = {
            "DACES": daces,
            "integralizado": integralizado,
            "turno": turno
        }
        
        # ID do Documento espelhando a lógica do Go -> "Unb_Matricula"
        doc_id = f"Unb_{matricula}"
        
        payload = {
            "matricula": matricula,
            "nomeAluno": nome,
            "cursoId": curso_id,
            "ativo": True,
            "semestre": semestre,
            "materiasConcluidas": materias_concluidas,
            "materiasReprovadas": materias_reprovadas,
            "ira": ira,
            "prioridades": prioridades,
            "nivelAcademico": "GRADUACAO",
            "horas": horas
        }
        
        try:
            db.collection("alunos").document(doc_id).set(payload)
            print(f"\n✅ [SUCESSO] Aluno {nome} criado no banco! Documento: {doc_id}")
        except Exception as e:
            print(f"\n❌ [ERRO] Falha ao salvar no banco: {e}")

if __name__ == "__main__":
    main()