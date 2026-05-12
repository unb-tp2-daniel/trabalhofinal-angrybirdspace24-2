<template>
    <section class="profile">
        <div class="profile__header">
            <img src="../../assets/images/perfil.png" alt="foto de perfil" class="profile__avatar" />

            <div class="profile__info">
                <h3 class="profile__name">Fulano da Silva</h3>
                <p class="profile__detail"><strong>Matrícula:</strong> 252765470</p>
                <p class="profile__detail"><strong>Status:</strong> ATIVO</p>
                <p class="profile__detail"><strong>Email:</strong> email@email.com</p>
                <button type="button" class="profile__editButton">Editar perfil</button>
            </div>
        </div>

        <section class="profile__section">
            <h4 class="profile__sectionTitle">Dados acadêmicos</h4>

            <div class="profile__row">
                <span class="profile__rowLabel">Curso</span>
                <span class="profile__rowValue profile__rowValue--text">Engenharia de Computação</span>
            </div>
            <div class="profile__row">
                <span class="profile__rowLabel">Período</span>
                <span class="profile__rowValue">5º</span>
            </div>
            <div class="profile__row">
                <span class="profile__rowLabel">Ingresso</span>
                <span class="profile__rowValue">2022.1</span>
            </div>
        </section>

        <section class="profile__section">
            <h4 class="profile__sectionTitle">Médias</h4>

            <div class="profile__row">
                <span class="profile__rowLabel">IRA</span>
                <span class="profile__rowValue">2.304903</span>
            </div>
            <div class="profile__row">
                <span class="profile__rowLabel">MP</span>
                <span class="profile__rowValue">4.554</span>
            </div>
        </section>

        <section class="profile__section">
            <h4 class="profile__sectionTitle">Integralizações</h4>

            <div class="profile__row">
                <span class="profile__rowLabel">CH obrigatória pendente</span>
                <span class="profile__rowValue">—</span>
            </div>
            <div class="profile__row">
                <span class="profile__rowLabel">CH obrigatória concluída</span>
                <span class="profile__rowValue">—</span>
            </div>
            <div class="profile__row">
                <span class="profile__rowLabel">CH total do currículo</span>
                <span class="profile__rowValue">—</span>
            </div>
            <div class="profile__row">
                <span class="profile__rowLabel">CH complementar pendente</span>
                <span class="profile__rowValue">—</span>
            </div>

            <div class="profile__progress" style="--pct: 49">
                <div class="profile__progressHeader">
                    <span class="profile__progressLabel">Conclusão</span>
                    <span class="profile__progressValue">49%</span>
                </div>
                <div class="profile__progressTrack" aria-label="Percentual de conclusão">
                    <div class="profile__progressFill" />
                </div>
            </div>
        </section>
    </section>

</template>

<style scoped>
.profile {
    display: flex;
    flex-direction: column;
    gap: 12px;

    width: 340px;
    max-width: 100%;

    background-color: #ffffff;
    padding: 16px 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.12);
}

.profile__header {
    display: flex;
    align-items: flex-start;
    gap: 16px;
    min-width: 0;
}

.profile__avatar {
    width: 88px;
    height: 88px;
    border-radius: 6px;
    object-fit: cover;
    border: 2px solid #1a5276;
    flex: 0 0 auto;
}

.profile__info {
    display: flex;
    flex-direction: column;
    gap: 2px;
    min-width: 0;
}

.profile__name {
    font-size: 14px;
    font-weight: 700;
    color: #1a5276;
    line-height: 1.2;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.profile__detail {
    font-size: 12px;
    font-weight: 600;
    color: #555555;
    line-height: 1.3;
    margin: 0;
}

.profile__detail strong {
    font-weight: 700;
    color: #333333;
}

.profile__section {
    background-color: transparent;
    border: none;
    border-top: 1px solid #e0e0e0;
    padding: 12px 0;
}

.profile__sectionTitle {
    font-size: 11px;
    font-weight: 700;
    color: #999999;
    letter-spacing: 1px;
    text-transform: uppercase;
    margin-bottom: 8px;
}

.profile__row {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 12px;
    padding: 8px 0;
    border-top: 1px solid #eeeeee;
}

.profile__row:first-of-type {
    border-top: none;
    padding-top: 0;
}

.profile__rowLabel {
    font-size: 12px;
    color: #333333;
    min-width: 0;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.profile__rowValue {
    font-size: 12px;
    font-weight: 700;
    color: #1a5276;
    flex: 0 0 auto;
}

.profile__rowValue--text {
    font-weight: 600;
    color: #555555;
    max-width: 170px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.profile__progress {
    margin-top: 10px;
}

.profile__progressHeader {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
    gap: 12px;
    margin-bottom: 6px;
}

.profile__progressLabel {
    font-size: 12px;
    color: #333333;
}

.profile__progressValue {
    font-size: 12px;
    font-weight: 700;
    color: #1a5276;
}

.profile__progressTrack {
    height: 10px;
    background-color: #e0e0e0;
    border-radius: 999px;
    overflow: hidden;
}

.profile__progressFill {
    height: 100%;
    width: calc(var(--pct) * 1%);
    background-color: #1a5276;
}

.profile__editButton {
    align-self: flex-start;
    margin-top: 6px;

    background-color: #1a5276;
    color: #ffffff;
    border: none;
    border-radius: 24px;
    padding: 7px 10px;
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.5px;
    cursor: pointer;
    transition: background-color 0.2s;
}

.profile__editButton:hover {
    background-color: #154360;
}

.profile__editButton:active {
    background-color: #0e2f44;
}

@media (max-width: 420px) {
    .profile {
        padding: 12px 12px;
        gap: 10px;
    }

    .profile__avatar {
        width: 48px;
        height: 48px;
    }
}
</style>