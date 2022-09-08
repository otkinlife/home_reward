package task

import "home-reward/server/base"

func getTasks(data *map[int64]Task) error {
	rows, err := base.DB.Query("select * from `task`")
	defer rows.Close()
	if err != nil {
		return err
	}
	r := *data
	for rows.Next() {
		p := Task{}
		err := rows.Scan(&p.ID, &p.Name, &p.Reward, &p.Status, &p.CharacterID)
		if err != nil {
			return err
		}
		r[p.ID] = p
	}
	*data = r
	return nil
}

func getTaskByID(ID int64, data *Task) error {
	sql := "select * from `task` where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	row := prepare.QueryRow(ID)
	if row.Err() != nil {
		return row.Err()
	}
	err = row.Scan(&data.ID, &data.Name, &data.Reward, &data.Status, &data.CharacterID)
	if err != nil {
		return err
	}
	return nil
}

func save(task Task) error {
	if task.ID == 0 {
		return insert(task)
	}
	c := Task{}
	err := getTaskByID(task.ID, &c)
	if err != nil {
		return err
	}
	if c.ID == 0 {
		return insert(task)
	} else {
		return update(task)
	}
}

func insert(task Task) error {
	sql := "insert into `task`(name, reward, status, character_id) values (?, ?, ?, ?)"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(
		task.Name,
		task.Reward,
		task.Status,
		task.CharacterID,
	)
	if err != nil {
		return err
	}
	return nil
}

func update(task Task) error {
	sql := "update `task` set name = ?, reward = ?, status = ?, character_id = ? where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(task.Name,
		task.Reward,
		task.Status,
		task.CharacterID,
		task.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func delete(task Task) error {
	sql := "delete from `task` where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(
		task.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
