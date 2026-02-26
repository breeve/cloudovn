package utils

import (
	"reflect"
	"testing"
)

func TestContainsMaps(t *testing.T) {
	tests := []struct {
		name   string
		parent map[string]string
		subset map[string]string
		want   bool
	}{
		{
			name:   "exact match",
			parent: map[string]string{"app": "ovn", "env": "prod"},
			subset: map[string]string{"app": "ovn", "env": "prod"},
			want:   true,
		},
		{
			name:   "is a true subset",
			parent: map[string]string{"app": "ovn", "env": "prod", "region": "cn"},
			subset: map[string]string{"app": "ovn"},
			want:   true,
		},
		{
			name:   "key exists but value differs",
			parent: map[string]string{"app": "ovn"},
			subset: map[string]string{"app": "kube-ovn"},
			want:   false,
		},
		{
			name:   "key does not exist",
			parent: map[string]string{"app": "ovn"},
			subset: map[string]string{"tag": "v1"},
			want:   false,
		},
		{
			name:   "empty subset is always true",
			parent: map[string]string{"app": "ovn"},
			subset: map[string]string{},
			want:   true,
		},
		{
			name:   "nil subset is true",
			parent: map[string]string{"app": "ovn"},
			subset: nil,
			want:   true,
		},
		{
			name:   "subset has more keys than parent",
			parent: map[string]string{"app": "ovn"},
			subset: map[string]string{"app": "ovn", "extra": "val"},
			want:   false,
		},
		{
			name:   "parent is nil, subset is empty (Should be True)",
			parent: nil,
			subset: map[string]string{},
			want:   true,
		},
		{
			name:   "parent is nil, subset is nil (Should be True)",
			parent: nil,
			subset: nil,
			want:   true,
		},
		{
			name:   "parent is nil, subset has values (Should be False)",
			parent: nil,
			subset: map[string]string{"app": "ovn"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsMaps(tt.parent, tt.subset); got != tt.want {
				t.Errorf("ContainsMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name      string
		base      map[string]string
		overrides map[string]string
		want      map[string]string
	}{
		{
			name:      "standard merge without overlap",
			base:      map[string]string{"a": "1"},
			overrides: map[string]string{"b": "2"},
			want:      map[string]string{"a": "1", "b": "2"},
		},
		{
			name:      "merge with override conflict",
			base:      map[string]string{"env": "dev", "app": "ovn"},
			overrides: map[string]string{"env": "prod"},
			want:      map[string]string{"env": "prod", "app": "ovn"},
		},
		{
			name:      "base is nil",
			base:      nil,
			overrides: map[string]string{"a": "1"},
			want:      map[string]string{"a": "1"},
		},
		{
			name:      "overrides is nil",
			base:      map[string]string{"a": "1"},
			overrides: nil,
			want:      map[string]string{"a": "1"},
		},
		{
			name:      "both are nil",
			base:      nil,
			overrides: nil,
			want:      map[string]string{}, // 返回初始化后的空 map 而非 nil，利于后续操作
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeMaps(tt.base, tt.overrides)
			// 使用 reflect.DeepEqual 比较 map 内容
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps() = %v, want %v", got, tt.want)
			}

			// 验证非破坏性：确保原 map 没有被修改
			if tt.base != nil && &got == &tt.base {
				t.Errorf("MergeMaps() returned the same map instance, should return a new one")
			}
		})
	}
}
