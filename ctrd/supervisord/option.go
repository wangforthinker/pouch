package supervisord

import "fmt"

// WithGRPCAddress sets the containerd address.
func WithGRPCAddress(addr string) Opt {
	return func(d *Daemon) error {
		if addr == "" {
			return fmt.Errorf("grpc address should not be empty")
		}

		d.cfg.GRPC.Address = addr
		return nil
	}
}

// WithLogLevel sets the log level of containerd.
func WithLogLevel(level string) Opt {
	return func(d *Daemon) error {
		d.cfg.Debug.Level = level
		return nil
	}
}

// WithOOMScore sets the OOMScore for containerd.
func WithOOMScore(score int) Opt {
	return func(d *Daemon) error {
		if score > 1000 || score < -1000 {
			return fmt.Errorf("oom-score range should be [-1000, 1000]")
		}

		d.cfg.OOMScore = score
		return nil
	}
}

// WithContainerdBinary sets the binary name or path of containerd.
func WithContainerdBinary(nameOrPath string) Opt {
	return func(d *Daemon) error {
		d.binaryName = nameOrPath
		return nil
	}
}
