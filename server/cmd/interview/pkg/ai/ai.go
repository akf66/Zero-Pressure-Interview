package ai

import (
	"context"
	"fmt"
	"zpi/server/cmd/interview/config"
)

// Client AI客户端接口
type Client interface {
	// GenerateFirstQuestion 生成第一个面试问题
	GenerateFirstQuestion(ctx context.Context, category string, difficulty int32, resumeContent string) (string, error)
	// GenerateNextQuestion 根据对话历史生成下一个问题
	GenerateNextQuestion(ctx context.Context, category string, history []Message) (string, bool, error)
	// EvaluateInterview 评估面试表现
	EvaluateInterview(ctx context.Context, category string, history []Message) (int32, string, error)
	// AnalyzeResume 分析简历
	AnalyzeResume(ctx context.Context, resumeContent string) (*ResumeAnalysis, error)
}

// Message 对话消息
type Message struct {
	Role    string // "interviewer" or "candidate"
	Content string
}

// ResumeAnalysis 简历分析结果
type ResumeAnalysis struct {
	Analysis         string   // 分析结果
	Suggestions      []string // 优化建议
	MatchedPositions []string // 匹配的岗位
}

// MockClient 模拟AI客户端（用于开发测试）
type MockClient struct{}

// NewMockClient 创建模拟客户端
func NewMockClient() *MockClient {
	return &MockClient{}
}

// GenerateFirstQuestion 生成第一个面试问题
func (c *MockClient) GenerateFirstQuestion(ctx context.Context, category string, difficulty int32, resumeContent string) (string, error) {
	questions := map[string][]string{
		"golang": {
			"请介绍一下 Go 语言的 goroutine 和 channel，以及它们是如何实现并发的？",
			"Go 语言中的 GC（垃圾回收）机制是怎样的？有哪些优化策略？",
			"请解释一下 Go 语言中的 interface 是如何实现的？",
		},
		"mysql": {
			"请介绍一下 MySQL 的索引类型，以及 B+ 树索引的原理。",
			"MySQL 的事务隔离级别有哪些？各自解决了什么问题？",
			"如何优化一个慢 SQL 查询？请说说你的思路。",
		},
		"redis": {
			"Redis 支持哪些数据结构？各自的使用场景是什么？",
			"Redis 的持久化机制有哪些？RDB 和 AOF 的区别是什么？",
			"如何解决 Redis 的缓存穿透、缓存击穿和缓存雪崩问题？",
		},
		"default": {
			"请做一个简单的自我介绍。",
			"你为什么选择这个技术方向？",
			"请介绍一下你最有成就感的项目经历。",
		},
	}

	categoryQuestions, ok := questions[category]
	if !ok {
		categoryQuestions = questions["default"]
	}

	// 根据难度选择问题
	idx := int(difficulty) - 1
	if idx < 0 || idx >= len(categoryQuestions) {
		idx = 0
	}

	return categoryQuestions[idx], nil
}

// GenerateNextQuestion 生成下一个问题
func (c *MockClient) GenerateNextQuestion(ctx context.Context, category string, history []Message) (string, bool, error) {
	// 模拟：每5轮对话后结束面试
	if len(history) >= 10 {
		return "", true, nil
	}

	followUpQuestions := []string{
		"能详细说说你是怎么实现的吗？",
		"在这个过程中遇到了什么困难？你是如何解决的？",
		"如果让你重新设计，你会做哪些改进？",
		"这个方案的优缺点是什么？",
		"你还有什么想补充的吗？",
	}

	idx := (len(history) / 2) % len(followUpQuestions)
	return followUpQuestions[idx], false, nil
}

// EvaluateInterview 评估面试
func (c *MockClient) EvaluateInterview(ctx context.Context, category string, history []Message) (int32, string, error) {
	// 模拟评分：根据回答数量给分
	answerCount := 0
	for _, msg := range history {
		if msg.Role == "candidate" {
			answerCount++
		}
	}

	score := int32(60 + answerCount*5)
	if score > 100 {
		score = 100
	}

	evaluation := fmt.Sprintf(`
## 面试评估报告

### 总体评分：%d/100

### 优点
- 回答问题积极主动
- 表达清晰有条理
- 技术基础扎实

### 待改进
- 可以更深入地分析问题
- 建议多举实际项目中的例子
- 可以更多地展示解决问题的思路

### 建议
继续加强 %s 相关知识的学习，多参与实际项目积累经验。
`, score, category)

	return score, evaluation, nil
}

// AnalyzeResume 分析简历
func (c *MockClient) AnalyzeResume(ctx context.Context, resumeContent string) (*ResumeAnalysis, error) {
	return &ResumeAnalysis{
		Analysis: `
## 简历分析

### 技术栈
- 后端开发：Go, Python, Java
- 数据库：MySQL, Redis, MongoDB
- 中间件：Kafka, RabbitMQ
- 云服务：AWS, Docker, Kubernetes

### 项目经验
具有丰富的微服务架构设计和开发经验，参与过多个高并发系统的设计与实现。

### 综合评价
技术栈全面，项目经验丰富，具备独立设计和开发复杂系统的能力。
`,
		Suggestions: []string{
			"建议增加更多量化的项目成果描述",
			"可以突出技术难点和解决方案",
			"建议添加开源项目或技术博客链接",
		},
		MatchedPositions: []string{
			"Golang 后端工程师",
			"分布式系统工程师",
			"云原生开发工程师",
		},
	}, nil
}

// GetClient 获取AI客户端
func GetClient() Client {
	// TODO: 根据配置返回真实的AI客户端
	_ = config.GlobalServerConfig.AIInfo
	return NewMockClient()
}
