package pkg

import (
	hbase "zpi/server/cmd/api/biz/model/base"
	kbase "zpi/server/shared/kitex_gen/base"
)

func ConvertInterview(l *hbase.Interview) *kbase.Interview {
	return &kbase.Interview{
		UserId:     l.UserID,
		Type:       kbase.InterviewType(l.Type),
		Category:   l.Category,
		Round:      kbase.InterviewRound(l.Round),
		Status:     kbase.InterviewStatus(l.Status),
		Score:      l.Score,
		Evaluation: l.Evaluation,
		CreatedAt:  l.CreatedAt,
		FinishedAt: l.FinishedAt,
	}
}
