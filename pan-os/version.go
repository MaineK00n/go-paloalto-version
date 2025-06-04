package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"
)

// Version represents a PAN-OS version
type Version struct {
	Major       int
	Minor       int
	Maintenance int
	Hotfix      *int
}

// NewVersion returns a parsed version
func NewVersion(ver string) (Version, error) {
	lhs, rhs, ok := strings.Cut(ver, "-")

	ss := strings.Split(lhs, ".")
	if len(ss) != 3 {
		return Version{}, fmt.Errorf("unexpected PAN-OS version format. expected: %q, actual: %q", "<major>.<minor>.<maintenance>", lhs)
	}

	major, err := strconv.Atoi(ss[0])
	if err != nil {
		return Version{}, fmt.Errorf("parse major version. err: %w", err)
	}

	minor, err := strconv.Atoi(ss[1])
	if err != nil {
		return Version{}, fmt.Errorf("parse minor version. err: %w", err)
	}

	maintenance, err := strconv.Atoi(ss[2])
	if err != nil {
		return Version{}, fmt.Errorf("parse maintenance version. err: %w", err)
	}

	v := Version{
		Major:       major,
		Minor:       minor,
		Maintenance: maintenance,
	}

	if ok {
		if !strings.HasPrefix(rhs, "h") {
			return Version{}, fmt.Errorf("unexpected PAN-OS hotfix prefix. expected: %q, actual: %q", "h", rhs)
		}

		hotfix, err := strconv.Atoi(strings.TrimPrefix(rhs, "h"))
		if err != nil {
			return Version{}, fmt.Errorf("parse hotfix version. err: %w", err)
		}

		v.Hotfix = &hotfix
	}

	return v, nil
}

// Compare returns an integer comparing two version.
// The result will be 0 if v1==v2, -1 if v1 < v2, and +1 if v1 > v2.
func (v1 Version) Compare(v2 Version) int {
	return cmp.Or(
		cmp.Compare(v1.Major, v2.Major),
		cmp.Compare(v1.Minor, v2.Minor),
		cmp.Compare(v1.Maintenance, v2.Maintenance),
		func() int {
			switch {
			case v1.Hotfix == nil && v2.Hotfix == nil:
				return 0
			case v1.Hotfix == nil && v2.Hotfix != nil:
				return -1
			case v1.Hotfix != nil && v2.Hotfix == nil:
				return +1
			default:
				return cmp.Compare(*v1.Hotfix, *v2.Hotfix)
			}
		}(),
	)
}

// String returns the full version string
func (v Version) String() string {
	if v.Hotfix == nil {
		return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Maintenance)
	}

	return fmt.Sprintf("%d.%d.%d-h%d", v.Major, v.Minor, v.Maintenance, *v.Hotfix)
}
