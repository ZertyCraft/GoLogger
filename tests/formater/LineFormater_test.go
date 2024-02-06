package formater_test

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/ZertyCraft/GoLogger/formater"
	"github.com/ZertyCraft/GoLogger/levels"
)

func TestNewLineFormater(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want *formater.LineFormater
	}{
		{
			name: "TestNewLineFormater",
			want: &formater.LineFormater{
				BaseFormater: *formater.NewBaseFormater("%d %l %m"),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := formater.NewLineFormater(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("NewLineFormater() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestLineFormater_SetFormat(t *testing.T) {
	t.Parallel()

	type fields struct {
		format string
	}

	type args struct {
		format string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestLineFormater_SetFormat",
			fields: fields{
				format: "%d %l %m",
			},
			args: args{
				format: "%d - [%l] - %m",
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			f := &formater.LineFormater{
				BaseFormater: *formater.NewBaseFormater(test.fields.format),
			}
			f.SetFormat(test.args.format)
			if f.GetFormat() != test.args.format {
				t.Errorf("SetFormat() = %v, want %v", f.GetFormat(), test.args.format)
			}
		})
	}
}

type fields struct {
	format string
}

type args struct {
	level   levels.Level
	message string
}

func loadTestsFromFile(t *testing.T) []struct {
	name   string
	fields fields
	args   args
	want   string
} {
	t.Helper()

	data, err := os.ReadFile("LineFormater_Tests.json")
	if err != nil {
		t.Fatalf("Failed to read tests.json: %v", err)
	}

	var tests []struct {
		name   string
		fields fields
		args   args
		want   string
	}

	if err := json.Unmarshal(data, &tests); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return tests
}

// TestLineFormater_Format tests the Format method of the LineFormater struct.
func TestLineFormater_Format(t *testing.T) {
	t.Parallel()

	tests := loadTestsFromFile(t)

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			f := &formater.LineFormater{
				BaseFormater: *formater.NewBaseFormater(test.fields.format),
			}
			got, err := f.Format(test.args.level, test.args.message)
			if err != nil {
				t.Errorf("Format() error = %v", err)

				return
			}
			if got != test.want {
				t.Errorf("Format() = %v, want %v", got, test.want)
			} else {
				t.Logf("Format() = %v, want %v", got, test.want)
			}
		})
	}
}
