package sql_start

const CreateTable = "CREATE TABLE" +
	" IF NOT EXISTS msg " +
	"(" +
	"time TIMESTAMP, " +
	"tag VARCHAR(32), " +
	"msg VARCHAR(32)" +
	");  " +
	"CREATE INDEX IF NOT EXISTS msg_tag ON msg(tag)"

const DeleteTable = "DELETE TABLE IF NOT EXISTS msg "
