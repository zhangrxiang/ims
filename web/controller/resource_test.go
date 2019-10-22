package controller

//func TestResourceAdd(t *testing.T) {
//	//创建一个缓冲区对象,后面的要上传的body都存在这个缓冲区里
//	bodyBuf := &bytes.Buffer{}
//	bodyWriter := multipart.NewWriter(bodyBuf)
//	//要上传的文件
//	file1 := "./common.go"
//	//创建第一个需要上传的文件,filepath.Base获取文件的名称
//	fileWriter1, _ := bodyWriter.CreateFormFile("file", filepath.Base(file1))
//	//打开文件
//	fd1, _ := os.Open(file1)
//	defer fd1.Close()
//	//把第一个文件流写入到缓冲区里去
//	_, _ = io.Copy(fileWriter1, fd1)
//	//获取请求Content-Type类型,后面有用
//	contentType := bodyWriter.FormDataContentType()
//	_ = bodyWriter.Close()
//	//创建一个http客户端请求对象
//	client := &http.Client{}
//	//请求url
//	//创建一个post请求
//	req, _ := http.NewRequest("POST", "http://127.0.0.1:8081/api/v1/resource/add", bytes.NewBuffer([]byte(url.Values{
//		"name":    {"test"},
//		"desc":    {"descdescdesc"},
//		"type":    {"1"},
//		"version": {"1.0.1"},
//	}.Encode())))
//	//设置请求头
//	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
//	//这里的Content-Type值就是上面contentType的值
//	req.Header.Set("Content-Type", contentType)
//	//转换类型
//	req.Body = ioutil.NopCloser(bodyBuf)
//	//发送数据
//	data, _ := client.Do(req)
//	//读取请求返回的数据
//	bytes, _ := ioutil.ReadAll(data.Body)
//	defer data.Body.Close()
//	//打印数据
//	fmt.Println(string(bytes))
//}
