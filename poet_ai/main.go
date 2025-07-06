// 文件: main.go

package main

import (
	"bufio"
	"context"
	"eino_demo/poet_ai/chat" // 确保这个路径和你的项目结构一致
	"errors"
	"fmt"
	"github.com/cloudwego/eino/schema" // *** 新增导入 ***
	"io"
	"os"
	"strings"
)

func main() {
	ctx := context.Background()
	llm := chat.CreateGeminiChatModel(ctx)
	scanner := bufio.NewScanner(os.Stdin)

	// *** 新增: 创建一个变量来存储对话历史 ***
	var chatHistory []*schema.Message

	fmt.Println("✒️  模板诗人已启动。(输入 'exit' 退出)")

	fmt.Print("> ")
	for scanner.Scan() {
		userInput := scanner.Text()

		if strings.TrimSpace(userInput) == "exit" {
			break
		}

		if strings.TrimSpace(userInput) == "" {
			fmt.Print("> ")
			continue
		}

		// 将用户输入添加到对话历史中
		messages := chat.CreateMessageFormTemplate(userInput, chatHistory)

		fmt.Print("诗人回应: ")
		// 非流输出
		//response := chat.Generate(ctx, llm, messages)
		//// 打印生成的诗歌
		//if len(response.MultiContent) == 0 {
		//	fmt.Println(response.Content)
		//} else {
		//	fmt.Println(response.MultiContent)
		//}
		////获取刚才发送给模型的用户消息
		//userMessage := messages[len(messages)-1]
		//// 将用户消息和模型的回复一起追加到历史记录中
		//chatHistory = append(chatHistory, userMessage, response)

		// 流式输出
		stream := chat.Stream(ctx, llm, messages)
		srs := stream.Copy(2)
		go func() {
			fullMsgs := make([]*schema.Message, 0)

			defer func() {
				srs[1].Close()

				chatHistory = append(chatHistory, schema.UserMessage(userInput))

				fullMsgs, err := schema.ConcatMessages(fullMsgs)
				if err != nil {
					fmt.Println("error concatenating messages: ", err.Error())
				}
				chatHistory = append(chatHistory, fullMsgs)
			}()

		outer:
			for true {
				select {
				case <-ctx.Done():
					fmt.Println("context done", ctx.Err())
					return
				default:
					chunk, err := srs[1].Recv()
					if err != nil {
						if errors.Is(err, io.EOF) {
							break outer
						}
					}
					fullMsgs = append(fullMsgs, chunk)
				}
			}
		}()

		for {
			msg, err := srs[0].Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Printf("Error receiving message: %v\n", err)
				break
			}
			fmt.Print(msg.Content)
		}

		fmt.Print("\n") // 为了格式清晰，只换一个新行
		fmt.Print("> ")
	}

	fmt.Println("\n感谢使用，再见!")
}
