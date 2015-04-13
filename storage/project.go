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

// QueryProject will retrieve all user projects
func QueryProject(userID int, dbmap *gorp.DbMap) (projects []Project, err error) {
	_, err = dbmap.Select(&projects, "select p.* from projects p"+
		" left join user_roles ur on (ur.project_id = p.id or ur.project_id = 0)"+
		" join roles r on r.id = ur.role_id"+
		" join role_privileges rp on rp.role_id = r.id"+
		" where ur.user_id = ? and (ur.project_id is p.id or r.scope = ?) and rp.resource = ? and rp.operation = ?", userID, RoleGlobal, ResourceProject, OperationGet)
	return
}

// GetProject will retrieve specified project by id and user
func GetProject(projectID int, userID int, dbmap *gorp.DbMap) (project *Project, err error) {
	err = dbmap.SelectOne(&project, "select p.* from projects p"+
		" left join user_roles ur on (ur.project_id = p.id or ur.project_id = 0)"+
		" join roles r on r.id = ur.role_id"+
		" join role_privileges rp on rp.role_id = r.id"+
		" where p.id = ? and ur.user_id = ? and (ur.project_id = ? or r.scope = ?) and rp.resource = ? and rp.operation = ?", projectID, userID, projectID, RoleGlobal, ResourceProject, OperationGet)
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
