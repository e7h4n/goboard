package logical

import (
	"github.com/yuantiku/goboard/storage"
	"github.com/yuantiku/goboard/web/vo"
)

// GetDataSource retrieve data source by id
func GetDataSource(dataSourceID int, ctx *Context) (dataSource *vo.DataSource, err error) {
	mDataSource, err := storage.GetDataSource(dataSourceID, ctx.DbMap)
	if err != nil {
		return nil, err
	}

	dataSource, err = vo.NewDataSource(mDataSource)
	if err != nil {
		return nil, err
	}

	return
}

// SaveDataSource update or new data source
func SaveDataSource(projectID int, dataSource *vo.DataSource, ctx *Context) (err error) {
	if dataSource.ID > 0 {
		mDataSource, err := storage.GetDataSource(dataSource.ID, ctx.DbMap)
		if err != nil {
			return err
		}

		dataSource.ProjectID = mDataSource.ID
	} else {
		dataSource.ProjectID = projectID
	}

	mDataSource, err := dataSource.Model()

	if err != nil {
		return err
	}

	err = mDataSource.Save(ctx.DbMap)
	if err != nil {
		return err
	}

	newDataSource, err := vo.NewDataSource(mDataSource)
	if err != nil {
		return err
	}

	*dataSource = *newDataSource

	return
}

// QueryDataSource will retrieve dataSources by project and parent
func QueryDataSource(projectID int, folderID int, ctx *Context) (dataSources []vo.DataSource, err error) {
	mDataSources, err := storage.QueryDataSource(projectID, folderID, ctx.DbMap)

	dataSources = make([]vo.DataSource, len(mDataSources))
	for i, v := range mDataSources {
		dataSource, err := vo.NewDataSource(&v)
		if err != nil {
			return nil, err
		}

		dataSources[i] = *dataSource
	}

	return
}

// RemoveDataSource delete a data source
func RemoveDataSource(dataSourceID int, ctx *Context) (err error) {
	mDataSource := &storage.DataSource{ID: dataSourceID}

	return mDataSource.Remove(ctx.DbMap)
}
