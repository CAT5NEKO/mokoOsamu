package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	var osamize int
	flag.IntVar(&osamize, "osamize", 5, "う、うわあああああああ（もこもこもこもこ）！表示するファイルの数を引数で渡してね")
	flag.Parse()

	targetPath := flag.Arg(0)
	if targetPath == "" {
		fmt.Println("使用方法: ./mokoOsamu.exe -osamize <intNum> <targetPath>")
		os.Exit(1)
	}

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		fmt.Println("Error: 指定されたパスが存在しないか、無効です。")
		os.Exit(1)
	}

	storageUsageMap := make(map[string]float64)
	err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}

		// おさマイゼーションの力をもってしても精査できなかったファイルはスキップ
		if !info.IsDir() && !strings.EqualFold(info.Name(), "System Volume Information") {
			storageUsageMap[path] = float64(info.Size())
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	sortedFiles := sortFilesByStorageUsage(storageUsageMap)

	displayTopNFiles(sortedFiles, osamize)
}

func sortFilesByStorageUsage(storageUsageMap map[string]float64) []string {
	type kv struct {
		Key   string
		Value float64
	}
	var sortedFiles []kv

	for k, v := range storageUsageMap {
		sortedFiles = append(sortedFiles, kv{k, v})
	}

	sort.Slice(sortedFiles, func(i, j int) bool {
		return sortedFiles[i].Value > sortedFiles[j].Value
	})

	var result []string
	for _, kv := range sortedFiles {
		result = append(result, fmt.Sprintf("%s (容量: %.2f MB)", kv.Key, kv.Value/(1024*1024)))
	}
	return result
}

func displayTopNFiles(files []string, n int) {
	fmt.Printf("もこ！ %d 件精査したよ。このファイル、でかいっ！:\n", n)
	for i := 0; i < n && i < len(files); i++ {
		fmt.Printf("%d. %s\n", i+1, files[i])
	}
}
