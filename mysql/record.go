package mysql

type Record struct {
	sid string
	usr string
	pwd string
}

func CreateRecord(sid string, usr string, pwd string) Record {
	return Record{sid, usr, pwd}
}

func getRecordSid(record Record) string {
	return record.sid
}

func getRecordUsr(record Record) string {
	return record.usr
}

func getRecordPwd(record Record) string {
	return record.pwd
}
