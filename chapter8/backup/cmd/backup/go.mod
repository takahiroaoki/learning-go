module backup

go 1.22.2

require github.com/matryer/filedb v0.0.0-20141103144311-3641db67a8c9 // indirect

require backup v0.0.0

replace backup => ../../internal/backup
