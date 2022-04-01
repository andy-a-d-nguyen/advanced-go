package log

var (
	destination string = "./application.log"
	lw          logWriter
)

func Run(dest string) {
}

type logWriter struct{}

func (logWriter) Write(msg []byte) (num int, err error) {
	return 0, nil
}
