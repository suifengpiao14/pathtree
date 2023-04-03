package cmd

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/logchan/v2"
)

func TestResetAllPath(t *testing.T) {
	err := ResetAllPath()
	logchan.UntilFinished(10 * time.Second)
	require.NoError(t, err)
}
