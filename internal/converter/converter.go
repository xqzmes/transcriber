package converter

import "os/exec"

func ConvertAudio(from, to string) error {
	cmd := exec.Command("ffmpeg", "-i", from, to)
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}

	return nil
}
