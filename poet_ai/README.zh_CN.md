# Poet AI

Poet AI 是一个基于 Go 语言开发的应用程序，构建于 Eino 框架之上，能够使用语言模型生成诗歌。它支持非流式和流式的内容生成输出。

## 功能

- **交互式 CLI**：输入文本并实时接收生成的诗歌。
- **流式输出**：从语言模型中流式接收响应，提供动态体验。
- **对话历史**：维护用户输入和模型响应的历史记录。
- **Eino 框架**：利用 Eino 的 `prompt` 和 `schema` 组件进行聊天模板管理。

## 前置条件

- Go 1.20 或更高版本
- Gemini 语言模型的有效 API 密钥

## 安装

1. 克隆仓库：
   ```bash
   git clone git@github.com:Miradui/eino_demo.git
   cd eino_demo
   ```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

3. 配置 `.env` 文件，添加 Gemini API 凭据：
   ```dotenv
   GEMINI_MODEL_NAME=gemini-2.5-flash
   GEMINI_API_KEY=your-api-key
   ```

## 使用方法

1. 运行应用程序：
   ```bash
   go run poet_ai/main.go
   ```

2. 交互式使用 CLI：
    - 输入文本以生成诗歌。
    - 输入 `exit` 退出程序。

## 项目结构

- `poet_ai/main.go`：应用程序的入口文件。
- `poet_ai/chat`：包含与语言模型交互的功能。

## 环境变量

应用程序使用以下环境变量：

- `GEMINI_MODEL_NAME`：指定模型名称。
- `GEMINI_API_KEY`：用于身份验证的 API 密钥。
