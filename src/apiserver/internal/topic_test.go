package project

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/johnwtracy/personal/src/apiserver/internal/ent"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/project"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/topic"
	_ "github.com/mattn/go-sqlite3"
)

var client *ent.Client

func init() {
	var err error
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func TestTopic(t *testing.T) {
	defer client.Close()
	ctx := context.Background()

	topics := []*ent.Topic{
		client.Topic.Create().
			SetTag("1").
			SaveX(ctx),
		client.Topic.Create().
			SetTag("2").
			SaveX(ctx),
		client.Topic.Create().
			SetTag("3").
			SaveX(ctx),
	}
	projects := []*ent.Project{
		client.Project.Create().
			SetHead("H1").
			SetBody("B1").
			SetStarted(time.Now().Add(-time.Hour)).
			SaveX(ctx),
		client.Project.Create().
			SetHead("H2").SetBody("B2").
			AddTags(topics[0]).
			SetStarted(time.Now().Add(-time.Hour)).
			SetCompleted(time.Now()).
			SaveX(ctx),
		client.Project.Create().
			SetHead("H3").
			SetBody("B3").
			AddTags(topics...).
			SetStarted(time.Now()).
			SaveX(ctx),
		client.Project.Create().
			SetHead("H4").
			SetBody("B4").
			AddTags(topics[1]).
			SetStarted(time.Now()).
			SaveX(ctx),
	}

	for _, p := range projects {
		t.Run("CheckTags", func(t *testing.T) {
			// Query tags by project id
			tags := client.Topic.Query().Where(
				topic.HasProjectsWith(
					project.IDEQ(p.ID),
				),
			).AllX(ctx)

			knownTags := map[int]struct{}{}
			for _, to := range p.QueryTags().AllX(ctx) {
				knownTags[to.ID] = struct{}{}
			}

			t.Logf("expected topics: %v", knownTags)
			t.Logf("found topics: %v", tags)
			for _, to := range tags {
				if _, ok := knownTags[to.ID]; !ok {
					t.Errorf("topics are not equivalent")
				}
			}
		})
	}
}
