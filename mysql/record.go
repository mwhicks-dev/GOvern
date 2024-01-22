package mysql

type Record struct {
	sid string
	usr string
	pwd string
}

func CreateRecord(sid string, usr string, pwd string) Record {
	return Record{sid, usr, pwd}
}

func GetRecordSid(record Record) string {
	return record.sid
}

func GetRecordUsr(record Record) string {
	return record.usr
}

func GetRecordPwd(record Record) string {
	return record.pwd
}
