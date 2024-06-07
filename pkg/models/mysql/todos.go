package mysql

import (
	"database/sql"
	"todomysql/pkg/models"
)

// Defining a TodosModel type which wraps a sql.DB connection pool.
type TodosModel struct {
	DB *sql.DB
}

// creating a custom insert function for the struct TodosModel(object m)
// This will insert a new todo into the database.
func (m *TodosModel) Insert(title string) (int, error) {
	//inserting new row into table todos by using SQL statement
	stmt := `INSERT INTO todos_v2 (title, created, expires)
              VALUES(?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	//pass title and content to replace ? in SQL statement
	result, err := m.DB.Exec(stmt, title, 7)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// This will return a specific todo based on its id.
func (m *TodosModel) GetSingle(id int) (*models.Todos, error) {
	//retrieve the row where id = user input
	stmt := `SELECT * FROM todos_v2
       WHERE id = ?`
	//pass the id to replace ? in Sql statememnt
	row := m.DB.QueryRow(stmt, id)
	s := &models.Todos{}
	//row.scan is used to copy the values from database
	err := row.Scan(&s.ID, &s.Title, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *TodosModel) GetMultiple() ([]*models.Todos, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT * FROM todos_v2`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []*models.Todos{}
	for rows.Next() {
		s := &models.Todos{}
		err = rows.Scan(&s.ID, &s.Title, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		todos = append(todos, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil

}

func (m *TodosModel) Delete(ID int) error {
	stmt := `DELETE FROM todos_v2 WHERE ID = ?`
	_, err := m.DB.Exec(stmt, ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *TodosModel) Update(Id int, updateValue string) error {
	stmt := `UPDATE todos_v2 SET title = ? WHERE ID = ?`
	_, err := m.DB.Exec(stmt, updateValue, Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *TodosModel) Upadateform(ID int, title string) error {
	stmt := `UPDATE todos_v2 SET title = ? WHERE id = ?`
	_, err := m.DB.Exec(stmt, title, ID)
	if err != nil {
		return err
	}
	return nil

}
