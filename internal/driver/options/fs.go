package options

type FileSystemDriver struct {
	TasksFile string
}

func FileSystemDriverFromMap(m map[string]any) *FileSystemDriver {
	return &FileSystemDriver{
		TasksFile: getOrEmpty(m, "file"),
	}
}

func (a *FileSystemDriver) Merge(b *FileSystemDriver) *FileSystemDriver {
	return &FileSystemDriver{
		TasksFile: replaceOnEmpty(a.TasksFile, b.TasksFile),
	}
}

func (a *FileSystemDriver) ToMap() map[string]any {
	return map[string]any{
		"file": a.TasksFile,
	}
}
