package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	endpint    = "oss-cn-chengdu.aliyuncs.com"
	acessKey   = "LTAI5tMvG5NA51eiH3ENZDaa"
	secretKey  = "xxx"
	bucketName = "cloud-station"
	uploadFile = "go.sum"
)

var (
	help bool
)

func main() {
	loadParam()

	if err := validate(); err != nil {
		fmt.Printf("validate paras error, %s\n", err)
		os.Exit(1)
	}

	if err := upload(uploadFile); err != nil {
		fmt.Printf("upload file error, %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("upload file %s success\n", uploadFile)
}

func loadParam() {
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:
`)
	flag.PrintDefaults()
}

func upload(filePaht string) error {
	client, err := oss.New(endpint, acessKey, secretKey)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(filePaht, filePaht)
	if err != nil {
		return err
	}

	signedURL, err := bucket.SignURL(filePaht, oss.HTTPGet, 60*60*24)
	if err != nil {
		return fmt.Errorf("SignURL error, %s", err)
	}
	fmt.Printf("下载链接: %s\n", signedURL)
	fmt.Println("\n注意: 文件下载有效期为1天, 中转站保存时间为3天, 请及时下载")

	return nil
}

func validate() error {
	if endpint == "" {
		return fmt.Errorf("endpoint missed")
	}

	if acessKey == "" || secretKey == "" {
		return fmt.Errorf("access key or secret key missed")
	}

	if bucketName == "" {
		return fmt.Errorf("bucket name missed")
	}

	return nil
}

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&uploadFile, "f", "", "指定本地文件")
}
