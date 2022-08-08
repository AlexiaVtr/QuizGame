package cmd

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

type gCache struct {
	scores gcache.Cache
}

const (
	cacheSize = 1_000
	cacheTTL  = 1 * time.Hour // default expiration
)

func newGCache() *gCache {
	return &gCache{
		scores: gcache.New(cacheSize).Expiration(cacheTTL).ARC().Build(),
	}
}

func (gc *gCache) update(score int, expireIn time.Duration) error {
	return gc.scores.SetWithExpire("quizzer", score, expireIn)
}

func (gc *gCache) read(gs GameScore) error {
	var err error
	scores := gc.scores.GetALL(true)

	for _, score := range scores {
		gs.scores = append(GameScores.scores, score.(int))
	}

	if err != nil {
		return fmt.Errorf("get: %w", err)
	}

	return nil
}

func (gc *gCache) delete() {
	gc.scores.Remove("quizzer")
}
