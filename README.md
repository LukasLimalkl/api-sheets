Notion to Google Sheets Integration
Descrição
Este projeto oferece uma solução automatizada para exportar dados de um banco de dados do Notion para o Google Sheets. Ideal para equipes e indivíduos que precisam sincronizar informações entre estas duas plataformas populares, facilitando a visualização, manipulação e análise de dados. A integração é realizada através de um script Python que acessa a API do Notion para extrair dados e a API do Google Sheets para inseri-los em uma planilha especificada.

Recursos
Extração de dados de bancos de dados do Notion usando a API do Notion.
Envio automático de dados para o Google Sheets.
Configuração flexível para diferentes estruturas de banco de dados.
Agendamento opcional para atualizações automáticas.
Pré-requisitos
Antes de iniciar, certifique-se de que você possui:

Python 3.8 ou superior.
Acesso ao Notion e uma integração API configurada.
Acesso ao Google Sheets e credenciais de API configuradas.
Configuração
Configuração do Ambiente
Clone o repositório:
bash
Copy code
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio

Configurando as Credenciais do Notion
Obtenha o token de integração da API do Notion seguindo as instruções aqui.
Crie um arquivo .env na raiz do projeto e adicione seu token:
makefile
Copy code
NOTION_TOKEN=<Seu_Token_Aqui>
Configurando as Credenciais do Google Sheets
Siga as instruções para criar um projeto no Google Cloud, habilitar a API do Sheets e obter as credenciais de acesso em formato JSON aqui.
Salve o arquivo JSON das credenciais na raiz do projeto.
Adicione o caminho do arquivo de credenciais ao seu arquivo .env:
makefile
Copy code
GOOGLE_APPLICATION_CREDENTIALS=<caminho_para_suas_credenciais.json>
Uso
Para executar o script e sincronizar os dados, use o seguinte comando:

bash
Copy code
python main.py
Contribuindo
Contribuições são sempre bem-vindas! Se você tem uma sugestão para melhorar esta integração, sinta-se à vontade para fazer um fork do repositório e enviar um pull request, ou abrir um issue com as tags "melhoria" ou "bug".


