package logger

import (
	"flag"
	"github.com/dwdcth/consoleEx"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Setup() {
	release := flag.Bool("release", false, "sets log level to debug")
	flag.Parse()
	if *release {
		// UNIX Time is faster and smaller than most timestamps
		// If you set zerolog.TimeFieldFormat to an empty string,
		// logs will write with UNIX time
		zerolog.TimeFieldFormat = ""
		// Default level for this example is info, unless debug flag is present
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		// Pretty logging on the console is made possible using the provided (but inefficient)
		outWriter := consoleEx.ConsoleWriterEx{Out: colorable.NewColorableStdout()}
		// 这里根据实际，另外获取的是Msg调用处的文件路径和行号
		zerolog.CallerSkipFrameCount = 2
		log.Logger = zerolog.New(outWriter).With().Caller().Timestamp().Logger()
	}
}
