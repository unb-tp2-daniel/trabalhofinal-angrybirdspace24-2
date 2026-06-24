import firebase_admin
from firebase_admin import credentials, auth

# 1. Configure o acesso (baixe o arquivo JSON na aba Contas de Serviço do Firebase)
cred = credentials.Certificate("ServiceAccountKey.json")
firebase_admin.initialize_app(cred)

def limpar_usuarios():
    # 2. Lista todos os usuários (limite de 1000 por vez)
    page = auth.list_users()
    
    while page:
        for user in page.users:
            # Avalia a condição (o ID no seu caso parece ser numérico)
            # Se o UID não começar com "U", vamos deletar
            if not user.uid.startswith("U"):
                print(f"Deletando usuário {user.uid}...")
                try:
                    auth.delete_user(user.uid)
                    print(f"Usuário {user.uid} excluído com sucesso.")
                except Exception as e:
                    print(f"Erro ao excluir {user.uid}: {e}")
        
        # Pega a próxima página de usuários
        page = page.get_next_page()

if __name__ == "__main__":
    limpar_usuarios()