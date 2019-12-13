package logger

import "sync"

//日志通道
type FileChannel struct {
	lock    sync.Mutex
	channel []string
}

var channels *FileChannel = &FileChannel{}

//单通道
func Channel(channelName string) *FileChannel {
	channels.lock.Lock()
	channels.channel = []string{channelName}
	return channels
}

//多通道
func Channels(channelNames []string) *FileChannel {
	channels.lock.Lock()
	channels.channel = channelNames
	return channels
}

// Painc logs a message at emergency level and panic.
func (*FileChannel) Painc(f interface{}, v ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			channels.lock.Unlock()
			panic(r)
		}
	}()
	defaultLogger.Panic(formatLog(f, v...))
}

// Fatal logs a message at emergency level and exit.
func (*FileChannel) Fatal(f interface{}, v ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			channels.lock.Unlock()
			panic(r)
		}
	}()
	defaultLogger.Fatal(formatLog(f, v...))
}

// Emer logs a message at emergency level.
func (*FileChannel) Emer(f interface{}, v ...interface{}) {
	defaultLogger.Emer(formatLog(f, v...))
	channels.lock.Unlock()
}

// Alert logs a message at alert level.
func (*FileChannel) Alert(f interface{}, v ...interface{}) {
	defaultLogger.Alert(formatLog(f, v...))
	channels.lock.Unlock()
}

// Crit logs a message at critical level.
func (*FileChannel) Crit(f interface{}, v ...interface{}) {
	defaultLogger.Crit(formatLog(f, v...))
	channels.lock.Unlock()
}

// Error logs a message at error level.
func (*FileChannel) Error(f interface{}, v ...interface{}) {
	defaultLogger.Error(formatLog(f, v...))
	channels.lock.Unlock()
}

// Warn logs a message at warning level.
func (*FileChannel) Warn(f interface{}, v ...interface{}) {
	defaultLogger.Warn(formatLog(f, v...))
	channels.lock.Unlock()
}

// Info logs a message at info level.
func (*FileChannel) Info(f interface{}, v ...interface{}) {
	defaultLogger.Info(formatLog(f, v...))
	channels.lock.Unlock()
}

// Notice logs a message at debug level.
func (*FileChannel) Debug(f interface{}, v ...interface{}) {
	defaultLogger.Debug(formatLog(f, v...))
	channels.lock.Unlock()
}

// Trace logs a message at trace level.
func (*FileChannel) Trace(f interface{}, v ...interface{}) {
	defaultLogger.Trace(formatLog(f, v...))
	channels.lock.Unlock()
}
