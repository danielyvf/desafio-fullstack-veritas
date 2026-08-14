## Sistema de Gerenciamento de Tarefas (Kanban)

Este projeto consiste em uma aplicação full stack para gerenciamento de tarefas em formato Kanban. A solução permite a criação, visualização, alteração de status e exclusão de tarefas em um fluxo contínuo.

---

## Decisões Técnicas
* Desacoplamento entre Frontend e Backend: A aplicação utiliza uma arquitetura baseada em API RESTful. O backend em Go fornece os endpoints e o frontend em React/TypeScript consome esses dados.

* Mapeamento via JSON Tags: Utilização de tags explícitas nas structs em Go (json:"title", json:"status", etc.) para assegurar que a tipagem entre os contratos do frontend (TypeScript) e a desserialização no backend permaneçam consistentes.

* Validação no Cliente e no Servidor: O frontend valida a presença do título obrigatório antes de desabilitar botões para prevenir múltiplos cliques (isSubmitting). O backend revalida os dados através do método Validate() garantindo a integridade da aplicação.

* Tratamento de Estado Local e Reatividade: O gerenciamento das tarefas no React é feito com hooks nativos (useState e useEffect), garantindo re-renderizações eficientes ao alterar colunas ou criar novos registros sem recarregar a página.

---

## Instruções para Execução

### Pré-requisitos
*  Go (versão 1.18 ou superior)
*  Node.js (versão 18 ou superior) e npm

## Executando o Backend (Go)

1. Navegue até o diretório do backend:
   ```bash
   cd backend
   ```

2. Instale as dependências (caso aplicável) e execute a aplicação:
   ```bash
   go run .
   ```
   * O servidor estará rodando na porta `8080` (`http://localhost:8080`).

3. *(Opcional)* Para compilar e validar o binário de produção:
   ```bash
   go build .
   ```

---

## Limitações Conhecidas e Melhorias Futuras

### Limitações Conhecidas
* Persistência em Memória: Os dados atualmente estão armazenados na memória runtime do backend. Reiniciar o servidor Go limpará as tarefas criadas.

* Autenticação: A aplicação não possui sistema de usuários ou controle de acesso (multitenancy).

### Melhorias Futuras
Persistência de Dados: Integração com banco de dados relacional (PostgreSQL ou SQLite) para persistência permanente.

* Drag and Drop: Implementação de suporte nativo a arrastar e soltar cartões entre as colunas usando bibliotecas como @hello-pangea/dnd.

* UI&UX: Melhorias na interface.

* Filtro e Busca: Implementação de um campo para busca textual e filtragem de tarefas por título ou descrição.

---

## Arquitetura e User Flow

### Fluxo do Usuário (User Flow)

O diagrama abaixo ilustra a jornada do usuário dentro da aplicação, cobrindo o fluxo de criação, validação e movimentação de tarefas entre os status.

```mermaid
flowchart TD
    A([Acessar a Aplicação]) --> B[Visualizar Quadro de tarefas]
    B --> C[Clicar em + Nova Tarefa]
    C --> D[Abrir Modal de Criação]
    D --> E[Preencher Título e Descrição]
    E --> F{Título foi preenchido?}
    
    F -- Não --> G[Exibir Banner de Erro]
    G --> E
    
    F -- Sim --> H[Enviar Dados para a API]
    H --> I[Tarefa Adicionada em 'A Fazer']
    I --> J[Mover para 'Em Progresso']
    J --> K[Mover para 'Concluído']
    K --> L([Fluxo Finalizado])



