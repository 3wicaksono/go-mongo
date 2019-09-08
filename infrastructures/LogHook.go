package infrastructures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-mongo/constants"
	"go-mongo/helpers"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
	sLog "log"
)

var (
	// LoggerHookFormat the default log format
	LoggerHookFormat = "{{.Level}} {{.StartTime}} {{.Message}} ==> ({{.Function}} {{.File}}:{{.Line}})\n"

	//osStat os stats
	osStat = os.Stat
)

// Hook the representation hook of logrus
type Hook struct {
	logType  string
	levels   []log.Level
	Env      string
	template *template.Template
	WLogger
	format     string
	dateFormat string
	logPath    string
	filename   string
	lWriter    *sLog.Logger
	isDebug    bool
	file       *os.File
	FormatType string
	mutex      sync.Mutex
	rotateType string
}

// LogEntry  is the structure
type LogEntry struct {
	Level     string      `json:"level"`
	StartTime string      `json:"time"`
	Message   string      `json:"message"`
	Function  string      `json:"function"`
	File      string      `json:"file"`
	Line      string      `json:"line_number"`
	Data      interface{} `json:"data"`
}

// NewLogHook the logger
func NewLogHook() *Hook {
	var appName string
	appName = config.GetString("app.name")
	h := &Hook{
		levels: []log.Level{
			log.PanicLevel,
			log.WarnLevel,
			log.ErrorLevel,
			log.FatalLevel,
			log.DebugLevel,
		},
		WLogger: sLog.New(os.Stdout, fmt.Sprintf("[%s] ", appName), 0),
		lWriter: &sLog.Logger{},
	}
	h.SetFormat(LoggerHookFormat)
	h.SetDateFormat("2006-01-02 15:04:05")
	return h
}

// WLogger interface
type WLogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// SetLogLevel logging
func (h *Hook) SetLogLevel(logLevel int) *Hook {
	switch logLevel {
	case 0:
		h.levels = []log.Level{
			log.DebugLevel,
			log.InfoLevel,
			log.WarnLevel,
			log.ErrorLevel,
			log.FatalLevel,
			log.PanicLevel,
		}
	case 1:
		h.levels = []log.Level{
			log.InfoLevel,
			log.WarnLevel,
			log.ErrorLevel,
			log.FatalLevel,
			log.PanicLevel,
		}
	case 2:
		h.levels = []log.Level{
			log.WarnLevel,
			log.ErrorLevel,
			log.FatalLevel,
			log.PanicLevel,
		}
	default:
		h.levels = []log.Level{
			log.ErrorLevel,
			log.FatalLevel,
			log.PanicLevel,
		}
	}

	return h
}

// SetRotateLog daily or static
func (h *Hook) SetRotateLog(typeName string) *Hook {

	h.rotateType = typeName
	return h
}

// Levels log level
func (h *Hook) Levels() []log.Level {
	return h.levels
}

// SetFormat the log output
func (h *Hook) SetFormat(format string) *Hook {
	h.template = template.Must(template.New("log_parser").Parse(format))
	return h
}

// SetLogType this for set log type
func (h *Hook) SetLogType(logType string) *Hook {
	h.logType = logType
	return h
}

// SetFormatType the log output
func (h *Hook) SetFormatType(formatType string) *Hook {
	h.FormatType = formatType
	return h
}

// SetDateFormat output
func (h *Hook) SetDateFormat(format string) *Hook {
	h.dateFormat = format
	return h
}

// SetLogPath location
func (h *Hook) SetLogPath(path string) *Hook {
	h.logPath = path
	return h
}

// Fire trigger event log hook
func (h *Hook) Fire(entry *log.Entry) error {

	start := time.Now()
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	lo := &LogEntry{
		Message:   entry.Message,
		Level:     strings.ToUpper(entry.Level.String()),
		StartTime: start.Format(h.dateFormat),
		Data:      entry.Data,
	}

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			lo.File = path.Base(file)
			lo.Function = path.Base(name)
			lo.Line = strconv.Itoa(line)
			break
		}
	}

	buff := &bytes.Buffer{}

	if h.FormatType == constants.LogFormatJSON {
		json.NewEncoder(buff).Encode(lo)
	} else {
		h.template.Execute(buff, lo)
	}

	if h.logType == constants.LogTypePrint {
		fmt.Print(string(buff.String()))
	}

	if h.logType == constants.LogTypeFile {
		go func() {
			h.mutex.Lock()
			logPath := config.GetString("log.path")

			rotate := logPath + "/app-" + start.Format("2006-01-02")

			if h.rotateType == "static" {
				rotate = logPath + "/app"
			}

			dir, base := filepath.Split(filepath.Clean(rotate))

			fileName := dir + base + ".log"

			if helpers.PathExist(logPath) {
				h.openNew(fileName)
				if h.file != nil {
					h.lWriter.SetOutput(h.file)
					h.lWriter.Print(string(buff.String()))
				}
			}
			h.mutex.Unlock()
		}()
	}
	return nil
}

// openNew file loc if not exist
func (h *Hook) openNew(fileLoc string) {
	if h.CurrentFileSize(fileLoc) == -1 || h.file == nil {
		if h.file != nil {
			h.file.Close()
		}
		f, err := os.OpenFile(fileLoc, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			h.Println(err)
		} else {
			h.file = f
		}
	}
	//defer mutex.Unlock()
}

// CurrentFileSize the get current file size
func (h *Hook) CurrentFileSize(fileLoc string) int64 {
	fl, err := osStat(fileLoc)
	if err != nil {
		h.WLogger.Println(err)
		return -1
	}
	return fl.Size()
}
