package module

import (
	"context"
	"database/sql"
	"fullcycleservice/database"
	"log"
	"strconv"
	"strings"
	"time"
)

func getModules() ([]Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `
		SELECT moduleId,
			name,
			active
		FROM modules
	`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	modules := make([]Module, 0)

	for results.Next() {
		var module Module
		results.Scan(
			&module.ModuleID,
			&module.Name,
			&module.Active)
		modules = append(modules, module)
	}
	return modules, nil
}

func getModule(moduleID int) (*Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `
		SELECT moduleId,
			name,
			active
		FROM modules
		WHERE moduleId = ?
	`, moduleID)

	module := &Module{}
	err := row.Scan(
		&module.ModuleID,
		&module.Name,
		&module.Active)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Panicln(err)
		return nil, err
	}
	return module, nil
}

func insertModule(module Module) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `
		INSERT INTO modules (
			name,
			active
		) VALUES (?,?)`,
		module.Name,
		module.Active,
	)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}

func updateModule(module Module) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `
		UPDATE modules SET
			name=?,
			active=?
		WHERE moduleId=?`,
		module.Name,
		module.Active,
		module.ModuleID,
	)
	if err != nil {
		return err
	}
	return nil
}

func removeModule(moduleID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM modules WHERE moduleId = ?`, moduleID)
	if err != nil {
		return err
	}
	return nil
}

func searchForModuleData(moduleFilter ModuleFilter) ([]Module, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var queryArgs = make([]interface{}, 0)
	var queryBuilder strings.Builder
	queryBuilder.WriteString(`
		SELECT moduleId,
			LOWER(name),
			active
			FROM modules
		WHERE 1=1
	`)
	if len(moduleFilter.NameFilter) > 0 {
		queryBuilder.WriteString(`AND name LIKE ? `)
		queryArgs = append(queryArgs, "%"+strings.ToLower(moduleFilter.NameFilter)+"%")
	}
	if len(moduleFilter.ActiveFilter) > 0 {

		active, err := strconv.ParseBool(moduleFilter.ActiveFilter)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		queryBuilder.WriteString(`AND active = ? `)
		queryArgs = append(queryArgs, active)
	}
	results, err := database.DbConn.QueryContext(ctx, queryBuilder.String(), queryArgs...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	modules := make([]Module, 0)
	for results.Next() {
		var module Module
		results.Scan(
			&module.ModuleID,
			&module.Name,
			&module.Active)
		modules = append(modules, module)
	}
	return modules, nil
}
