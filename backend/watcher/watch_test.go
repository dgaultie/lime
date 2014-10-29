package watcher

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"reflect"
	"testing"
)

func reinit() {
	var err error
	if wchr, err = fsnotify.NewWatcher(); err != nil {
		fmt.Printf("Could not create watcher due to: %s", err)
		return
	}
	watched = make(map[string][]func())
	watchers = nil
	dirs = nil
}

func dum() {
	return
}

func TestExistsIn(t *testing.T) {
	test := struct {
		array []string
		elms  []string
		exps  []bool
	}{
		[]string{"a", "b", "c", "d"},
		[]string{"a", "t", "A"},
		[]bool{true, false, false},
	}
	for i, exp := range test.exps {
		if existIn(test.array, test.elms[i]) != exp {
			t.Errorf("Expected in %v exist result of element %s be %v, but got %v", test.array, test.elms[i], exp, existIn(test.array, test.elms[i]))
		}
	}
}

func TestIsDir(t *testing.T) {
	test := struct {
		paths []string
		exps  []bool
	}{
		[]string{"testdata/dummy.txt", "testdata", ".", "test"},
		[]bool{false, true, true, false},
	}
	for i, path := range test.paths {
		if isDir(path) != test.exps[i] {
			t.Errorf("Expected %s isDir result be %v, but got %v", path, test.exps[i], isDir(path))
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		slice  []string
		remove string
		exp    []string
	}{
		{
			[]string{"a", "b", "c"},
			"a",
			[]string{"c", "b"},
		},
		{
			[]string{"a", "b", "c"},
			"k",
			[]string{"a", "b", "c"},
		},
	}
	for i, test := range tests {
		if exp := remove(test.slice, test.remove); !reflect.DeepEqual(exp, test.exp) {
			t.Errorf("Test %d: Expected %v be equal to %v", i, exp, test.exp)
		}
	}
}

func TestWatch(t *testing.T) {
	tests := []struct {
		paths       map[string]func()
		expWatched  []string
		expWatchers []string
	}{
		{
			map[string]func(){
				"testdata/dummy.txt": dum,
				"testdata/test.txt":  dum,
			},
			[]string{"testdata/dummy.txt", "testdata/test.txt"},
			[]string{"testdata/dummy.txt", "testdata/test.txt"},
		},
		{
			map[string]func(){
				"testdata":           dum,
				"testdata/dummy.txt": dum,
				"testdata/test.txt":  dum,
			},
			[]string{"testdata", "testdata/dummy.txt", "testdata/test.txt"},
			[]string{"testdata"},
		},
		{
			map[string]func(){
				"testdata/dummy.txt": dum,
				"testdata/test.txt":  dum,
				"testdata":           dum,
			},
			[]string{"testdata", "testdata/dummy.txt", "testdata/test.txt"},
			[]string{"testdata"},
		},
		// 2 path refer same file or dir but different(e.g abs path and relative path)
	}
	for i, test := range tests {
		for path, action := range test.paths {
			Watch(path, action)
		}
		if len(watched) != len(test.expWatched) {
			t.Errorf("Test %d: Expected len of watched %d, but got %d", i, len(test.expWatched), len(watched))
		}
		for _, p := range test.expWatched {
			if _, exist := watched[p]; !exist {
				t.Errorf("Test %d: Expected %s exist in watched", i, p)
			}
		}
		if !reflect.DeepEqual(test.expWatchers, watchers) {
			t.Errorf("Test %d: Expected watchers %v, but got %v", i, test.expWatchers, watchers)
		}
		for _, path := range test.expWatchers {
			UnWatch(path)
		}
		reinit()
	}
}

func TestUnWatch(t *testing.T) {
	tests := []struct {
		watchs      []string
		unWatchs    []string
		expWatched  []string
		expWatchers []string
	}{
		{
			[]string{"testdata/dummy.txt", "testdata/test.txt"},
			[]string{"testdata/dummy.txt"},
			[]string{"testdata/test.txt"},
			[]string{"testdata/test.txt"},
		},
		{
			[]string{"testdata", "testdata/dummy.txt", "testdata/test.txt"},
			[]string{"testdata"},
			[]string{"testdata/dummy.txt", "testdata/test.txt"},
			[]string{"testdata/test.txt", "testdata/dummy.txt"},
		},
	}
	for i, test := range tests {
		for _, path := range test.watchs {
			Watch(path, dum)
		}
		for _, path := range test.unWatchs {
			UnWatch(path)
		}
		if len(watched) != len(test.expWatched) {
			t.Errorf("Test %d: Expected len of watched %d, but got %d", i, len(test.expWatched), len(watched))
		}
		for _, p := range test.expWatched {
			if _, exist := watched[p]; !exist {
				t.Errorf("Test %d: Expected %s exist in watched", i, p)
			}
		}
		if !reflect.DeepEqual(test.expWatchers, watchers) {
			t.Errorf("Test %d: Expected watchers %v, but got %v", i, test.expWatchers, watchers)
		}
		for _, path := range test.expWatchers {
			UnWatch(path)
		}
		reinit()
	}
}

func TestObserve(t *testing.T) {

}

func TestObserveDirectory(t *testing.T) {

}

func TestObserveCreateEvent(t *testing.T) {

}

func TestObserveDeleteEvent(t *testing.T) {

}

func TestObserveModifyEvent(t *testing.T) {

}

func TestObserveRenameEvent(t *testing.T) {

}
