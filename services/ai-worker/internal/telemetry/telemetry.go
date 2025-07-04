package telemetry

import (
	"github.com/ollama/ollama/api"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Service struct {
	TotalCreatedChats        *prometheus.CounterVec
	TotalChatMessages        *prometheus.CounterVec
	OllamaTotalDuration      *prometheus.GaugeVec
	OllamaLoadDuration       *prometheus.GaugeVec
	OllamaPromptEvalCount    *prometheus.GaugeVec
	OllamaPromptEvalDuration *prometheus.GaugeVec
	OllamaEvalCount          *prometheus.GaugeVec
	OllamaEvalDuration       *prometheus.GaugeVec
}

func NewService() *Service {
	createdChats := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "created_chats",
		Help: "Total number of created chats",
	}, []string{})
	prometheus.MustRegister(createdChats)

	chatMessages := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "chat_messages",
		Help: "Total number of chat messages",
	}, []string{"model"})
	prometheus.MustRegister(chatMessages)

	ollamaTotalDuration := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ollama_total_duration",
		Help: "Total duration of chat operations",
	}, []string{"model"})
	prometheus.MustRegister(ollamaTotalDuration)

	ollamaLoadDuration := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ollama_load_duration",
		Help: "Duration of loading the Ollama model",
	}, []string{"model"})
	prometheus.MustRegister(ollamaLoadDuration)

	ollamaPromptEvalCount := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ollama_prompt_eval_count",
		Help: "Count of prompt evaluations in Ollama",
	}, []string{"model"})
	prometheus.MustRegister(ollamaPromptEvalCount)

	ollamaPromptEvalDuration := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ollama_prompt_eval_duration",
		Help: "Duration of prompt evaluations in Ollama",
	}, []string{"model"})
	prometheus.MustRegister(ollamaPromptEvalDuration)

	ollamaEvalCount := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ollama_eval_count",
		Help: "Count of evaluations in Ollama",
	}, []string{"model"})
	prometheus.MustRegister(ollamaEvalCount)

	ollamaEvalDuration := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ollama_eval_duration",
		Help: "Duration of evaluations in Ollama",
	}, []string{"model"})
	prometheus.MustRegister(ollamaEvalDuration)

	return &Service{
		TotalCreatedChats:        createdChats,
		TotalChatMessages:        chatMessages,
		OllamaTotalDuration:      ollamaTotalDuration,
		OllamaLoadDuration:       ollamaLoadDuration,
		OllamaPromptEvalCount:    ollamaPromptEvalCount,
		OllamaPromptEvalDuration: ollamaPromptEvalDuration,
		OllamaEvalCount:          ollamaEvalCount,
		OllamaEvalDuration:       ollamaEvalDuration,
	}
}

func (s *Service) RegisterHandler(mux *http.ServeMux) {
	mux.Handle("/metrics", promhttp.Handler())
}

func (s *Service) TrackNewChat() {
	s.TotalCreatedChats.With(prometheus.Labels{}).Inc()
}

func (s *Service) TrackOllamaResponse(resp *api.ChatResponse) {
	s.TotalChatMessages.With(prometheus.Labels{"model": resp.Model}).Inc()

	s.OllamaTotalDuration.WithLabelValues(resp.Model).Set(float64(resp.TotalDuration.Milliseconds()))
	s.OllamaLoadDuration.WithLabelValues(resp.Model).Set(float64(resp.LoadDuration.Milliseconds()))
	s.OllamaPromptEvalCount.WithLabelValues(resp.Model).Set(float64(resp.PromptEvalCount))
	s.OllamaPromptEvalDuration.WithLabelValues(resp.Model).Set(float64(resp.PromptEvalDuration.Milliseconds()))
	s.OllamaEvalCount.WithLabelValues(resp.Model).Set(float64(resp.EvalCount))
	s.OllamaEvalDuration.WithLabelValues(resp.Model).Set(float64(resp.EvalDuration.Milliseconds()))
}
