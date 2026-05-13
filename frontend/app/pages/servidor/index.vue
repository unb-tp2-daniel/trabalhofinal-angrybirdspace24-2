<template>
    <Header />
    <div class="dashboard">
        <aside class="sidebar">
            <h2>Sistema</h2>

            <button>Matérias</button>
            <button>Nova Matéria</button>
            <button>Matrículas</button>
            <button>Relatórios</button>
            <button>Configurações</button>
        </aside>

        <main class="content">
            <header class="topbar">
            <h1>Bem-vindo, Servidor</h1>
            </header>

            <section class="cards">

            <div class="card">
                <h3>124</h3>
                <p>Matérias</p>
            </div>

            <div class="card">
                <h3>2430</h3>
                <p>Matrículas</p>
            </div>

            <div class="card">
                <h3>18</h3>
                <p>Solicitações</p>
            </div>

            </section>

            <section class="recent">
            <h2>Últimas atividades</h2>

            <ul>
                <li>Matéria "Cálculo I" atualizada</li>
                <li>15 novas matrículas</li>
                <li>Nova disciplina criada</li>
            </ul>
            </section>
        </main>
    </div>
    <Footer/>
</template>

<style scoped>
    .dashboard{
        display:flex;
        min-height:100vh;
    }

    .sidebar{
        width:220px;
        background:#1a5276;
        color:white;
        padding:20px;
        display:flex;
        flex-direction:column;
        gap:10px;

        position:relative;
        
        border-right: 1px solid rgba(255,255,255,0.1);
        box-shadow: 4px 2px 12px rgba(0,0,0,0.25);
    }

    .sidebar button{
        margin-top: 6px;
        color: #1a5276;
        border: none;
        border-radius: 24px;
        font-size: 12px;
        font-weight: 700;
        padding:12px;
        border:none;
        cursor:pointer;
        transition: background-color 0.2s;
    }

    .sidebar h2{
        align-self: center;
        color:#f1f5f9;
    }

    .content{
        flex:1;
        padding:30px;
        background:#f1f5f9;
    }

    .cards{
        display:flex;
        gap:20px;
        margin-top:20px;
    }

    .card{
        background:white;
        padding:20px;
        border-radius:10px;
        width:180px;
        box-shadow:0 2px 5px rgba(0,0,0,0.1);
    }

    .recent{
        margin-top:40px;
    }
</style>