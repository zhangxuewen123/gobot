// Based on aplay audio adaptor written by @colemanserious (https://github.com/colemanserious)
package audio

import (
	"os/exec"
	"testing"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/gobottest"
)

var _ gobot.Adaptor = (*Adaptor)(nil)

func TestAudioAdaptor(t *testing.T) {
	a := NewAdaptor()

	gobottest.Assert(t, a.Connect(), nil)
	gobottest.Assert(t, a.Finalize(), nil)
}

func TestAudioAdaptorCommandsWav(t *testing.T) {
	cmd, _ := CommandName("whatever.wav")
	gobottest.Assert(t, cmd, "aplay")
}

func TestAudioAdaptorCommandsMp3(t *testing.T) {
	cmd, _ := CommandName("whatever.mp3")
	gobottest.Assert(t, cmd, "mpg123")
}

func TestAudioAdaptorCommandsUnknown(t *testing.T) {
	cmd, err := CommandName("whatever.unk")
	gobottest.Refute(t, cmd, "mpg123")
	gobottest.Assert(t, err.Error(), "Unknown filetype for audio file.")
}

func TestAudioAdaptorSoundWithNoFilename(t *testing.T) {
	a := NewAdaptor()

	errors := a.Sound("")
	gobottest.Assert(t, errors[0].Error(), "Requires filename for audio file.")
}

func TestAudioAdaptorSoundWithNonexistingFilename(t *testing.T) {
	a := NewAdaptor()

	errors := a.Sound("doesnotexist.mp3")
	gobottest.Assert(t, errors[0].Error(), "stat doesnotexist.mp3: no such file or directory")
}

func TestAudioAdaptorSoundWithValidMP3Filename(t *testing.T) {
	execCommand = gobottest.ExecCommand

	a := NewAdaptor()
	defer func() { execCommand = exec.Command }()

	errors := a.Sound("../../examples/laser.mp3")

	gobottest.Assert(t, len(errors), 0)
}