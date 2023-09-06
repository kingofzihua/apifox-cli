package importc

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/imroc/req/v3"
	"github.com/spf13/cobra"
)

var projectID string
var token string

var ImportCmd = &cobra.Command{
	Use:   "import",
	Short: "import to apifox from an openapi.[json|yaml] file",
	Long: `import to apifox from an openapi.[json|yaml] file
@see: https://apifox-openapi.apifox.cn`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handler(cmd, args)
		handlerErr(err)
	},
}

func init() {
	ImportCmd.Flags().StringVar(&projectID, "project", "", "project setting -> project id")
	ImportCmd.Flags().StringVar(&token, "token", "", "see: https://apifox.com/help/openapi")
	_ = ImportCmd.MarkFlagRequired("project")
	_ = ImportCmd.MarkFlagRequired("token")
}

func handler(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("input file is mandatory")
	}

	inp := args[0] // first arg is input file

	content, err := getFileContent(inp)
	if err != nil {
		return err
	}

	return importData(content, projectID, token)
}

// 获取文件内容
func getFileContent(p string) (string, error) {
	suffix, err := parseExtension(p)
	if err != nil {
		return "", err
	}

	file, err := os.ReadFile(p)
	if err != nil {
		return "", err
	}

	switch suffix {
	case "yml":
		json, err := yaml.YAMLToJSON(file)
		return string(json), err
	case "yaml":
		json, err := yaml.YAMLToJSON(file)
		return string(json), err
	default:
		return string(file), nil
	}
}

// 获取扩展名
func parseExtension(p string) (string, error) {
	// 获取后缀名
	idx := strings.LastIndex(p, ".")
	if idx == -1 {
		return "", fmt.Errorf("the file[%s] has no extension", p)
	}
	return p[idx+1:], nil //obtain the extension in ext variable
}

func importData(json string, project string, token string) error {
	var result map[string]any

	client := req.C().DevMode()

	resp, err := client.R().
		SetHeaders(map[string]string{
			"Content-Type":     "application/json+",
			"User-Agent":       "Apifox/1.0.0 (https://apifox.com)",
			"X-Apifox-Version": "2022-11-16",
		}).
		SetBearerAuthToken(token).
		SetPathParam("project_id", project).
		SetBody(newImportDataReq(json)).
		SetSuccessResult(result).
		Post("https://api.apifox.cn/api/v1/projects/{project_id}/import-data")

	if err != nil {
		return err
	}

	if !resp.IsSuccessState() {
		return fmt.Errorf("bad response status:%s", resp.Status)
	}

	return nil
}

// importDataReq see: https://apifox-openapi.apifox.cn/
type importDataReq struct {
	ImportFormat     string `json:"importFormat"`
	Data             string `json:"data"`
	ApiOverwriteMode string `json:"apiOverwriteMode"`
}

func newImportDataReq(data string) *importDataReq {
	return &importDataReq{ImportFormat: "openapi", Data: data, ApiOverwriteMode: "methodAndPath"}
}

func handlerErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
