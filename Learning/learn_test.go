package learning

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")
	res := m.Run()
	fmt.Println("test_down")

	os.Exit(res)
}

func TestAdd(t *testing.T) {
	//setup
	fmt.Println("Setup")

	// t.Cleanup(func() {
	// 	fmt.Println("teardown on cleanup")
	// })
	//import data in db

	t.Run("groupeA", func(t *testing.T) {
		t.Log("groupeA")
		t.Run("simpletest", func(t *testing.T) {
			t.Parallel()
			t.Log("simpletest")
			var x, y, result = 24, 24, 48

			newresult := Add(x, y)

			if newresult != result {
				t.Errorf("expected result: %d != %d real result", result, newresult)
			}
			t.Run("moresimpletest", func(t *testing.T) {
				r := Add(2, 2)
				if r != 4 {
					t.Error("faild")
				}
			})
		})
		t.Run("mediumtest", func(t *testing.T) {
			t.Parallel()
			t.Log("mediumtest")
			var x, y, result = 314, 300, 614

			newresult := Add(x, y)

			if newresult != result {
				t.Errorf("expected result: %d != %d real result", result, newresult)
			}
		})

		t.Run("hardtest", func(t *testing.T) {
			t.Parallel()
			t.Log("hardtest")
			var x, y, result = -23945, 124832, 100887

			newresult := Add(x, y)

			if newresult != result {
				t.Errorf("expected result: %d != %d real result", result, newresult)
			}
		})
	})
	t.Run("groupeB", func(t *testing.T) {
		t.Log("gropeB")
		t.Run("simpletest", func(t *testing.T) {
			t.Parallel()
			t.Log("simpletest")
			var x, y, result = 24, 24, 48

			newresult := Add(x, y)

			if newresult != result {
				t.Errorf("expected result: %d != %d real result", result, newresult)
			}
			t.Run("moresimpletest", func(t *testing.T) {
				r := Add(2, 2)
				if r != 4 {
					t.Error("faild")
				}
			})
		})
		t.Run("mediumtest", func(t *testing.T) {
			t.Parallel()
			t.Log("mediumtest")
			var x, y, result = 314, 300, 614

			newresult := Add(x, y)

			if newresult != result {
				t.Errorf("expected result: %d != %d real result", result, newresult)
			}
		})

		t.Run("hardtest", func(t *testing.T) {
			t.Parallel()
			t.Log("hardtest")
			var x, y, result = -23945, 124832, 100887

			newresult := Add(x, y)

			if newresult != result {
				t.Errorf("expected result: %d != %d real result", result, newresult)
			}
		})
	})
	t.Run("groupeC", func(t *testing.T) {
		panic("PANIC!")
	})

	require.Equal(t, "The testing panic.")

	//testdown
	// delete test data in db
}
