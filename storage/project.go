package storage

import (
	"errors"
	"time"

	"github.com/nu7hatch/gouuid"

	"gopkg.in/gorp.v1"
)

// ProjectNone is used for flag global or none project
const ProjectNone = 0

// Project is a namespace to isolate datasources, users, etc...
type Project struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	UUID      string    `db:"uuid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func initProjectTable(dbmap *gorp.DbMap) {
	projectTable := dbmap.AddTableWithName(Project{}, "projects")
	projectTable.SetKeys(true, "id")
	projectTable.ColMap("name").SetUnique(true)
}

// GetAllProject will retrieve all projects
func GetAllProject(dbmap *gorp.DbMap) (projects []Project, err error) {
	_, err = dbmap.Select(&projects, "select * from projects")
	return
}

// GetProjectByID will retrieve specified project by id
func GetProjectByID(projectID int, dbmap *gorp.DbMap) (project *Project) {
	err := dbmap.SelectOne(&project, "select * from projects where id = ?", projectID)
	if err != nil {
		return nil
	}

	return
}

// Save will insert or update a project record to database
func (p *Project) Save(dbmap *gorp.DbMap) (err error) {
	if p.UUID == "" {
		uuid, err := uuid.NewV4()
		if err != nil {
			return err
		}

		p.UUID = uuid.String()
	}

	var affectCount int64
	if p.ID > 0 {
		affectCount, err = dbmap.Update(p)
	} else {
		err = dbmap.Insert(p)
		affectCount = 1
	}

	if err != nil {
		return err
	}

	if affectCount == 0 {
		return errors.New("failed to save project, affectCount = 0")
	}

	return
}

// Remove a project record from database
func (p *Project) Remove(dbmap *gorp.DbMap) (err error) {
	_, err = dbmap.Delete(p)
	return
}
