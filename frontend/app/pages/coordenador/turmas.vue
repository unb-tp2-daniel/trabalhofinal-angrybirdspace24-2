<template>
    <Header />
    <div class="page">
        <div class="top-section">
            <div>
                <h1>Turmas</h1>
                <p>Gerencie todas as turmas cadastradas no sistema</p>
            </div>

            <button class="new-class-btn">
                + Nova Turma
            </button>
        </div>

        <div class="filters">
            <input type="text" placeholder="Buscar turma, professor ou código..."/>
            <select>
                <option>Todas as matérias</option>
                <option>Cálculo I</option>
                <option>Estrutura de Dados</option>
                <option>Banco de Dados</option>
            </select>
        </div>

        <section class="subject-section">
            <div class="subject-header">
                <h2>Cálculo I</h2>
                <span>3 turmas</span>
            </div>

            <table class="classes-table">
                <thead>
                    <tr>
                        <th>Turma</th>
                        <th>Código</th>
                        <th>Professor</th>
                        <th>Alunos</th>
                        <th>Período</th>
                        <th>Horário</th>
                        <th>Status</th>
                        <th>Ações</th>
                    </tr>
                </thead>

                <tbody>
                    <tr>
                        <td>Turma A</td>
                        <td>MAT607</td>
                        <td>Potchi O. Nopotushi</td>
                        <td>42</td>
                        <td>2026.1</td>
                        <td>08:00 - 10:00</td>
                        <td>
                            <span class="status warning">
                                Lotando
                            </span>
                        </td>

                        <td class="actions">
                            <button class="edit">
                                Editar
                            </button>
                            <button class="students">
                                Alunos
                            </button>
                            <button class="delete">
                                Excluir
                            </button>
                        </td>
                    </tr>

                    <tr>
                        <td>Turma B</td>
                        <td>MAT067</td>
                        <td>Wolfgang Amadeus Mozart</td>
                        <td>37</td>
                        <td>2026.1</td>
                        <td>10:00 - 12:00</td>

                        <td>
                            <span class="status active">
                                Ativa
                            </span>
                        </td>

                        <td class="actions">
                            <button class="edit">
                                Editar
                            </button>
                            <button class="students">
                                Alunos
                            </button>
                            <button class="delete">
                                Excluir
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </section>

        <section class="subject-section">
            <div class="subject-header">
                <h2>Estrutura de Dados</h2>
                <span>2 turmas</span>
            </div>

            <table class="classes-table">
                <thead>
                    <tr>
                        <th>Turma</th>
                        <th>Código</th>
                        <th>Professor</th>
                        <th>Alunos</th>
                        <th>Período</th>
                        <th>Horário</th>
                        <th>Status</th>
                        <th>Ações</th>
                    </tr>
                </thead>

                <tbody>
                    <tr>
                        <td>Turma A</td>
                        <td>CIC167</td>
                        <td>Frisquila</td>
                        <td>50</td>
                        <td>2026.1</td>
                        <td>14:00 - 16:00</td>

                        <td>
                            <span class="status full">
                                Lotada
                            </span>
                        </td>

                        <td class="actions">
                            <button class="edit">
                                Editar
                            </button>

                            <button class="students">
                                Alunos
                            </button>

                            <button class="delete">
                                Excluir
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </section>
    </div>
    <Footer />
</template>

<style scoped>

.page{
    padding:30px;
    background:#f1f5f9;
    min-height:100vh;
}

.top-section{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:30px;
}

.top-section h1{
    font-size:36px;
    margin-bottom:6px;
}

.top-section p{
    color:#1a5276;
}

.new-class-btn{
    background:#1a5276;
    color:white;
    border:none;
    padding:14px 20px;
    border-radius:10px;
    cursor:pointer;
    font-weight:700;
    transition:0.2s;
}

.new-class-btn:hover{
    opacity:0.9;
}

.filters{
    display:flex;
    gap:16px;
    margin-bottom:30px;
}

.filters input,
.filters select{
    padding:12px;
    border-radius:10px;
    border:1px solid #cbd5e1;
    background:white;
}

.filters input{
    width:320px;
}

.subject-section{
    background:white;
    border-radius:16px;
    padding:24px;
    margin-bottom:30px;
    box-shadow:0 2px 8px rgba(0,0,0,0.06);
}

.subject-header{
    display:flex;
    justify-content:space-between;
    align-items:center;
    margin-bottom:20px;
}

.subject-header h2{
    color:#1a5276;
}

.subject-header span{
    color:#1a5276;
    font-size:14px;
}

.classes-table{
    width:100%;
    border-collapse:collapse;
}

.classes-table th{
    background:#f8fafc;
    padding:14px;
    text-align:left;
    color:#334155;
    font-size:14px;
}

.classes-table td{
    padding:16px 14px;
    border-top:1px solid #e2e8f0;
}

.classes-table tr:hover{
    background:#f8fafc;
}

.actions{
    display:flex;
    gap:8px;
}

.actions button{
    border:none;
    padding:8px 12px;
    border-radius:8px;
    cursor:pointer;
    font-weight:600;
}

.edit{
    background:#dbeafe;
    color:#1d4ed8;
}

.students{
    background:#dcfce7;
    color:#166534;
}

.delete{
    background:#fee2e2;
    color:#b91c1c;
}

.status{
    padding:6px 10px;
    border-radius:999px;
    font-size:12px;
    font-weight:700;
}

.active{
    background:#dcfce7;
    color:#166534;
}

.warning{
    background:#fef3c7;
    color:#92400e;
}

.full{
    background:#fee2e2;
    color:#991b1b;
}

</style>