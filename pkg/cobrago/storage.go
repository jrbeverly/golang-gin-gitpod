package cobrago

type RemoteStorage interface {
	List(bucket string) []RemoteFile
}

type SystemWriter interface {
	Print(files []RemoteFile)
}

type RemoteFile struct {
	Key  string
	Size int64
}

func ListFilesFromStorage(bucket string, storage RemoteStorage, writer SystemWriter) {
	files := storage.List(bucket)
	writer.Print(files)
}
