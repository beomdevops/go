package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"park/goproject/first/models"
	"strconv"

	redis "github.com/go-redis/redis/v9"
	fiber "github.com/gofiber/fiber/v2"
)

var Rds *redis.Client

var ctx = context.Background()

func RedisConnect() {

	Rds = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		// use default DB
	})
}

type RedisTest struct {
	rd *redis.Client
}

func NewRedisTest(di_rd *redis.Client) *RedisTest {

	return &RedisTest{rd: di_rd}
}

type ModelEx struct {
	ID   int    `redis:"id"`
	Name string `redis:"name"`
}

func (r *RedisTest) SetDate(c *fiber.Ctx) error {
	r.rd.HSet(ctx, "3", "id", "3")
	r.rd.HSet(ctx, "3", "name", "yo")
	return nil
}

func (r *RedisTest) SetUser(u *models.User) error {
	r.rd.HSet(ctx, strconv.FormatUint(uint64(u.ID), 10), "id", strconv.FormatUint(uint64(u.ID), 10))
	r.rd.HSet(ctx, u.Name, "name", u.Name)

	fmt.Println(u.ID)
	fmt.Println(u.Name)
	return nil
}

func (r *RedisTest) GetData(c *fiber.Ctx) error {
	var model ModelEx
	err := r.rd.HMGet(ctx, "3", "id", "name").Scan(&model)

	if err != nil {
		panic(err)
	}
	ret, _ := json.Marshal(model)
	fmt.Println(string(ret))
	return nil
}

func (r *RedisTest) GetUser(id string) (*ModelEx, error) {
	model := &ModelEx{}
	err := r.rd.HMGet(ctx, id, "id", "name").Scan(model)

	if err != nil {
		return nil, err
	}
	if model.ID == 0 {
		return nil, errors.New("no data")
	}
	return model, nil
}
