package gochia

type BackupInfo struct {
	BackupHost string `json:"backup_host"`
	Downloaded bool   `json:"downloaded"`
}
