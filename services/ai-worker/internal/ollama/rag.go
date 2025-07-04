package ollama

import (
	"context"
	"fmt"
	"github.com/ollama/ollama/api"
	"github.com/qdrant/go-client/qdrant"
)

const collectionName = "msl-rag"

type RAGStore struct {
	qc             *qdrant.Client
	ollama         *api.Client
	embeddingModel string
}

type RAGConfiguration struct {
	QC             *qdrant.Client
	Ollama         *api.Client
	EmbeddingModel string
}

func NewRAGStore(cfg RAGConfiguration) *RAGStore {
	return &RAGStore{
		qc:             cfg.QC,
		ollama:         cfg.Ollama,
		embeddingModel: cfg.EmbeddingModel,
	}
}

func (s *RAGStore) AddDocument(ctx context.Context, docID uint64, content string) error {
	exists, err := s.qc.CollectionExists(ctx, collectionName)
	if err != nil {
		return fmt.Errorf("failed to check if collection exists: %w", err)
	}

	if !exists {
		err := s.qc.CreateCollection(ctx, &qdrant.CreateCollection{
			CollectionName: collectionName,
			VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
				Size:     4,
				Distance: qdrant.Distance_Cosine,
			}),
		})
		if err != nil {
			return fmt.Errorf("could not create collection: %w", err)
		}
	}

	embed, err := s.CreateEmbed(ctx, content)
	if err != nil {
		return fmt.Errorf("failed to create embed: %w", err)
	}

	_, err = s.qc.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(docID),
				Vectors: qdrant.NewVectorsMulti(embed),
				Payload: map[string]*qdrant.Value{
					"content": qdrant.NewValueString(content),
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to upsert points: %w", err)
	}

	return nil
}

func (s *RAGStore) Search(ctx context.Context, query string, maxResults int) ([]string, error) {
	embedding, err := s.CreateEmbedding(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}

	points, err := s.qc.Query(ctx, &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQuery(f64Tof32(embedding)...),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to query points: %w", err)
	}

	var res []string
	for i := 0; i < max(len(res), maxResults); i++ {
		res = append(res, points[i].Payload["content"].String())
	}

	return res, nil
}

func (c *RAGStore) CreateEmbed(ctx context.Context, input string) ([][]float32, error) {
	resp, err := c.ollama.Embed(ctx, &api.EmbedRequest{
		Model: c.embeddingModel,
		Input: input,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding: %w", err)
	}

	return resp.Embeddings, nil
}

func (c *RAGStore) CreateEmbedding(ctx context.Context, prompt string) ([]float64, error) {
	resp, err := c.ollama.Embeddings(ctx, &api.EmbeddingRequest{
		Model:  c.embeddingModel,
		Prompt: prompt,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get embedding: %w", err)
	}

	return resp.Embedding, nil
}

func f64Tof32(a []float64) []float32 {
	b := make([]float32, len(a))
	for i, v := range a {
		b[i] = float32(v)
	}
	return b
}
