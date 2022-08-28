package import_export

import (
	"os"
	"testing"
)

func TestImport(t *testing.T) {
	blogPath := "/Users/zouguodong/Code/Other/blog/source/_posts/Kubernetes节点管理.md"
	file, err := os.Open(blogPath)
	if err != nil {
		t.Error(err)
	}
	err = Import(file)
}
