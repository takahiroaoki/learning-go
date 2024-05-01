module backupd

go 1.22.2

require backup v0.0.0

require github.com/matryer/filedb v0.0.0-20141103144311-3641db67a8c9 // indirect

replace backup => ../../internal/backup
