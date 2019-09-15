package breaker

import (
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// CommandFuncs is implemented by clients, mandatory
type CommandFuncs interface {
	CommandFunc()
	DefaultFunc()
	CleanupFunc()
}

// Timeout is optionally implemented by clients to override the global circuit breaker timeout
type Timeout interface {
	timeout() time.Duration
}

// Breaker struct for circuit breaker control parameters
type Breaker struct {
	name                string        // For debudding purposes
	timeout             time.Duration // Timeout at breaker level, can be reset by specific consumer
	numConcurrent       int           // Number of concurrent requests
	semaphore           chan bool     // Controls access to execute tasks
	isOk                bool          // Can circuit take more load?
	isShutdown          bool          // Has circuit been shutdown completely?
	status              int           // States for a circuit, look at consts below
	HealthCheckInterval time.Duration // Scanning interval to reset tripped circuit
}

var log *logrus.Logger

// Init initializes the circuit breaker
func (b *Breaker) Init(name string, timeout time.Duration, numConcurrent int) {
	b.name = name
	b.timeout = timeout
	b.numConcurrent = numConcurrent
	b.semaphore = make(chan bool, b.numConcurrent)
	b.isOk = true
	b.HealthCheckInterval = 100
	go scanner(b)
	log = logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
}

const (
	iShutdown        = 0
	iCircuitStillBad = 1
	iCircuitGood     = 2
)

func Logger()

func scanner(b *Breaker) {
	for {
		if b.isShutdown {
			b.status = iShutdown
			return
		}
		time.Sleep(b.HealthCheckInterval * time.Millisecond)
		if !b.isOk {
			select {
			case b.semaphore <- true:
				<-b.semaphore
				b.closeCircuit()
				log.WithFields(logrus.Fields{"name": b.name}).Info("circuit repaired, load it normal")
				b.status = iCircuitGood
			default:
				log.WithFields(logrus.Fields{"name": b.name}).Info("attempt to repair circuit failed")
				b.status = iCircuitStillBad
			}
		}
	}
}

func (b *Breaker) openCircuit() bool {
	b.isOk = false
	b.status = iCircuitStillBad
	return b.isOk
}

func (b *Breaker) closeCircuit() bool {
	b.isOk = true
	b.status = iCircuitGood
	return b.isOk
}

var mutex = &sync.Mutex{}

// Shutdown is called by clients to completely stop circuit breaker from taking any more load
func (b *Breaker) Shutdown() {
	if b.isShutdown {
		return
	}
	mutex.Lock()
	b.isShutdown = true
	mutex.Unlock()
	b.status = iShutdown
}

// Execute is called by clients to initiate task
func (b *Breaker) Execute(commands CommandFuncs) chan error {
	errorch := make(chan error, 1)
	if b.isShutdown {
		be := Error{Err: errors.New("cicuit has been permanently shutdown. create a new one")}
		errorch <- be
		return errorch
	}
	go func() {
		select {
		case b.semaphore <- true:
			go func() {
				defer func() { <-b.semaphore }()
				done := make(chan bool, 1)
				go func() {
					commands.CommandFunc()
					done <- true
				}()
				select {
				case <-time.After(b.commandTimeout(commands)):
					commands.DefaultFunc()
					commands.CleanupFunc()
					//log.WithFields(logrus.Fields{"name": b.name}).Info("task timed out")
					be := Error{isTimeout: true, Err: errors.New("task timed out")}
					errorch <- be
				case <-done:
					errorch <- nil
				}
			}()
		default:
			commands.DefaultFunc()
			commands.CleanupFunc()
			b.openCircuit()
			errorch <- errors.New("reached threshold, cannot run your function")
		}
	}()
	return errorch
}

func (b *Breaker) commandTimeout(c CommandFuncs) time.Duration {
	if t, ok := c.(Timeout); ok {
		return t.timeout()
	}
	return b.timeout
}

// Error can be unwrappd by clients to determine exact nature of failure
type Error struct {
	Err        error
	isTimeout  bool
	isShutdown bool
}

func (b Error) Unwrap() error  { return b.Err }
func (b Error) Error() string  { return b.Err.Error() }
func (b Error) Timeout() bool  { return b.isTimeout }
func (b Error) Shutdown() bool { return b.isShutdown }
