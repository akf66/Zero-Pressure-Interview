package pkg

import (
	hbase "zpi/server/cmd/api/biz/model/base"
	kbase "zpi/server/shared/kitex_gen/base"
)

func ConvertInterview(l *hbase.Interview) *kbase.Interview {
	return &kbase.Interview{
		UserId:     l.UserID,
		Type:       l.Type,
		Category:   l.Category,
		Round:      l.Round,
		Status:     l.Status,
		Score:      l.Score,
		Evaluation: l.Evaluation,
		CreatedAt:  l.CreatedAt,
		FinishedAt: l.FinishedAt,
	}
}
