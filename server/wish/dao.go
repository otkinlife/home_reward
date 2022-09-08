package wish

import "home-reward/server/base"

func getWishes(data *map[int64]Wish) error {
	rows, err := base.DB.Query("select * from `wish`")
	defer rows.Close()
	if err != nil {
		return err
	}
	r := *data
	for rows.Next() {
		p := Wish{}
		err := rows.Scan(&p.ID, &p.Name, &p.Reward, &p.Status, &p.CharacterID, &p.Publisher, &p.Other)
		if err != nil {
			return err
		}
		r[p.ID] = p
	}
	*data = r
	return nil
}

func getWishByID(ID int64, data *Wish) error {
	sql := "select * from `wish` where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	row := prepare.QueryRow(ID)
	if row.Err() != nil {
		return row.Err()
	}
	err = row.Scan(&data.ID, &data.Name, &data.Reward, &data.Status, &data.CharacterID, &data.Publisher, &data.Other)
	if err != nil {
		return err
	}
	return nil
}

func save(wish Wish) error {
	if wish.ID == 0 {
		return insert(wish)
	}
	c := Wish{}
	err := getWishByID(wish.ID, &c)
	if err != nil {
		return err
	}
	if c.ID == 0 {
		return insert(wish)
	} else {
		return update(wish)
	}
}

func insert(wish Wish) error {
	sql := "insert into `wish`(name, reward, status, character_id, publisher, other) values (?, ?, ?, ?, ?, ?)"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(
		wish.Name,
		wish.Reward,
		wish.Status,
		wish.CharacterID,
		wish.Publisher,
		wish.Other,
	)
	if err != nil {
		return err
	}
	return nil
}

func update(wish Wish) error {
	sql := "update `wish` set name = ?, reward = ?, status = ?, character_id = ?, publisher = ?, other = ? where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(wish.Name,
		wish.Reward,
		wish.Status,
		wish.CharacterID,
		wish.Publisher,
		wish.Other,
		wish.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func delete(wish Wish) error {
	sql := "delete from `wish` where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(
		wish.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
