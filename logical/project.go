package logical

import (
	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
)

// QueryProject will retrieve all projects
func QueryProject(ctx *Context) (projects []vo.Project, err error) {
	mProjects, err := storage.QueryProject(ctx.UserID, ctx.DbMap)

	projects = make([]vo.Project, len(mProjects))
	for i, v := range mProjects {
		project := vo.NewProject(&v)
		projects[i] = *project
	}

	return
}

// GetProject retrieve project by id
func GetProject(projectID int, ctx *Context) (project *vo.Project, err error) {
	mProject, err := storage.GetProject(projectID, ctx.UserID, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	project = vo.NewProject(mProject)

	return
}

// SaveProject will save or update project
func SaveProject(project *vo.Project, ctx *Context) (err error) {
	if project.ID > 0 {
		mProject, err := storage.GetProject(project.ID, ctx.UserID, ctx.DbMap)
		if err != nil {
			return err
		}

		project.UUID = mProject.UUID
	}

	mProject := project.Model()

	err = mProject.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	newProject := vo.NewProject(mProject)

	*project = *newProject

	return
}

// RemoveProject will delete a project
func RemoveProject(projectID int, ctx *Context) (err error) {
	project, err := storage.GetProject(projectID, ctx.UserID, ctx.DbMap)
	if err != nil {
		return err
	}

	err = project.Remove(ctx.DbMap)
	return
}
