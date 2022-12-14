redis.go

  func (r *RedisTest) SetUser(u *models.User) error {
      data, err := json.Marshal(u.ToDto())

      if err != nil {
        return err
      }
      r.rd.HSet(ctx, strconv.FormatUint(uint64(u.ID), 10), "user", data)
      return nil
    }
    
    func (r *RedisTest) GetUser(id string) (*models.UserDto, error) {
      ret_mode := models.UserDto{}
      data, err := r.rd.HGet(ctx, id, "user").Result()

      if err != nil {
        return nil, err
      }
      //
      err = json.Unmarshal([]byte(data), &ret_mode)
      fmt.Println(ret_mode)

      if err != nil {
        return nil, err
      }

      return &ret_mode, nil
    }
