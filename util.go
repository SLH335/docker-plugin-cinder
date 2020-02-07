package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func getFilesystemType(dev string) (string, error) {
	out, err := exec.Command("blkid", "-s", "TYPE", "-o", "value", dev).CombinedOutput()

	if err != nil {
		if len(out) == 0 {
			return "", nil
		}

		return "", errors.New(string(out))
	}

	return string(out), nil
}

func formatFilesystem(dev string, label string) error {
	out, err := exec.Command("mkfs.ext4", "-L", label, dev).CombinedOutput()

	if err != nil {
		return errors.New(string(out))
	}

	return nil
}

func waitForDevice(dev string) error {
	_, err := os.Stat(dev)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		return nil
	}

	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)

		if _, err = os.Stat(dev); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			return nil
		}
	}

	return fmt.Errorf("Timeout waiting for file: %s", dev)
}
