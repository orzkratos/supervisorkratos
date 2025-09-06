package supervisorkratos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptBasic(t *testing.T) {
	// Test basic Opt operations
	// 测试 Opt 基本操作
	opt := NewOpt("default")
	require.Equal(t, "default", opt.Get())
	require.False(t, opt.IsSet())

	opt.Set("custom")
	require.Equal(t, "custom", opt.Get())
	require.True(t, opt.IsSet())
}

func TestOptAny(t *testing.T) {
	// Test Opt[any] for AutoRestart scenarios
	// 测试 AutoRestart 场景的 Opt[any]
	opt := NewOpt[any]("unexpected")
	require.False(t, opt.IsSet())

	opt.Set(true)
	require.True(t, opt.IsSet())
	require.Equal(t, true, opt.Get())

	opt.Set("false")
	require.Equal(t, "false", opt.Get())
}
