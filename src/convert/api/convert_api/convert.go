package convertapi

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Convert(input string, output string, convertType string) (string, error) {
	err := ffmpeg.Input(input).Output(output).Run()
	if err != nil {
		return "Trancoding failed", fmt.Errorf("error occured: %v", err.Error())
	}
	return "Transcoding successful", nil
}