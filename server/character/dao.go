package character

import (
	"home-reward/server/base"
)

func getCharacters(data *map[int64]Character) error {
	rows, err := base.DB.Query("select * from `character`")
	defer rows.Close()
	if err != nil {
		return err
	}
	r := *data
	for rows.Next() {
		c := Character{}
		err := rows.Scan(&c.ID, &c.Nick, &c.Reward, &c.Avatar)
		if err != nil {
			return err
		}
		r[c.ID] = c
	}
	*data = r
	return nil
}

func getCharacterByID(ID int64, data *Character) error {
	sql := "select * from `character` where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	row := prepare.QueryRow(ID)
	if row.Err() != nil {
		return row.Err()
	}
	err = row.Scan(&data.ID, &data.Nick, &data.Reward, &data.Avatar)
	if err != nil {
		return err
	}
	return nil
}

func save(character Character) error {
	if character.ID == 0 {
		return insert(character)
	}
	c := Character{}
	err := getCharacterByID(character.ID, &c)
	if err != nil {
		return err
	}
	if c.ID == 0 {
		return insert(character)
	} else {
		return update(character)
	}
}

func insert(character Character) error {
	sql := "insert into `character`(nick, reward, avatar) values (?, ?, ?)"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(character.Nick, character.Reward, character.Avatar)
	if err != nil {
		return err
	}
	return nil
}

func update(character Character) error {
	sql := "update `character` set nick = ?, reward = ?, avatar = ? where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(character.Nick, character.Reward, character.Avatar, character.ID)
	if err != nil {
		return err
	}
	return nil
}

func delete(character Character) error {
	sql := "delete from `character` where id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(character.ID)
	if err != nil {
		return err
	}
	return nil
}

func getCharacterByIp(ip string, data *Character) error {
	sql := "select character_id from `character_ip` where ip = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	row := prepare.QueryRow(ip)
	if row.Err() != nil {
		return row.Err()
	}
	var ID int64
	err = row.Scan(&ID)
	if err != nil {
		return err
	}
	return getCharacterByID(ID, data)
}

func bind(ip string, ID int64) error {
	sql := "insert into `character_ip`(ip, character_id) values (?, ?)"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(ip, ID)
	if err != nil {
		return err
	}
	return nil
}

func unbind(ip string, ID int64) error {
	sql := "delete from `character_ip` where ip = ? and character_id = ?"
	prepare, err := base.DB.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = prepare.Exec(ip, ID)
	if err != nil {
		return err
	}
	return nil
}
