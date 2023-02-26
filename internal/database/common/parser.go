package parser

import (
	"fmt"
)

const schema = "public"

func (db *DB) PushRecords(relation string, csvPath string) error {
	opts := "(FORMAT CSV, HEADER true, NULL 'null')"

	sql := fmt.Sprintf(
		"COPY \"%s\" FROM '%s' %s;",
		relation, csvPath, opts,
	)

	err := db.con.Exec(sql).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetHeader(relation string) ([]string, error) {
	sql := fmt.Sprintf(
		"SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' AND table_name = '%s';",
		schema, relation,
	)

	rows, err := db.con.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			db.logger.Fatal().Err(err).Msg("Failed to close rows")
		}
	}()

	var header []string
	for rows.Next() {
		var column string
		err = rows.Scan(&column)
		if err != nil {
			return nil, err
		}

		header = append(header, column)
	}

	return header, nil
}

func (db *DB) UpdateSequence(relation string) error {
	sql := fmt.Sprintf(
		"SELECT setval('\"%s_id_seq\"', (SELECT id FROM \"%s\" ORDER BY id DESC LIMIT 1));",
		relation, relation,
	)

	err := db.con.Exec(sql).Error
	if err != nil {
		return err
	}

	return nil
}
