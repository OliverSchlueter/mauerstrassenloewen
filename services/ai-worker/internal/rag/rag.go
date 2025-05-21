package rag

import (
	"context"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/ai-worker/internal/ollama"
	"github.com/qdrant/go-client/qdrant"
)

const collectionName = "msl-rag"

type Store struct {
	qc     *qdrant.Client
	ollama *ollama.Client
}

type Configuration struct {
	QC     *qdrant.Client
	Ollama *ollama.Client
}

func NewStore(cfg Configuration) *Store {
	return &Store{
		qc:     cfg.QC,
		ollama: cfg.Ollama,
	}
}

func (s *Store) AddDocument(ctx context.Context, docID uint64, content string) error {
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

	embed, err := s.ollama.CreateEmbed(ctx, content)
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

func (s *Store) Search(ctx context.Context, query string, maxResults int) ([]string, error) {
	embedding, err := s.ollama.CreateEmbedding(ctx, query)
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

func f64Tof32(a []float64) []float32 {
	b := make([]float32, len(a))
	for i, v := range a {
		b[i] = float32(v)
	}
	return b
}
