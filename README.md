#  Processamento de arquivo de trânsito
### TCC ADS 2022

## 🗂️ Índice  
1. [Descrição](#-descrição)  
2. [Tecnologias Utilizadas](#-tecnologias-utilizadas)  
3. [Funcionalidades da Aplicação](#️-funcionalidades-da-aplicação)  
4. [Como Utilizar](#-como-utilizar)  
5. [Desenvolvedores](#-desenvolvedores)  

---

## 📄 Descrição  
Este projeto desenvolveu uma aplicação para análise e agrupamento de dados de acidentes de trânsito com base em arquivos CSV disponibilizados pelo Ministério dos Transportes. 

---

## 💻 Tecnologias Utilizadas  

![Go](https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge) ![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white) ![Docker](https://img.shields.io/badge/docker-257bd6?style=for-the-badge&logo=docker&logoColor=white)  

---

## ⚙️ Funcionalidades da Aplicação  

### *1. Análise de dados*  
- Análise e agrupamento de informações trazidas por arquivos CSV.
- Guarda a informação em um banco Redis.

### *2. Rotas GET para retorno de informação*  
- Rotas GET com nomes específicos definidos como base as Keys salvas no banco e parâmetros esperados para retorno de informações de anos.

---

## 📚 Como Utilizar  

### Pré-requisitos

- Configurar as variáveis de ambiente.

### Acessar drive e baixar arquivos CSV.

- https://drive.google.com/drive/folders/1r2Pr1TFnJYFOovVZJaFNv3fgOMPAAC-7?usp=sharing

### Passos para Execução

#### *1. Rode o comando abaixo:*

docker compose up --build -d

---

## 👥 Desenvolvedores  
| [![Duarte](https://github.com/duarte25.png?size=120)](https://github.com/duarte25) |
|:------------------------------------------------------------------------------------------------:|
| [Gustavo Duarte](https://github.com/duarte25) |
