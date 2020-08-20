package project

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/johnwtracy/personal/src/apiserver/internal/project/ent"
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

func TestProject(t *testing.T) {
	defer client.Close()
	ctx := context.Background()
	t1 := client.Topic.Create().
		SetTag("t1").
		SaveX(ctx)
	t2 := client.Topic.Create().
		SetTag("t2").
		SaveX(ctx)

	p1 := client.Project.Create().
		SetHead("First Project").
		SetBody("My first Ent project went a little something like...").
		AddTags(t1).
		SetStarted(time.Now().Truncate(time.Hour)).
		SaveX(ctx)
	p2 := client.Project.Create().
		SetHead("Second Project").
		SetBody("On the otherhand, my second Ent project...").
		AddTags(t1, t2).
		SetStarted(time.Now()).
		SaveX(ctx)

	t.Run("t1", func(t *testing.T) {
		for _, p := range t1.QueryProjects().AllX(ctx) {
			t.Log(p)
		}
	})

	t.Run("t2", func(t *testing.T) {
		for _, p := range t2.QueryProjects().AllX(ctx) {
			t.Log(p)
		}
	})

	t.Run("p1", func(t *testing.T) {
		for _, to := range p1.QueryTags().AllX(ctx) {
			t.Log(to)
		}
	})

	t.Run("p2", func(t *testing.T) {
		for _, to := range p2.QueryTags().AllX(ctx) {
			t.Log(to)
		}
	})

}
