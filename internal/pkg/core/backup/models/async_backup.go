package models

type AsyncBackup struct {
	RequestId string
	State     string
	Backup    *Backup
}
